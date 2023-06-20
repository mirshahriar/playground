package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result []User

	// Panic for custom scanner
	if err := DB.Find(&result).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// The following one is also a panic

	//if err := DB.Where("1 = 1").Find(&result).Error; err != nil {
	//	t.Errorf("Failed, got error: %v", err)
	//}

	// The following one is working fine.

	//if err := DB.Where("1 = ?", 1).Find(&result).Error; err != nil {
	//	t.Errorf("Failed, got error: %v", err)
	//}
}
