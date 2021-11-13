package handlers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yigitaltunay/go-assignment/app/models"
	"github.com/yigitaltunay/go-assignment/app/repositories"
	"github.com/yigitaltunay/go-assignment/app/services"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	log.Println("yigit")
}

func Currency(c *gin.Context) {
	currencyType := c.Param("type")
	err := repositories.R.IsExistsCurrentType(currencyType)
	if err != nil {
		log.Println("IsExistsCurrentType")
		c.JSON(404, err)
		return
	}
	// Cache control for 10 minutes
	val, err := repositories.R.MemoryDbFetchWithKey(currencyType)
	if err == nil {
		c.JSON(200, val.Value)
		return
	}
	// Get data from providers
	resp, err := services.FindCurrencyCurrencies(currencyType)
	if err != nil {
		c.JSON(400, resp)
		return
	}
	repositories.R.MemoryDbInsert(models.InMemoryModel{Key: currencyType, Value: resp, CreateDate: time.Now()})

	c.JSON(200, resp)

}
