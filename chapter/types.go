package chapter

type ChapterInfo struct {
	Title string   `json:"title"`
	Pages []string `json:"pages"`
}

type MangaInfo struct {
	Title       string        `json:"title"`
	ChapterList []ChapterInfo `json:"chapterList"`
}
