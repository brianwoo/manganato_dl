package output_file

import (
	"encoding/json"
	"io/ioutil"
	"manganato_dl/chapter"
)

func WriteJson(fileName string, mangaInfo *chapter.MangaInfo) {
	json, err := json.MarshalIndent(mangaInfo, "", "  ")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(fileName, json, 0600)
}

func ReadJSON(fileName string) *chapter.MangaInfo {
	datas := chapter.MangaInfo{}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &datas)
	return &datas
}
