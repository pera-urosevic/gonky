package widget

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/pera-urosevic/gonky/util"
)

// Weather //
type Weather struct {
	Key    string
	Lat    string
	Lon    string
	url    string
	Every  time.Duration
	Render func(*Weather)
	State  WeatherBitData
}

// WeatherBit //
type WeatherBit struct {
	Count int              `json:"count"`
	Data  []WeatherBitData `json:"data"`
}

// WeatherBitData //
type WeatherBitData struct {
	AppTemp      float64               `json:"app_temp"`
	Aqi          float64               `json:"aqi"`
	CityName     string                `json:"city_name"`
	Clouds       float64               `json:"clouds"`
	CountryCode  string                `json:"country_code"`
	Datetime     string                `json:"datetime"`
	Dewpt        float64               `json:"dewpt"`
	Dhi          float64               `json:"dhi"`
	Dni          float64               `json:"dni"`
	ElevAngle    float64               `json:"elev_angle"`
	Ghi          float64               `json:"ghi"`
	HAngle       float64               `json:"h_angle"`
	LastObTime   string                `json:"last_ob_time"`
	Lat          float64               `json:"lat"`
	Lon          float64               `json:"lon"`
	ObTime       string                `json:"ob_time"`
	Pod          string                `json:"pod"`
	Precip       float64               `json:"precip"`
	Pres         float64               `json:"pres"`
	Rh           float64               `json:"rh"`
	Slp          float64               `json:"slp"`
	Snow         float64               `json:"snow"`
	SolarRad     float64               `json:"solar_rad"`
	StateCode    string                `json:"state_code"`
	Station      string                `json:"station"`
	Sunrise      string                `json:"sunrise"`
	Sunset       string                `json:"sunset"`
	Temp         float64               `json:"temp"`
	Timezone     string                `json:"timezone"`
	Ts           float64               `json:"ts"`
	Uv           float64               `json:"uv"`
	Vis          float64               `json:"vis"`
	Weather      WeatherBitDataWeather `json:"weather"`
	WindCdir     string                `json:"wind_cdir"`
	WindCdirFull string                `json:"wind_cdir_full"`
	WindDir      float64               `json:"wind_dir"`
	WindSpd      float64               `json:"wind_spd"`
}

// WeatherBitDataWeather //
type WeatherBitDataWeather struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Start //
func (weather Weather) Start() {
	weather.url = fmt.Sprintf("https://api.weatherbit.io/v2.0/current?units=M&key=%s&lat=%s&lon=%s", weather.Key, weather.Lat, weather.Lon)
	go runner.Run(weather.Every, weather.tick)
}

func (weather *Weather) tick() {
	weather.sensor()
	weather.Render(weather)
}

func (weather *Weather) sensor() {
	if util.Debug {
		weather.State = WeatherBitData{
			Weather: WeatherBitDataWeather{
				Description: "Few clouds",
			},
			Temp:    27,
			AppTemp: 28,
			Rh:      61,
			Precip:  0,
			Pres:    1005.4,
			WindSpd: 5.7,
		}
	} else {
		client := http.Client{}
		req, e := http.NewRequest(http.MethodGet, weather.url, nil)
		if e != nil {
			util.ErrorLog(e)
			return
		}
		res, e := client.Do(req)
		if e != nil {
			util.ErrorLog(e)
			return
		}
		body, e := ioutil.ReadAll(res.Body)
		response := WeatherBit{}
		e = json.Unmarshal(body, &response)
		if e != nil {
			util.ErrorLog(e)
			return
		}
		if len(response.Data) < 1 {
			util.Log(response)
			e = errors.New("Weather received empty data")
			return
		}
		weather.State = response.Data[0]
	}
}
