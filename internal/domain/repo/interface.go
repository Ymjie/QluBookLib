package repo

import (
	"cklib/internal/domain/DO/library"
	"context"
	"gorm.io/gorm"
)

type DBOption func(*gorm.DB) *gorm.DB
type IlkbRepo interface {
	WithSID(sid int) DBOption
	WithArea(area string) DBOption
	WithName(name string) DBOption
	GetBookInfo(ctx context.Context, opts ...DBOption) (*library.Lookbook, error)
	Add(*library.Lookbook) error
	Update(*library.Lookbook) error
	Delete(*library.Lookbook) error
}
