package dbmodel

//go:generate go run github.com/vektah/dataloaden ArticleLoader int *github.com/hiroyky/legato/domain/dbmodel.Article
//go:generate go run github.com/vektah/dataloaden ArticleSliceLoader int []*github.com/hiroyky/legato/domain/dbmodel.Article
