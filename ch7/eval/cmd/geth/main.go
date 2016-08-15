package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/chzyer/readline"
	//"github.com/soudy/mathcat"
	"github.com/jfinken/gopl.io/ch7/eval"
)

var precision = flag.Int("precision", 2, "decimal precision used in results")

func getHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func repl() {

	pi := "\xCE\xA0"
	fmt.Printf("Come to the dork side we have %s\n", pi)
	// readline for the prompt and history, a go implementation of
	// the gnu libreadline library.
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      "geth> ",
		HistoryFile: getHomeDir() + "/.geth_history",
	})

	if err != nil {
		panic(err)
	}
	defer rl.Close()

	//env := eval.Env{"x": 12, "y": 1}
	g := eval.New()
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		// detect EOF before sending it to the lexer
		if line == "" {
			continue
		}
		/*
			expr, err := eval.Parse(line)
			if err != nil {
				//fmt.Println(err)
				continue
			}
			//res := fmt.Sprintf("%.6g", expr.Eval(env))
			res := expr.Eval(env)
		*/
		res := g.Run(line)
		if eval.IsWholeNumber(res) {
			fmt.Printf("%d\n", int64(res))
		} else {
			fmt.Printf("%.*f\n", *precision, res)
		}
	}
}

func main() {
	flag.Parse()
	repl()
}
