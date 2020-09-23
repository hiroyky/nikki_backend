package viewerresolvers

import (
	"github.com/hiroyky/nikki_backend/domain/gql"
	"github.com/hiroyky/nikki_backend/domain/gql/viewermodel"
	"github.com/hiroyky/nikki_backend/service"
)

func ArticlesComplexity(childComplexity int, sort []*gql.SortOrder, page *viewermodel.Pagination) int {
	limit, _ := service.ValidateViewerPagination(page, 10, 100)
	return limit
}

func ArticleComplexity(childComplexity int, id string) int {
	return 1
}
