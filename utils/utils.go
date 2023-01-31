package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

func (utils *Utils) GetConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./config")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic("config err: " + err.Error())
	}
	return config
}

func (utils *Utils) ReturnAPIResult(c *gin.Context, success bool, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
	})
}
