package usecase

import "cklib/internal/domain/DO/library"

type IlkbkUsecase interface {
	Add(lookbook *library.Lookbook)
	GET(area string) *library.Lookbook
	Update()
}
