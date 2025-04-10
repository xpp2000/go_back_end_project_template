package dao

import (
	"gogofly/dto"

	"gorm.io/gorm"
)

// ====================================
// = common paging func

func Paginate(p dto.PagingDto) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		return orm.Offset((p.GetPage() - 1) * p.GetLimit()).Limit(p.GetLimit())
	}
}
