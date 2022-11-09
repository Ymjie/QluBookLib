package lkbkCrawler

import (
	lkbkUsecase "cklib/internal/usecase/lkbkCrawler"
	"cklib/pkg/logger"
)

type lkbkCrawler struct {
	Log  *logger.MyLogger
	lkbk lkbkUsecase.CrawlerUsecase
}

func NewlkbkCrawlerService(lkbk lkbkUsecase.CrawlerUsecase, Log *logger.MyLogger) *lkbkCrawler {
	return &lkbkCrawler{
		Log:  Log,
		lkbk: lkbk,
	}
}

func (c *lkbkCrawler) Run() {

}
