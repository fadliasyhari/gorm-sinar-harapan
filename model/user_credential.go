package model

type UserCredential struct {
	BaseModel
	UserName string `gorm:"unique;size:50;not null"`
	Password string `gorm:"size:10;not null"`
	IsActive bool   `gorm:"default:false"`
}

func (UserCredential) TableName() string {
	return "mst_user"
}
