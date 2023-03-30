package usecase

import (
	"fmt"

	"enigmacamp.camp/gorm-sinar-harapan-makmur/model"
	"enigmacamp.camp/gorm-sinar-harapan-makmur/repository"
)

type VehicleUseCase interface {
	RegisterNewVehicle(newVehicle model.Vehicle) error
	FindAllVehicle() ([]model.Vehicle, error)
	GetVehicle(id string) (model.Vehicle, error)
	UpdateVehicle(newVehicle model.Vehicle) error
	Delete(id string) error
}

type vehicleUseCase struct {
	vehicleRepo repository.VehicleRepository
}

func (v *vehicleUseCase) RegisterNewVehicle(newVehicle model.Vehicle) error {

	errValidation := vehicleValidation(newVehicle)
	if errValidation != nil {
		return fmt.Errorf("failed to create new vehicle : %v", errValidation)
	}

	err := v.vehicleRepo.Create(newVehicle)
	if err != nil {
		return fmt.Errorf("failed to create new vehicle : %v", err)
	}
	return nil
}

func (v *vehicleUseCase) UpdateVehicle(newVehicle model.Vehicle) error {
	errValidation := vehicleValidation(newVehicle)
	if errValidation != nil {
		return fmt.Errorf("failed to update new vehicle : %v", errValidation)
	}

	err := v.vehicleRepo.Update(newVehicle)
	if err != nil {
		return fmt.Errorf("failed to update new vehicle : %v", err)
	}
	return nil
}

func (v *vehicleUseCase) FindAllVehicle() ([]model.Vehicle, error) {
	return v.vehicleRepo.List()
}

func (v *vehicleUseCase) GetVehicle(id string) (model.Vehicle, error) {
	return v.vehicleRepo.Get(id)
}

func (v *vehicleUseCase) Delete(id string) error {
	return v.vehicleRepo.Delete(id)
}

func vehicleValidation(payload model.Vehicle) error {
	if payload.Brand == "" || payload.Model == "" || payload.Color == "" {
		return fmt.Errorf("brand, model, and color are required")
	}

	if !payload.IsValidStatus() {
		return fmt.Errorf("invalid status : %s", payload.Status)
	}

	if payload.SalePrice <= 0 {
		return fmt.Errorf("sale price can't zero or negative")
	}

	if payload.Stock < 0 {
		return fmt.Errorf("stock can't negative")
	}
	return nil
}

func NewVehicleUseCase(vehicleRepo repository.VehicleRepository) VehicleUseCase {
	return &vehicleUseCase{
		vehicleRepo: vehicleRepo,
	}
}
