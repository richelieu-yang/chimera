package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/tableKit"
	"github.com/richelieu-yang/chimera/v2/src/thirdParty/gaode/gaodeKit"
	"github.com/richelieu-yang/chimera/v2/src/thirdParty/gaode/weatherKit"
	"github.com/sirupsen/logrus"
)

func main() {
	gaodeKit.MustSetUp("b15c36bf1df4c272e92f3f1875a127f1")

	table := tableKit.CreateTable1()
	table.AddHeaders("地区", "当前天气", "白天天气", "夜晚天气")

	cities := []string{"320200", "320206", "320211", "320213", "310000"}
	for _, city := range cities {
		live, err := weatherKit.GetLive(city)
		if err != nil {
			logrus.WithError(err).WithField("city", city).Error("Fail to get live")
			continue
		}
		todayCast, err := weatherKit.GetTodayCast(city)
		if err != nil {
			logrus.WithError(err).WithField("city", city).Error("Fail to get today cast")
			continue
		}

		table.AddRow(live.City, fmt.Sprintf("%s(%s)",
			live.Temperature, live.Weather),
			fmt.Sprintf("%s(%s)", todayCast.DayTemp, todayCast.DayWeather),
			fmt.Sprintf("%s(%s)", todayCast.NightTemp, todayCast.NightWeather))
	}

	fmt.Println(timeKit.FormatCurrent(timeKit.FormatEntire))
	fmt.Println(table.Render())
}
