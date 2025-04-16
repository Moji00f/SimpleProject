package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id int `gorm:"primarykey"`

	CreateAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifyAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null "`
	DELETEDAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy int            `gorm:"not null"`
	ModifyBy  *sql.NullInt64 `gorm:"null"`
	DeletedBy *sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if value != nil {
		userId = int(value.(float64))
	}
	m.CreateAt = time.Now().UTC()
	m.CreatedBy = userId
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.ModifyAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.ModifyBy = userId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userId")
	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.DELETEDAt = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	m.DeletedBy = userId
	return
}
