package download

import (
	"fmt"
	"io"
	"manganato_dl/chapter"
	"manganato_dl/utils"
	"net/http"
	"os"
)

func DownloadChapter(chapterInfo chapter.ChapterInfo, saveToDirectory string) {

	saveToPath, ok := utils.CreateDirIfNotExists(saveToDirectory, chapterInfo.Title)
	fmt.Printf("Saving to %s\n", saveToPath)
	if ok {
		for fileIndex, url := range chapterInfo.Pages {
			DownloadFile(url, saveToPath, fileIndex+1)
		}
	} else {
		panic(fmt.Sprintf("Unable to create directory: %s", saveToPath))
	}

}

func DownloadFile(url string, saveToDirectoryPath string, fileIndex int) bool {

	fileExtension, ok := utils.GetFileExtension(url)
	if !ok {
		return false
	}

	fileSaveAsPath := utils.GetFileSaveAsPath(saveToDirectoryPath, fileIndex, fileExtension)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("sec-ch-ua", `"Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"`)
	req.Header.Add("Referer", "https://readmanganato.com/")
	req.Header.Add("DNT", "1")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", `"Linux"`)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fileHandle, err := os.OpenFile(fileSaveAsPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	_, err = io.Copy(fileHandle, resp.Body)
	if err != nil {
		panic(err)
	}

	return true
}
