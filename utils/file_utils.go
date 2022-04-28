package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func GetFileExtension(url string) (string, bool) {
	if len(url) == 0 {
		return "", false
	}

	fileExtension := filepath.Ext(url)
	return fileExtension, true
}

func GetFileSaveAsPath(saveToDirectory string, fileIndex int, fileExtension string) string {
	if len(saveToDirectory) == 0 {
		saveToDirectory = "."
	}

	fileSaveAsPath := fmt.Sprintf(
		"%s%s%s%s",
		saveToDirectory,
		string(os.PathSeparator),
		strconv.Itoa(fileIndex),
		fileExtension,
	)
	return fileSaveAsPath
}

func CreateDirIfNotExists(baseDir string, dirToCreate string) (string, bool) {

	dirPath := fmt.Sprintf("%s%s%s", baseDir, string(os.PathSeparator), dirToCreate)

	// If file exists and it's not a directory
	if fileInfo, err := os.Stat(dirPath); err == nil && !fileInfo.IsDir() {
		return dirPath, false
	}

	// create new directory
	err := os.MkdirAll(dirPath, 0700)
	if err != nil {
		panic(err)
	}

	return dirPath, true
}
