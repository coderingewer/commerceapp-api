package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name     string `gorm:"size:255;not null; unique" json:"UserName"`
	Email    string `gorm:"size:255;not null;unique" json:"Email"`
	Phone    string `gorm:"size:255;not null;unique" json:"Phone"`
	Message  string `gorm:"size:255;not null;unique" json:"Message"`
	Answered bool   `gorm:"size:255;not null;unique" json:"Answered"`
}

func (contact *Contact) CreateContact() (*Contact, error) {
	err := GetDB().Debug().Create(&contact).Error
	if err != nil {
		return &Contact{}, err
	}

	return contact, nil
}
func (contact Contact) FindForms() ([]Contact, error) {
	contacts := []Contact{}
	err := GetDB().Debug().Table("contacts").Where("answered=?", false).Find(&contacts).Error
	if err != nil {
		return []Contact{}, err
	}
	return contacts, nil
}

func (contact Contact) FindAnswered() ([]Contact, error) {
	contacts := []Contact{}
	err := GetDB().Debug().Table("contacts").Where("answered=?", true).Find(&contacts).Error
	if err != nil {
		return []Contact{}, err
	}
	return contacts, nil
}

func (contact Contact) FindByID(id uint) (Contact, error) {
	err := GetDB().Debug().Table("contacts").Where("id=?", id).Take(&contact).Error
	if err != nil {
		return Contact{}, err
	}
	return contact, nil
}

func (contact *Contact) DeleteContact(id uint) error {
	db := GetDB().Debug().Table("contacts").Where("id=?", id).Take(&contact).Delete(&Contact{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (contact *Contact) Update(id uint) (*Contact, error) {
	db := GetDB().Table("dealers").Where("id=?", id).UpdateColumn(
		map[string]interface{}{
			"answered":   contact.Answered,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Contact{}, db.Error
	}
	err := GetDB().Table("contacts").Where("id=?", id).Take(&contact).Error
	if err != nil {
		return &Contact{}, err
	}
	return contact, nil
}
