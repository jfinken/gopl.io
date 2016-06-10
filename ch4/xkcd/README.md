
# IndeXKCD

One application to textually index every XKCD comic and one to search and retrieve from the index.

## NOTES:

 * Build the indexer and the search (cli) service via `make`.
 * The core index is powered by [bleve](https://github.com/blevesearch/bleve)

## Consider:

```
 $ firefox `./service/service compiling |head -n1`
 $ firefox `./service/service bobby tables |head -n1`
 $ firefox `./service/service knuth languages`
```
