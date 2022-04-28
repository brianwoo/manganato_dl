## MangaNato Downloader

### A CLI tool to download manga from MangaNato, written in GoLang

```bash
$ go run main.go -h
Usage:
  -d string
    	Base dir for saving manga chapters (default ".")
  -from int
    	Starting chapter to download (default 1)
  -i string
    	Open manga info file as input
  -o string
    	Save manga info to file
  -to int
    	Ending chapter to download
  -u string
    	Manga URL on readmanganato

E.g.: manganato_dl -u=https://readmanganato.com/manga-hu985355 -from=1 -to=50 -d=./downloads -o=./dball.json
E.g.: manganato_dl -i=./dball.json -from=1 -to=50 -d=./downloads

Note: One of the flags: -i or -u is required
```
