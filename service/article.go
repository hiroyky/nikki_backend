package service

import (
	"context"
	"github.com/hiroyky/nikki_backend/domain/dbmodel"
	"github.com/volatiletech/sqlboiler/boil"
)

func NewArticle(ctx context.Context, article dbmodel.Article) (*dbmodel.Article, error) {
	if err := article.InsertG(ctx, boil.Infer()); err != nil {
		return nil, err
	}
	return &article, nil
}

func UpdateArticle(ctx context.Context, id int, article dbmodel.Article) (*dbmodel.Article, error) {
	found, err := dbmodel.FindArticleG(ctx, id)
	if err != nil {
		return nil, err
	}

	article.ArticleID = found.ArticleID
	if _, err := article.UpdateG(ctx, boil.Infer()); err != nil {
		return nil, err
	}
	return &article, nil
}
