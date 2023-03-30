package cli

import (
	"fmt"

	"enigmacamp.camp/gorm-sinar-harapan-makmur/config"
	"enigmacamp.camp/gorm-sinar-harapan-makmur/model"
	"enigmacamp.camp/gorm-sinar-harapan-makmur/repository"
	"enigmacamp.camp/gorm-sinar-harapan-makmur/usecase"
	"gorm.io/gorm"
)

func VehicleCLI() {
	db := createConnection()
	vehicleRepo := repository.NewVehicleRepository(db)
	vehicleUseCase := usecase.NewVehicleUseCase(vehicleRepo)

	vech := model.Vehicle{
		Brand:          "Mazda",
		Model:          "CX-5",
		ProductionYear: 2019,
		Color:          "Red",
		IsAutomatic:    true,
		Stock:          3,
		SalePrice:      590000000,
		Status:         "Bekas",
	}

	if err := vehicleUseCase.RegisterNewVehicle(vech); err != nil {
		panic(err)
	}

}

func createConnection() *gorm.DB {
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

	return newDbConn.Conn()
}
