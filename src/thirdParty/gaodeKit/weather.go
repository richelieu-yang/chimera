package gaodeKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/reqKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
)

const (
	weatherUrl = "https://restapi.amap.com/v3/weather/weatherInfo"
)

// GetLive 获取"实况"天气.
/*
@param city 城市编码
*/
func (client *Client) GetLive(city string) (*Live, error) {
	_, data, err := reqKit.Get(weatherUrl, map[string][]string{
		"key":        {client.key},
		"city":       {city},
		"extensions": {"base"},
	})
	if err != nil {
		return nil, err
	}

	resp := &WeatherResponse{}
	if err := jsonKit.Unmarshal(data, resp); err != nil {
		return nil, errorKit.Wrap(err, "Fail to unmarshal with json: %s", string(data))
	}
	if err := resp.IsSuccess(); err != nil {
		return nil, err
	}
	if len(resp.Lives) == 0 {
		return nil, errorKit.New("len(resp.Lives) == 0")
	}
	return resp.Lives[0], nil
}

// GetTodayCast 获取今天的"预报"天气.
func (client *Client) GetTodayCast(city string) (*Cast, error) {
	forecast, err := client.GetForecast(city)
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
func (client *Client) GetForecast(city string) (*Forecast, error) {
	_, data, err := reqKit.Get(weatherUrl, map[string][]string{
		"key":        {client.key},
		"city":       {city},
		"extensions": {"all"},
	})
	if err != nil {
		return nil, err
	}

	resp := &WeatherResponse{}
	if err := jsonKit.Unmarshal(data, resp); err != nil {
		return nil, errorKit.Wrap(err, "Fail to unmarshal with json: %s", string(data))
	}
	if err := resp.IsSuccess(); err != nil {
		return nil, err
	}
	if len(resp.Forecasts) == 0 {
		return nil, errorKit.New("len(resp.Forecasts) == 0")
	}
	return resp.Forecasts[0], nil
}
