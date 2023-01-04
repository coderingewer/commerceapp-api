package models

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"size:255;not null; unique" json:"UserName"`
	Email    string `gorm:"size:255;not null;unique" json:"Email"`
	Password string `gorm:"size:255;not null;unique" json:"Password"`
}

func (admin *Admin) CreateAdmin() (*Admin, error) {
	err := GetDB().Debug().Create(&admin).Error
	if err != nil {
		return &Admin{}, err
	}

	return admin, nil
}

func (admin Admin) FindByUserName(username string) (Admin, error) {
	err := GetDB().Debug().Table("admins").Where("username=?", username).Take(&admin).Error
	if err != nil {
		return Admin{}, err
	}
	return admin, nil
}

func (admin Admin) FindByID(id uint) (Admin, error) {
	err := GetDB().Debug().Table("admins").Where("id=?", id).Take(&admin).Error
	if err != nil {
		return Admin{}, err
	}
	return admin, nil
}

func (admin *Admin) DeleteAdmin(id uint) error {
	db := GetDB().Debug().Table("admins").Where("id=?", id).Take(&admin).Delete(&Admin{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
