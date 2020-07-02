// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package adminmodel

import (
	"time"
)

type Connection interface {
	IsConnection()
}

type Edge interface {
	IsEdge()
}

// Common
type Node interface {
	IsNode()
}

type Article struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	Description    string    `json:"description"`
	PublishStatus  int       `json:"publishStatus"`
	ThumbnailImage string    `json:"thumbnailImage"`
	PostedAt       time.Time `json:"postedAt"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (Article) IsNode() {}

type ArticleConnection struct {
	PageInfo *PageInfo      `json:"pageInfo"`
	Edges    []*ArticleEdge `json:"edges"`
	Nodes    []*Article     `json:"nodes"`
}

func (ArticleConnection) IsConnection() {}

type ArticleEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Article `json:"node"`
}

func (ArticleEdge) IsEdge() {}

type ArticleMutationInput struct {
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	Description    string    `json:"description"`
	PublishStatus  int       `json:"publishStatus"`
	ThumbnailImage string    `json:"thumbnailImage"`
	PostedAt       time.Time `json:"postedAt"`
}

type PageInfo struct {
	Page             int  `json:"page"`
	PaginationLength int  `json:"paginationLength"`
	HasNextPage      bool `json:"hasNextPage"`
	HasPreviousPage  bool `json:"hasPreviousPage"`
	Count            int  `json:"count"`
	TotalCount       int  `json:"totalCount"`
	Limit            int  `json:"limit"`
	Offset           int  `json:"offset"`
}

type Pagination struct {
	Limit  *int `json:"Limit"`
	Offset *int `json:"Offset"`
}
