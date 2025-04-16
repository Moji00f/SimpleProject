package model

type Country struct {
	BaseModel
	Name  string `gorm:"size:15;type:string;not null"`
	Cites []City
}

type City struct {
	BaseModel
	Name      string  `gorm:"size:10;type:string;not null"`
	CountryId int     //`gorm:"type:bigint;not null"`
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
