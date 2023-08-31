package services

import (
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/Moji00f/SimpleProject/api/dto"
	"github.com/Moji00f/SimpleProject/common"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/constants"
	"github.com/Moji00f/SimpleProject/data/db"
	"github.com/Moji00f/SimpleProject/data/models"
	"github.com/Moji00f/SimpleProject/pkg/logging"
	"github.com/Moji00f/SimpleProject/pkg/service_error"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger   logging.Logger
	preloads []preload
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
		Logger:   logging.NewLogger(cfg),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, _ := common.TypeConverter[T](req)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.
		Create(model).
		Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	// return common.TypeConverter[Tr](model)
	bm, _ := common.TypeConverter[models.BaseModel](model)

	return s.GetById(ctx, bm.Id)

}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {

	updateMap, _ := common.TypeConverter[map[string]interface{}](req)

	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}

	snakeMap["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	snakeMap["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	// (*updateMap)["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	// (*updateMap)["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	if err := tx.Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(snakeMap).
		Error; err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return s.GetById(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin()

	model := new(T)

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}

	if ctx.Value(constants.UserIdKey) == nil {
		return &service_error.ServiceError{EndUserMessage: service_error.PermissionDenied}
	}
	if cnt := tx.
		Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(deleteMap).
		RowsAffected; cnt == 0 {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, service_error.RecordNotFound, nil)
		return &service_error.ServiceError{EndUserMessage: service_error.RecordNotFound}
	}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	err := s.Database.
		Where("id = ? and deleted_by is null", id).
		First(model).
		Error
	if err != nil {
		return nil, err
	}
	return common.TypeConverter[Tr](model)
}

func NewPageList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PageList[T] {
	pl := &dto.PageList[T]{
		PageNumber: pageNumber,
		TotalRows:  count,
		Items:      items,
	}

	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasPreviousPage = pl.PageNumber > 1

	return pl
}

// Paginate
func Paginate[T any, Tr any](pagination *dto.PaginationInputWithFilter, preloads []preload, db *gorm.DB) (*dto.PageList[Tr], error) {
	model := new(T)
	var items *[]T
	var rItems *[]Tr
	db = Preload(db, preloads)
	query := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)

	var totalRows int64 = 0

	db.Where(model).Where(query).Count(&totalRows)

	err := db.Where(query).Offset(pagination.GetOffset()).Limit(pagination.GetPageSize()).Order(sort).Find(&items).Error

	if err != nil {
		return nil, err
	}

	rItems, err = common.TypeConverter[[]Tr](items)
	if err != nil {
		return nil, err
	}

	return NewPageList(rItems, totalRows, pagination.PageNumber, int64(pagination.PageSize)), err

}

// GetQuery
func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			fld, ok := typeT.FieldByName(name)
			if ok {
				fld.Name = common.ToSnakeCase(fld.Name)
				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "notcontains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From))
				case "endsWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From))
				case "equals":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEqual":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lessThan":
					query = append(query, fmt.Sprintf("%s < %s", fld.Name, filter.From))
				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > %s", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
					}
				}
			}
		}
	}

	return strings.Join(query, " AND ")

}

// GetSort
func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range *filter.Sort {
			fld, ok := typeT.FieldByName(tp.ColId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}

//Preload

func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)
	}

	return db
}
