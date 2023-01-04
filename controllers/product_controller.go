package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/auth"
	"go.mod/models"
)

func CreateProduct(c *gin.Context) {
	admin := models.Admin{}
	aid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
		return
	}
	admn, err := admin.FindByID(aid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "Böyle bir admin yok çekilelemedi"})
		return
	}
	if admn.ID != aid {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
		return
	}
	product := models.Product{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	err = json.Unmarshal(body, &product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	productCreated, err := product.CreateProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productCreated)
}
func GetProductByID(c *gin.Context) {
	product := models.Product{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}

	productt, err := product.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, productt)
}

func GetAllProducts(c *gin.Context) {
	product := models.Product{}
	products, err := product.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func DeleteProductByID(c *gin.Context) {
	product := models.Product{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}

	err = product.DeleteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, nil)
}
