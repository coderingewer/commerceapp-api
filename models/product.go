package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ModelName      string `gorm:"not null;" json:"ModelName"`
	KnittingModel  string `gorm:"not null;" json:"KnittingModel"`
	Code           string `gorm:"not null;" json:"Code"`
	Detail         string `gorm:"not null;" json:"Detail"`
	Content        string `gorm:"not null;" json:"Content"`
	Pieces         int    `gorm:"not null;" json:"Pieces"`
	Measures       string `gorm:"not null;" json:"Measures"`
	MetalColours   string `gorm:"not null;" json:"MetalColours"`
	RattanColours  string `gorm:"not null;" json:"RattanColours"`
	CushionColours string `gorm:"not null;" json:"CushionColours"`
	Weight         string `gorm:"not null;" json:"Weight"`

	Price  int            `gorm:"not null; unique" json:"Price"`
	Images []ProductImage `json:"Images"`
}

func (product *Product) CreateProduct() (*Product, error) {
	err := GetDB().Debug().Create(&product).Error
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (product *Product) FindByID(id uint) (*Product, error) {
	err := GetDB().Debug().Table("products").Where("id=?", id).Take(&product).Error
	if err != nil {
		return &Product{}, err
	}
	err = GetDB().Debug().Table("product_images").Where("product_id=?", product.ID).Find(&product.Images).Error
	if err != nil {
		return &Product{}, err
	}
	if len(product.Images) > 0 {
		for a, _ := range product.Images {
			err := GetDB().Debug().Table("images").Where("id=?", product.Images[a].ID).Take(&product.Images[a].Image).Error
			if err != nil {
				return &Product{}, err
			}
		}
	}

	return product, nil
}

func (product *Product) FindAll() (*[]Product, error) {
	products := []Product{}
	err := GetDB().Debug().Table("products").Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	if len(products) > 0 {
		for i, _ := range products {
			err := GetDB().Debug().Table("product_images").Where("product_id=?", products[i].ID).Find(&products[i].Images).Error
			if err != nil {
				return &[]Product{}, err
			}
			if len(products[i].Images) > 0 {
				for a, _ := range products[i].Images {
					err := GetDB().Debug().Table("images").Where("id=?", products[i].Images[a].ID).Take(&products[i].Images[a].Image).Error
					if err != nil {
						return &[]Product{}, err
					}
				}
			}
		}
	}
	return &products, nil
}

func (product *Product) DeleteByID(id uint) error {
	db := GetDB().Debug().Table("products").Where("id=?", id).Take(&product).Delete(&Product{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
