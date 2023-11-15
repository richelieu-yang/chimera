package weatherKit

import "github.com/richelieu-yang/chimera/v2/src/component/web/reqKit"

const (
	Url = "https://restapi.amap.com/v3/weather/weatherInfo"
)

// GetLiveWeather 获取"实况"天气.
/*
@param city 城市编码
*/
func GetLiveWeather(city string) (*Live, error) {
	reqKit.GetDefaultClient()
}

// GetForecastWeather 获取"预报"天气.
/*
@param city 城市编码
*/
func GetForecastWeather(city string) ([]*Forecast, error) {

}
