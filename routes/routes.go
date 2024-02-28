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

// GetForm function is used to get the form page for entering latitude and longitude
func GetForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.page.tmpl", gin.H{
		"error": false,
	})
}

// GetTemperature function is used to get the weather data for given latitude and longitude
func GetTemperature(c *gin.Context) {
	lat := c.PostForm("latitude")
	long := c.PostForm("longitude")
	if len(lat) == 0 || len(long) == 0 {
		c.HTML(http.StatusOK, "form.page.tmpl", gin.H{
			"error": true,
		})
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("Error in loading the .env ", err)
		return
	}
	api_key := os.Getenv("WEATHER_API_KEY")
	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", lat, long, api_key)
	resp, err := http.Get(URL)
	if err != nil {
		log.Println("Error in Getting the data ", err)
		c.HTML(http.StatusOK, "form.page.tmpl", gin.H{
			"error": true,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in Reading the output ", err)
		c.HTML(http.StatusOK, "form.page.tmpl", gin.H{
			"error": true,
		})
		return
	}
	var weatherData models.WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Println("Error in unmarshalling the data ", err)
		c.HTML(http.StatusOK, "form.page.tmpl", gin.H{
			"error": true,
		})
		return

	}
	c.HTML(http.StatusOK, "response.page.tmpl", gin.H{
		"temp":     weatherData.Main.Temp,
		"pressure": weatherData.Main.Pressure,
		"humidity": weatherData.Main.Humidity,
		"wind":     weatherData.Wind.Speed,
		"area":     weatherData.Name,
	})

}
