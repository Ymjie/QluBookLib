package repo

import (
	"cklib/internal/domain/DO/library"
	"cklib/internal/domain/repo"
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type LkbkRepo struct {
	db  *gorm.DB
	log *log.Logger
}

func NewLkbkRepo(db *gorm.DB, logger *log.Logger) repo.IlkbRepo {
	return &LkbkRepo{
		db:  db,
		log: logger,
	}
}

func (l *LkbkRepo) WithSID(sid int) repo.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return l.db.Where(&library.Lookbook{StuID: sid})
	}

}

func (l *LkbkRepo) WithArea(area string) repo.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return l.db.Where(&library.Lookbook{Area: area})
	}
}

func (l *LkbkRepo) WithName(name string) repo.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return l.db.Where(&library.Lookbook{Name: name})
	}
}

func (l *LkbkRepo) GetBookInfo(ctx context.Context, opts ...repo.DBOption) (*library.Lookbook, error) {
	var V1 *gorm.DB
	for k, v := range opts {
		if k == 0 {
			V1 = v(l.db)
		}
		V1 = v(V1)
	}
	var result *library.Lookbook
	V1.WithContext(ctx).First(result)
	return result, nil
}

func (l *LkbkRepo) Add(lookbook *library.Lookbook) error {
	result := l.db.Create(lookbook)
	fmt.Println(result)
	return nil
}

func (l *LkbkRepo) Update(lookbook *library.Lookbook) error {
	//TODO implement me
	panic("implement me")
}

func (l *LkbkRepo) Delete(lookbook *library.Lookbook) error {
	//TODO implement me
	panic("implement me")
}
