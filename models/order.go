package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	DealerID  uint          `gorm:"not null" json:"dealerID"`
	Dealer    Dealer        `json:"dealer"`
	Delivered bool          `json:"Delivered"`
	Products  []CartOnOrder `json:"Products"`
}

type CartOnOrder struct {
	gorm.Model
	OrderID   uint    `gorm:"not null" json:"OrderID"`
	ProductID uint    `gorm:"not null" json:"ProductID"`
	Product   Product `json:"Products"`
}

func (order *Order) CreateOrder() (*Order, error) {
	err := GetDB().Debug().Create(&order).Error
	if err != nil {
		return &Order{}, err
	}
	return order, nil
}

func (cartonOrder *CartOnOrder) AddToOrder() error {
	err := GetDB().Debug().Create(&cartonOrder).Error
	if err != nil {
		return err
	}
	return nil
}

func (order *Order) FindAll() ([]Order, error) {
	orders := []Order{}
	err := GetDB().Debug().Table("orders").Find(&orders).Error
	if err != nil {
		return []Order{}, err
	}
	if len(orders) > 0 {
		for i, _ := range orders {
			err := GetDB().Debug().Table("cart_on_orders").Where("order_id=?", orders[i].ID).Find(&orders[i].Products).Error
			if err != nil {
				return []Order{}, err
			}
			if len(orders[i].Products) > 0 {
				for a, _ := range orders[i].Products {
					err := GetDB().Debug().Table("products").Where("id=?", orders[i].Products[a].ProductID).Take(&orders[i].Products[a].Product).Error
					if err != nil {
						return []Order{}, err
					}
				}
			}
		}
	}
	return orders, nil
}

func (order *Order) UpdateOrder(id uint) (*Order, error) {
	db := GetDB().Table("orders").Where("id=?", id).UpdateColumn(
		map[string]interface{}{
			"delivered":  order.Delivered,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Order{}, db.Error
	}
	return order, nil
}
