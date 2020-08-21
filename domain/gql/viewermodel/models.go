package viewermodel

import "time"

type Article struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Description     string    `json:"description"`
	ThumbnailImage  string    `json:"thumbnailImage"`
	PostedAt        time.Time `json:"postedAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	PreviousArticle string    `json:"previousArticle"`
	NextArticle     string    `json:"nextArticle"`
}

func (Article) IsNode() {}
