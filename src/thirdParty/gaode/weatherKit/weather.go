package weatherKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/reqKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/thirdParty/gaode/gaodeKit"
)

const (
	Url = "https://restapi.amap.com/v3/weather/weatherInfo"
)

// GetLiveWeather 获取"实况"天气.
/*
@param city 城市编码
*/
func GetLiveWeather(city string) (*Live, error) {
	apiKey, err := gaodeKit.GetApiKey()
	if err != nil {
		return nil, err
	}

	_, data, err := reqKit.Get(Url, map[string][]string{
		"key":        {apiKey},
		"city":       {city},
		"extensions": {"base"},
	})
	if err != nil {
		return nil, err
	}

	resp := &GaodeResponse{}
	if err := jsonKit.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	if err := resp.Check(); err != nil {
		return nil, err
	}
	if len(resp.Lives) == 0 {
		return nil, errorKit.New("length of lives is zero")
	}
	return resp.Lives[0], nil
}

//// GetForecastWeather 获取"预报"天气.
///*
//@param city 城市编码
//*/
//func GetForecastWeather(city string) ([]*Forecast, error) {
//
//}
