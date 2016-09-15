// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the concurrently generated Mandelbrot fractal.
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"runtime/pprof"
	"sync"
)

/*
	Approach:
		- generate a pool of N goroutines consuming a buffered input chan of type Pixel
			of cap width*height
		- the for loops below produce Pixel onto the input chan, closing on for-loop exit
		- the workers output Pixel (with mandelbrot(z))	onto an output chan (range over the chan)
		- the main goroutine consumes the output chan (range) calling img.Set
		- a goroutine calls wg.Wait() and closes the output chan.
*/
type Pixel struct {
	x, y int
	//z    complex128
	//c    color.Color
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

/*
func buildWorkerPool(n int, wg *sync.WaitGroup, in <-chan Pixel, out chan<- Pixel) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for pixel := range in {
				out <- Pixel{x: pixel.x, y: pixel.y, c: mandelbrot(pixel.z)}
			}
		}()
	}
}
*/
func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		N                      = 5
	)
	in := make(chan *Pixel, width*height)
	//out := make(chan *Pixel, width*height)
	wg := new(sync.WaitGroup)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// build and start the worker pool
	//buildWorkerPool(N, wg, in, out)

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(img *image.RGBA) {
			defer wg.Done()
			for pixel := range in {
				x := float64(pixel.x)/width*(xmax-xmin) + xmin
				y := float64(pixel.y)/height*(ymax-ymin) + ymin
				z := complex(x, y)
				//out <- &Pixel{x: pixel.x, y: pixel.y, c: mandelbrot(pixel.z)}
				img.Set(pixel.x, pixel.y, mandelbrot(z))
			}
		}(img)
	}

	// PRODUCER
	for py := 0; py < height; py++ {
		//y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			//x := float64(px)/width*(xmax-xmin) + xmin
			//z := complex(x, y)
			// Image point (px, py) represents complex value z.
			//img.Set(px, py, mandelbrot(z))
			//in <- &Pixel{x: px, y: py, z: z}
			in <- &Pixel{x: px, y: py}
		}
	}
	// finished producing input pixels
	close(in)

	// output chan closer
	/*
		go func() {
			wg.Wait()
			close(out)
		}()
		// consume the output chan
		for p := range out {
			img.Set(p.x, p.y, p.c)
		}
	*/
	wg.Wait()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
