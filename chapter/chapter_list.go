package chapter

import (
	"fmt"
	"manganato_dl/utils"

	"github.com/gocolly/colly"
)

func GetChapterList(mangaUrl string) *MangaInfo {

	mangaInfo := MangaInfo{
		Title:       "",
		ChapterList: make([]ChapterInfo, 0),
	}

	chapterUrlList := make([]string, 0)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnHTML("div.story-info-right h1", func(h *colly.HTMLElement) {
		mangaInfo.Title = h.Text
	})

	c.OnHTML("ul.row-content-chapter li.a-h a", func(h *colly.HTMLElement) {
		chapterUrlList = append(chapterUrlList, h.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		utils.Reverse(chapterUrlList)
		for _, chapterUrl := range chapterUrlList {
			chapterInfo := GetChapter(chapterUrl)
			mangaInfo.ChapterList = append(mangaInfo.ChapterList, chapterInfo)
		}
	})

	c.Visit(mangaUrl)
	return &mangaInfo
}
