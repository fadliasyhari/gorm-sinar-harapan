package repository

import (
	"enigmacamp.camp/gorm-sinar-harapan-makmur/model"
	"enigmacamp.camp/gorm-sinar-harapan-makmur/utils"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	BaseRepository[model.Vehicle]
}

type vehicleRepo struct {
	db *gorm.DB
}

func (v *vehicleRepo) Create(newData model.Vehicle) error {
	result := v.db.Debug().Create(newData)
	err := utils.CheckErr(result)
	if err != nil {
		return err
	}
	return nil
}

func (v *vehicleRepo) Update(newData model.Vehicle) error {
	result := v.db.Model(&model.Vehicle{}).Debug().Where("id = ?", newData.ID).Updates(newData)
	err := utils.CheckErr(result)
	if err != nil {
		return err
	}
	return nil
}

func (v *vehicleRepo) Delete(id string) error {
	result := v.db.Delete(&model.Vehicle{}, "id = ?", id)
	err := utils.CheckErr(result)
	if err != nil {
		return err
	}
	return nil
}

func (v *vehicleRepo) List() ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	result := v.db.Debug().Find(&vehicles)
	err := utils.CheckErr(result)
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (v *vehicleRepo) Get(id string) (model.Vehicle, error) {
	var vehicle model.Vehicle
	result := v.db.Debug().Find(&vehicle)
	err := utils.CheckErr(result)
	if err != nil {
		return model.Vehicle{}, err
	}
	return vehicle, nil
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepo{
		db: db,
	}
}
