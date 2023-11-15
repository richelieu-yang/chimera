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

// GetLive 获取"实况"天气.
/*
@param city 城市编码
*/
func GetLive(city string) (*Live, error) {
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
		return nil, errorKit.New("len(resp.Lives) == 0")
	}
	return resp.Lives[0], nil
}

// GetTodayCast 获取今天的"预报"天气.
func GetTodayCast(city string) (*Cast, error) {
	forecast, err := GetForecast(city)
	if err != nil {
		return nil, err
	}

	if len(forecast.Casts) == 0 {
		return nil, errorKit.New("len(forecast.Casts) == 0")
	}

	return forecast.Casts[0], nil
}

// GetForecast 获取"预报"天气.
/*
@param city 城市编码
*/
func GetForecast(city string) (*Forecast, error) {
	apiKey, err := gaodeKit.GetApiKey()
	if err != nil {
		return nil, err
	}

	_, data, err := reqKit.Get(Url, map[string][]string{
		"key":        {apiKey},
		"city":       {city},
		"extensions": {"all"},
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
	if len(resp.Forecasts) == 0 {
		return nil, errorKit.New("len(resp.Forecasts) == 0")
	}
	return resp.Forecasts[0], nil
}
