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

func GetAllDealer(c *gin.Context) {
	dealer := models.Dealer{}

	dealers, err := dealer.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealers)
}

func GetAllDealerRequests(c *gin.Context) {
	dealer := models.Dealer{}

	dealers, err := dealer.FindDealerRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealers)
}
func GetAllDealerForms(c *gin.Context) {
	dealerform := models.DealerForm{}

	dealerforms, err := dealerform.FindAllForms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealerforms)
}

func GetFormByID(c *gin.Context) {
	dealerform := models.DealerForm{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}

	dealerformm, err := dealerform.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealerformm)
}

func GetByID(c *gin.Context) {
	dealer := models.Dealer{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "geçersiz istek"})
		return
	}

	dealers, err := dealer.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealers)
}

func CreateDealer(c *gin.Context) {
	dealer := models.Dealer{}
	admin := models.Admin{}
	aid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hata": "ID çekilelemedi"})
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
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	err = json.Unmarshal(body, &dealer)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	dealerr, err := dealer.CreateDealer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Veriler alınırken sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealerr)

}

func CreateDealerForm(c *gin.Context) {
	dealerform := models.DealerForm{}
	dealerformm, err := dealerform.CreateDealer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"hata": "Sunucu hatası yaşandı"})
		return
	}
	c.JSON(http.StatusOK, dealerformm)

}

func DealerLogin(c *gin.Context) {
	dealer := models.Dealer{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	err = json.Unmarshal(body, &dealer)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity})
		return
	}
	dealerr, err := dealer.FindByEmail(dealer.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	if dealer.Password != dealerr.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"ERROR": "Şifre Hatalı"})
		return
	}
	token, err := auth.CreateToken(dealerr.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"hata": "Token verilirken hata oluştu"})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
