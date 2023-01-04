package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/auth"
	"go.mod/models"
)

func UploadProductImageImage(c *gin.Context) {
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
	image := models.Image{}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "Dosya alınırken hata oluştu"})
		return
	}
	defer file.Close()
	uploadUrl, err := models.NewMediaUpload().FileUpload(models.File{File: file})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "Dosya alınırken hata oluştu"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}
	image.Url = uploadUrl
	imagee, err := image.SaveImage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": fmt.Sprintf("Dosya kaydedilirken hata oluştu %s", err)})
		return
	}
	productimg := models.ProductImage{}
	productimg.ImageID = imagee.ID
	productimg.ProductID = uint(id)
	err = productimg.SaveProductImage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": fmt.Sprintf("Dosya kaydedilirken hata oluştu %s", err)})
		return
	}

	c.JSON(http.StatusOK, imagee)
}
