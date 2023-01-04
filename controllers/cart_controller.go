package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/models"
)

func CreateCart(c *gin.Context) {
	cart := models.Cart{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	err = json.Unmarshal(body, &cart)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	cart.ProductID = uint(id)
	cartCreated, err := cart.CreateCart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartCreated)
}

func FinCartByDealerID(c *gin.Context) {
	cart := models.Cart{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}

	carts, err := cart.FinByDealerID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, carts)
}

func DeleteCartByID(c *gin.Context) {
	cart := models.Cart{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}

	err = cart.DeleteCart(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, nil)
}
