package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/navneetshukl/models"
)

func GetForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.page.tmpl", nil)
}

func GetTemperature(c *gin.Context) {
	lat := c.PostForm("latitude")
	long := c.PostForm("longitude")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error in loading the .env ", err)
		return
	}
	api_key := os.Getenv("WEATHER_API_KEY")
	fmt.Println("Api Key ", api_key)
	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", lat, long, api_key)
	resp, err := http.Get(URL)
	if err != nil {
		log.Println("Error in Getting the data ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in Reading the output ", err)
	}
	var weatherData models.WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Println("Error in unmarshalling the data ", err)
		return
	}
	c.HTML(http.StatusOK, "response.page.tmpl", gin.H{
		"data": weatherData.Main.Temp,
		"area": weatherData.Name,
	})

}
