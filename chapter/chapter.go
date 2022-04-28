package chapter

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func GetChapter(chapterUrl string) ChapterInfo {

	chapterInfo := ChapterInfo{
		Title: "",
		Pages: make([]string, 0),
	}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Building Metadata: ", r.URL)
	})

	c.OnHTML("div.panel-breadcrumb span a.a-h", func(h *colly.HTMLElement) {

		link := h.Attr("href")
		if link == chapterUrl {
			chapterInfo.Title = strings.TrimSpace(h.Text)
		}
	})

	c.OnHTML("div.container-chapter-reader img[src]", func(h *colly.HTMLElement) {
		chapterInfo.Pages = append(chapterInfo.Pages, h.Attr("src"))
	})

	c.Visit(chapterUrl)
	return chapterInfo
}
