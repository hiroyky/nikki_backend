// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package viewermodel

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
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}
