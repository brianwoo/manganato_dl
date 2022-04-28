package main

import (
	"flag"
	"fmt"
	"manganato_dl/chapter"
	"manganato_dl/download"
	"manganato_dl/output_file"
	"os"
)

func printParams(inputFile *string, outputFile *string, mangaUrl *string, baseUrl *string, fromChapter *int, toChapter *int) {
	if len(*inputFile) > 0 {
		fmt.Printf("InputFile   : %s\n", *inputFile)
	}
	if len(*outputFile) > 0 {
		fmt.Printf("OutputFile  : %s\n", *outputFile)
	}
	if len(*mangaUrl) > 0 {
		fmt.Printf("MangaUrl    : %s\n", *mangaUrl)
	}
	if len(*baseUrl) > 0 {
		fmt.Printf("Save To     : %s\n", *baseUrl)
	}
	fmt.Printf("From Chapter: %d\n", *fromChapter)
	if *toChapter == 0 {
		fmt.Println("To Chapter  : END")
	} else {
		fmt.Printf("To Chapter  : %d\n", *toChapter)
	}
}

func validateParams(inputFile *string, mangaUrl *string) {
	if len(*inputFile) == 0 && len(*mangaUrl) == 0 {
		panic("One of the flags: -i or -u is required. -h for usage")
	} else if len(*inputFile) > 0 && len(*mangaUrl) > 0 {
		panic("Please enter -i or -u, but not both. -h for usage")
	}
}

func getMangaInfo(inputFile *string, mangaUrl *string) *chapter.MangaInfo {

	if len(*inputFile) > 0 {
		return output_file.ReadJSON(*inputFile)
	} else {
		return chapter.GetChapterList(*mangaUrl)
	}
}

func storeMangaInfo(inputFile *string, outputFile *string, mangaInfo *chapter.MangaInfo) {
	if len(*outputFile) == 0 {
		return
	}
	if *inputFile == *outputFile {
		fmt.Println("WARNING: InputFile path is the same as OutputFile path, ignoring save...")
		return
	}

	output_file.WriteJson(*outputFile, mangaInfo)
}

func downloadChapters(baseDir *string, mangaInfo *chapter.MangaInfo, fromChapter *int, toChapter *int) {

	endChapter := 0
	if *toChapter == 0 {
		endChapter = len(mangaInfo.ChapterList)
	} else {
		endChapter = *toChapter
	}

	for i := *fromChapter - 1; i < endChapter; i++ {
		download.DownloadChapter(mangaInfo.ChapterList[i], *baseDir)
	}
}

func main() {
	inputFile := flag.String("i", "", "Open manga info file as input")
	outputFile := flag.String("o", "", "Save manga info to file")
	mangaUrl := flag.String("u", "", "Manga URL on readmanganato")
	baseDir := flag.String("d", ".", "Base dir for saving manga chapters")
	fromChapter := flag.Int("from", 1, "Starting chapter to download")
	toChapter := flag.Int("to", 0, "Ending chapter to download")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "E.g.: manganato_dl -u=https://readmanganato.com/manga-hu985355 -from=1 -to=50 -d=./downloads -o=./dball.json\n")
		fmt.Fprintf(os.Stderr, "E.g.: manganato_dl -i=./dball.json -from=1 -to=50 -d=./downloads\n\n")
		fmt.Fprintf(os.Stderr, "Note: One of the flags: -i or -u is required\n\n")
	}

	flag.Parse()

	validateParams(inputFile, mangaUrl)
	printParams(inputFile, outputFile, mangaUrl, baseDir, fromChapter, toChapter)

	mangaInfo := getMangaInfo(inputFile, mangaUrl)
	storeMangaInfo(inputFile, outputFile, mangaInfo)

	downloadChapters(baseDir, mangaInfo, fromChapter, toChapter)
}
