package main

import (
	"fmt"

	"enigmacamp.camp/gorm-sinar-harapan-makmur/config"
	"enigmacamp.camp/gorm-sinar-harapan-makmur/model"
	"gorm.io/gorm"
)

type VehicleBrandCount struct {
	Brand string
	Total int
}

func main() {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	newDbConn, err := config.NewDbConnection(c)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success Get Connection")
	}

	newDbConn.Conn()

	// SelectVehicle(db)
	// CreateVehicle(db)
}

func CreateVehicle(db *gorm.DB) {

	newVehicle := model.Vehicle{
		Brand:          "Nissan",
		Model:          "Terra",
		ProductionYear: 2022,
		Color:          "Hitam",
		IsAutomatic:    true,
		Stock:          2,
		SalePrice:      350000000,
		Status:         "Bekas",
	}

	// result := db.Debug().Save(&newVehicle)
	result := db.Select("Brand", "Model", "Stock", "SalePrice", "CreatedAt").Create(&newVehicle)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println("Success Create New Vehicle")
	}
}

func UpdateVehicle(db *gorm.DB) {
	// newVehicle := model.Vehicle{
	// 	Brand:          "Toyota",
	// 	Model:          "Fortuner",
	// 	ProductionYear: 2021,
	// 	Color:          "Hitam",
	// 	IsAutomatic:    true,
	// 	Stock:          2,
	// 	SalePrice:      450000000,
	// 	Status:         "Bekas",
	// }

	// newVehicle.ID = "48d7946a-fac7-4797-9639-b7711d5c1940"
	// newVehicle.SalePrice = 510000000
	// result := db.Debug().Save(&newVehicle)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }

	// ------------------- UPDATE SINGLE COLUMN --------------------
	// UPDATE table_name SET  bla bla where bla bla AND deleted_at IS NULL
	// newVehicle.ID = "48d7946a-fac7-4797-9639-b7711d5c1940"
	// result := db.Model(&newVehicle).Update("sale_price", 455000000)
	// Update() -> pada argument pertama itu optional
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// result := db.Debug().Model(&model.Vehicle{}).Where("Brand = ?", "Honda").Update("model", "Supra-x")
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// perintah Unscoped ini untuk mengabaikan deleted_at / menghilangkan validasi deleted_at IS NULL
	// db.Unscoped().Debug()

	// ----------------- UPDATE MULTI COLUMN ------------------------
	newVehicle := model.Vehicle{}

	newVehicle.ID = "48d7946a-fac7-4797-9639-b7711d5c1940"
	result := db.Model(&model.Vehicle{}).Where("id = ?", newVehicle.ID).Updates(&newVehicle)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func DeleteVehicle(db *gorm.DB) {

	// SOFT DELETE
	// result := db.Delete(&model.Vehicle{BaseModel: model.BaseModel{
	// 	ID: "48d7946a-fac7-4797-9639-b7711d5c1940",
	// }})
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	result := db.Delete(&model.Vehicle{}, "id = ?", "48d7946a-fac7-4797-9639-b7711d5c1940")
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	// menggunakan unscoped membuat hard Delete, menghilangkan validasi deleted_at is NULL
	// result := db.Unscoped().Delete(&model.Vehicle{}, "id = ?", "48d7946a-fac7-4797-9639-b7711d5c1940")
}

func SelectVehicle(db *gorm.DB) {
	// var vehicles []model.Vehicle

	// SELECT * FROM table_name
	// Semua records akan di assign ke dalam slice model vehicle (vehicles)
	// Note : yang terselect adalah column deleted_at yang null.
	// gunakan Unscoped() jika ingin mengabaikannya
	// result := db.Find(&vehicles)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// fmt.Println(result.RowsAffected)

	// for _, v := range vehicles {
	// 	fmt.Println("ID :", v.ID)
	// 	fmt.Println("Brand :", v.Brand)
	// 	fmt.Println("Model :", v.Model)
	// 	fmt.Println("ProductionYear :", v.ProductionYear)
	// 	fmt.Println("SalePrice :", v.SalePrice)
	// 	fmt.Println()
	// }

	// FIRST
	// var vehicle model.Vehicle
	// result := db.First(&vehicle, "id = ?", "a23a6e25-d7b1-4006-94af-d5364266258f")
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// } else if result.RowsAffected == 0 {
	// 	fmt.Println("failed to get data")
	// } else {
	// 	fmt.Println("ID :", vehicle.ID)
	// 	fmt.Println("Brand :", vehicle.Brand)
	// 	fmt.Println("Model :", vehicle.Model)
	// 	fmt.Println("ProductionYear :", vehicle.ProductionYear)
	// 	fmt.Println("SalePrice :", vehicle.SalePrice)
	// 	fmt.Println()
	// }

	// ------------------------ Refining SELECT ----------------------------------
	var vehicles []model.Vehicle

	// WHERE CLAUSE
	// result := db.Where("brand = ? AND sale_price > ?", "Toyota", 250000000).Or("color > ?", "Hitam").Or("stock > ?", 5).Find(&vehicles)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// fmt.Println(result.RowsAffected)

	// for _, v := range vehicles {
	// 	fmt.Println("ID :", v.ID)
	// 	fmt.Println("Brand :", v.Brand)
	// 	fmt.Println("Model :", v.Model)
	// 	fmt.Println("ProductionYear :", v.ProductionYear)
	// 	fmt.Println("SalePrice :", v.SalePrice)
	// 	fmt.Println()
	// }

	// IN
	// result := db.Where("model IN ?", []string{"M3", "Avanza"}).Find(&vehicles)
	// CheckErr(result.Error)
	// fmt.Println(result.RowsAffected)

	// for _, v := range vehicles {
	// 	fmt.Println("ID :", v.ID)
	// 	fmt.Println("Brand :", v.Brand)
	// 	fmt.Println("Model :", v.Model)
	// 	fmt.Println("ProductionYear :", v.ProductionYear)
	// 	fmt.Println("SalePrice :", v.SalePrice)
	// 	fmt.Println()
	// }

	// LIKE
	// result := db.Where("model LIKE ?", "%-x%").Find(&vehicles)
	// CheckErr(result)
	// fmt.Println(result)

	// limit offset
	// page := 2
	// itemPerPage := 2
	// offset := itemPerPage * (page - 1)
	// result := db.Order("created_at").Limit(itemPerPage).Offset(offset).Find(&vehicles)
	// CheckErr(result)

	// for _, v := range vehicles {
	// 	fmt.Println("ID :", v.ID)
	// 	fmt.Println("Brand :", v.Brand)
	// 	fmt.Println("Model :", v.Model)
	// 	fmt.Println("ProductionYear :", v.ProductionYear)
	// 	fmt.Println("SalePrice :", v.SalePrice)
	// 	fmt.Println()
	// }

	// COUNT
	// var total int64
	// result := db.Debug().Model(&vehicles).Count(&total)
	// CheckErr(result)
	// fmt.Println("total kendaraan:", total)

	// COUNT with GROUPBY
	var vehicleBrandCount []VehicleBrandCount
	result := db.Model(&vehicles).Select("brand, COUNT(*) AS total").Group("brand").Find(&vehicleBrandCount)
	CheckErr(result)
	for _, v := range vehicleBrandCount {
		fmt.Println("Brand :", v.Brand)
		fmt.Println("Total :", v.Total)
	}

}

func CheckErr(result *gorm.DB) {
	if result.Error != nil {
		fmt.Println(result.Error)
	} else if result.RowsAffected == 0 {
		fmt.Println("Data tidak ada")
	}
}
