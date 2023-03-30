package model

type Vehicle struct {
	Brand          string `gorm:"varchar;size:30"`
	Model          string `gprm:"varchar;size:30"`
	ProductionYear int    `gorm:"size:4"`
	Color          string `gorm:"varchar;size:30"`
	IsAutomatic    bool   `gorm:"default:true"`
	Stock          int    `gorm:"check:stock >= 0"`
	SalePrice      int64  `gorm:"check:sale_price > 0"`
	Status         string `gorm:"check:status IN ('Baru', 'Bekas')"`
	BaseModel
}

func (Vehicle) TableName() string {
	return "mst_vehicle"
}

func (v *Vehicle) IsValidStatus() bool {
	return v.Status == "Baru" || v.Status == "Bekas"
}
