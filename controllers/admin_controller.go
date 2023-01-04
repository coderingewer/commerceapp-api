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

func CreateAdmin(c *gin.Context) {
	admin := models.Admin{}
	fmt.Println(c.Request.Header)

	aid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": err})
		return
	}
	fmt.Println(aid)
	admn, err := admin.FindByID(aid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "Böyle bir admin yok çekilelemedi"})
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
	err = json.Unmarshal(body, &admin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	adminCreated, err := admin.CreateAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adminCreated)
}

func Login(c *gin.Context) {
	admin := models.Admin{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	admn, err := admin.FindByUserName(admin.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	if admin.Password != admn.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"ERROR": "Şifre Hatalı"})
		return
	}
	token, err := auth.CreateToken(admn.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "Token verilirken hata oluştu"})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func DeleteAdminByID(c *gin.Context) {
	_, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
		return
	}
	admin := models.Admin{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}
	err = admin.DeleteAdmin(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"durum": "admin silindi"})
}
