package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/auth"
	"go.mod/models"
)

func CreteOrder(c *gin.Context) {
	dealer := models.Dealer{}
	did, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
		return
	}
	dealerr, err := dealer.FindByID(did)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "Böyle bir bayii yok"})
		return
	}
	if dealerr.ID != did {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "Buna yetkiniz yok"})
		return
	}

	order := models.Order{}

	order.DealerID = dealer.ID
	order.Delivered = false

	orderr, err := order.CreateOrder()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Sunucu hatası"})
		return
	}

	cart := models.Cart{}

	carts, err := cart.FinByDealerID(dealer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Sunucu hatası"})
		return
	}
	if len(carts) > 0 {
		for i, _ := range carts {
			cartInOrder := models.CartOnOrder{}
			cartInOrder.ProductID = carts[i].ProductID
			cartInOrder.OrderID = orderr.ID
			cartInOrder.AddToOrder()
			cart.DeleteCart(carts[i].ID)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "Sepet boş"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func UpdateOrder(c *gin.Context) {
	order := models.Order{}
	admin := models.Admin{}
	aid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
		return
	}
	admn, err := admin.FindByID(aid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "Böyle bir admin yok"})
		return
	}
	if admn.ID != aid {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	err = json.Unmarshal(body, &order)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}
	dealers, err := order.UpdateOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealers)

}

func GetAllOrders(c *gin.Context) {
	order := models.Order{}
	admin := models.Admin{}
	aid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
		return
	}
	fmt.Println(aid)
	admn, err := admin.FindByID(uint(aid))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "Böyle bir admin yok"})
		return
	}
	if admn.ID != aid {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
		return
	}
	dealers, err := order.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealers)
}
