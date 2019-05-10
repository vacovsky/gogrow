package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vacovsky/gogrow/ggmodels"
)

var (
	apiKey string
	zip    string
	units  string
)

// units=metric
func init() {
	apiKey = os.Getenv("GG_OPENWEATHER_API_KEY")
	units = os.Getenv("GG_OPENWEATHER_API_UNITS")
	zip = os.Getenv("GG_ZIPCODE")
}

func GetWeatherFromAPI() ggmodels.Weather {
	wr := makeRequest()
	w := ggmodels.Weather{}

	w.WindDeg = wr.List[0].Wind.Deg
	w.WindSpeed = wr.List[0].Wind.Speed
	w.TimeStamp = time.Now()
	w.Temperature = wr.List[0].Main.Temp
	w.SeaLevel = wr.List[0].Main.SeaLevel
	w.GroundLevel = wr.List[0].Main.GrndLevel
	w.Clouds = wr.List[0].Clouds.All
	w.Humidity = wr.List[0].Main.Humidity
	w.AirPressure = wr.List[0].Main.Pressure
	w.WeatherIcon = wr.List[0].Weather[0].Icon
	w.Description = wr.List[0].Weather[0].Description

	return w

}

func makeRequest() ggmodels.WeatherRaw {
	weather := ggmodels.WeatherRaw{}
	url := "http://api.openweathermap.org/data/2.5/forecast?zip=" + zip + "&APPID=" + apiKey + "&units=" + units
	fmt.Println(url)
	resp, err := http.Get(url)
	// req.Header.Add("User-Agent", "GreenGuard")
	if err != nil {
		log.Println(err)
	}

	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &weather)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return weather
}
