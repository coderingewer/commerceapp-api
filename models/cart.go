package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	DealerID  uint    `gorm:"not null" json:"DealerID"`
	ProductID uint    `gorm:"not null" json:"ProductID"`
	Unit      int     `gorm:"not null; unique" json:"Unit"`
	Product   Product `json:"Product"`
}

func (cart *Cart) CreateCart() (*Cart, error) {
	err := GetDB().Debug().Create(&cart).Error
	if err != nil {
		return &Cart{}, err
	}
	err = GetDB().Debug().Table("products").Where("id=?", cart.ProductID).Take(&cart.Product).Error
	if err != nil {
		return &Cart{}, err
	}
	return cart, nil
}

func (cart *Cart) FinByDealerID(did uint) ([]Cart, error) {
	carts := []Cart{}
	err := GetDB().Debug().Table("carts").Where("dealer_id=?", did).Find(&carts).Error
	if len(carts) > 0 {
		for i, _ := range carts {
			err = GetDB().Debug().Table("products").Where("id=?", carts[i].ProductID).Take(&carts[i].Product).Error
			if err != nil {
				return []Cart{}, err
			}
		}
	}
	if err != nil {
		return []Cart{}, err
	}
	return carts, nil
}

func (cart *Cart) DeleteCart(id uint) error {
	db := GetDB().Debug().Table("carts").Where("id=?", id).Take(&cart).Delete(&Cart{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
