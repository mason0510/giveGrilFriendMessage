package provider

import (
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"giveGrilFriendMessage/config"
	"io/ioutil"
	"net/http"
)

type DataFetcher interface {
	Fetch() string
}

// 定义具体的实现
type RealDataFetcher struct{}

// 定义获取天气的方法
func (r *RealDataFetcher) Fetch() string {
	type weatherData struct {
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
	}
	apiKey := config.LoadConfig().WetherApi
	city := "London,UK"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var weather weatherData
	if err := json.Unmarshal(data, &weather); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Temperature in %s: %.2f°C\n", city, weather.Main.Temp-273.15)
	fmt.Printf("Humidity in %s: %d%%\n", city, weather.Main.Humidity)
	return "hello"
}

// 定义二者关系
var WireSet = wire.NewSet(
	//绑定具体的实现
	wire.Struct(new(RealDataFetcher), "*"),
	//绑定接口和具体的实现
	wire.Bind(new(DataFetcher), new(*RealDataFetcher)),
)
