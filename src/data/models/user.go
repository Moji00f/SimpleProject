package models

type User struct {
	BaseModel
	UserName     string `gorm:"type:string;size:20;not null;unique;"`
	FirstName    string `gorm:"type:string;size:15;null;"`
	LastName     string `gorm:"type:string;size:25;null;"`
	MobileNumber string `gorm:"type:string;size:11;null;unique;default:null"`
	Email        string `gorm:"type:string;size:64;null;unique;default:null"`
	Password     string `gorm:"type:string;size:64;not null;"`
	Enabled      bool   `gorm:"default:true"`
	UserRoles    *[]UserRole
}

/*
user -> n roles
role -> n users

n user -> n roles

user
user_tole -> user_id, role_id, ...
role
*/