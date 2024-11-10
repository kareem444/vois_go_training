package db

import (
	"errors"

	"example.com/test/core/initializers"
	"gorm.io/gorm"
)

type Options struct {
	Where    map[string]interface{}
	WhereNot map[string]interface{}
	WhereRaw []interface{}
	Limit    int
	Offset   int
	OrderBy  string
}

func Create[T any](t T) (T, error) {
	res := initializers.DB.Create(&t)

	if res.Error != nil {
		return t, res.Error
	}

	return t, nil
}

func CreateMany[T any](t []T) ([]T, error) {
	res := initializers.DB.CreateInBatches(&t, 100)

	if res.Error != nil {
		return t, res.Error
	}

	return t, nil
}

func Find[T any]() ([]T, error) {
	var t []T
	res := initializers.DB.Find(&t)

	if res.Error != nil {
		return t, res.Error
	}

	return t, nil
}

func Query[T any](options Options) ([]T, error) {
	var t []T
	query, err := _handleQuery(options)

	if err != nil {
		return t, err
	}

	res := query.Find(&t)

	if res.Error != nil {
		return t, res.Error
	}

	return t, nil
}

func Paginate[T any](page, pageSize int, options Options) ([]T, int64, error) {
	var t []T
	var totalRecords int64

	query, err := _handleQuery(options)
	if err != nil {
		return t, totalRecords, err
	}

	// Count total records
	query.Model(&t).Count(&totalRecords)

	// Apply pagination
	query = query.Offset((page - 1) * pageSize).Limit(pageSize)

	res := query.Find(&t)
	if res.Error != nil {
		return t, totalRecords, res.Error
	}

	return t, totalRecords, nil
}

func Update[T any](t T) (T, error) {
	res := initializers.DB.Save(&t)

	if res.Error != nil {
		return t, res.Error
	}

	return t, nil
}

func _handleQuery(options Options) (*gorm.DB, error) {
	query := initializers.DB

	if options.WhereRaw != nil {
		if len(options.WhereRaw) < 2 {
			return nil, errors.New("possible injections: WhereRaw should have at least 2 elements")
		} else {
			query = query.Where(options.WhereRaw[0], options.WhereRaw[1:]...)
		}
	}

	if options.Where != nil {
		query = query.Where(options.Where)
	}
	if options.WhereNot != nil {
		query = query.Not(options.WhereNot)
	}
	if options.Limit > 0 {
		query = query.Limit(options.Limit)
	}
	if options.Offset > 0 {
		query = query.Offset(options.Offset)
	}
	if options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}

	return query, nil
}
