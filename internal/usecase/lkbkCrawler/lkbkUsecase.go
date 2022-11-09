package lkbkUsecase

import (
	"cklib/internal/domain/DO/library"
	"cklib/internal/domain/repo"
	"context"
	"time"
)

func NewLkbkUsecase(repo repo.IlkbRepo) *CrawlerUsecase {
	return &CrawlerUsecase{Repo: repo}
}

func (l *CrawlerUsecase) GET(area string) *library.Lookbook {
	//l.Repo.WithName(area)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	info, _ := l.Repo.GetBookInfo(ctx, l.Repo.WithArea(area))
	return info
}

func (l *CrawlerUsecase) Update() {
	//TODO implement me
	panic("implement me")
}

func (l *CrawlerUsecase) Add(lookbook *library.Lookbook) {
	l.Repo.Add(lookbook)
}
