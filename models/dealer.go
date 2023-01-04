package models

import (
	"github.com/jinzhu/gorm"
)

type Dealer struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;" json:"DealerName"`
	Email    string `gorm:"size:255;not null; unique" json:"Email"`
	Password string `gorm:"size:255;" json:"Password"`
	Active   bool   `gorm:"size:255;" json:"Active"`
}

type DealerForm struct {
	gorm.Model
	Name                 string `gorm:"size:255;not null;" json:"Name"`
	Phone                string `gorm:"size:255;not null; unique" json:"Phone"`
	Gsm                  string `gorm:"size:255;" json:"Gsm"`
	City                 string `gorm:"size:255;" json:"city"`
	Sistrict             string `gorm:"size:255;" json:"District"`
	Job                  string `gorm:"size:255;" json:"Job"`
	Investment           string `gorm:"size:255;" json:"Investment"`
	DealershipPreference string `gorm:"size:255;" json:"DealershipPreference"`
	CurrentSales         string `gorm:"size:255;" json:"CurrentSales"`
	Experience           string `gorm:"size:255;" json:"Experience"`
	WhyIsItPreferred     string `gorm:"size:255;" json:"WhyIsItPreferred"`
	MostPreferred        string `gorm:"size:255;" json:"MostPreferred"`
	Management           string `gorm:"size:255;" json:"Management"`
	ExploreWhere         string `gorm:"size:255;" json:"ExploreWhere"`
	WhyDoBe              string `gorm:"size:255;" json:"WhyDoBe"`

	Active bool `gorm:"size:255;" json:"Active"`
}

func (dealer *Dealer) Prepare() {
	dealer.Active = true
}

func (dealer *Dealer) CreateDealer() (*Dealer, error) {
	err := GetDB().Debug().Create(&dealer).Error
	if err != nil {
		return &Dealer{}, err
	}
	return dealer, nil
}

func (dealer *Dealer) FindByID(id uint) (*Dealer, error) {
	err := GetDB().Debug().Table("dealers").Where("id=?", id).Take(&dealer).Error
	if err != nil {
		return &Dealer{}, err
	}
	return dealer, nil
}

func (dealer *Dealer) FindByEmail(email string) (*Dealer, error) {
	err := GetDB().Debug().Table("dealers").Where("email=?", email).Take(&dealer).Error
	if err != nil {
		return &Dealer{}, err
	}
	return dealer, nil
}

func (dealer *Dealer) FindAll() ([]Dealer, error) {
	dealers := []Dealer{}
	err := GetDB().Debug().Table("dealers").Where("active=?", true).Find(&dealers).Error
	if err != nil {
		return []Dealer{}, err
	}
	return dealers, nil
}

func (dealer *Dealer) FindDealerRequests() ([]Dealer, error) {
	dealers := []Dealer{}
	err := GetDB().Debug().Table("dealers").Where("active=?", false).Find(&dealers).Error
	if err != nil {
		return []Dealer{}, err
	}
	return dealers, nil
}

// Dealer Forms
func (dealerform *DealerForm) FindAllForms() ([]DealerForm, error) {
	forms := []DealerForm{}
	err := GetDB().Debug().Table("dealerforms").Find(&forms).Error
	if err != nil {
		return []DealerForm{}, err
	}
	return forms, nil
}

func (dealerform *DealerForm) FindByID(id uint) (*DealerForm, error) {
	err := GetDB().Debug().Table("dealerforms").Where("id=?", id).Take(&dealerform).Error
	if err != nil {
		return &DealerForm{}, err
	}
	return dealerform, nil
}

func (dealerdform *DealerForm) CreateDealer() (*DealerForm, error) {
	err := GetDB().Debug().Create(&dealerdform).Error
	if err != nil {
		return &DealerForm{}, err
	}
	return dealerdform, nil
}
