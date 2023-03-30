package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func CheckErr(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("data tidak ada")
	}
	return nil
}
