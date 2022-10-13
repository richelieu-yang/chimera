package main

import (
	"bufio"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	rootPwd := "Cyy7587141200"
	/*
		真实的开始结束时间: time0、time1
		虚假的开始结束时间: timeA、timeB
	*/
	var time0, time1, timeA, timeB time.Time

	logrusKit.InitializeByDefault()

	// time0
	time0 = time.Now()
	logrus.Infof("time0: [%s].", timeKit.FormatTimeToString(time0, timeKit.EntireFormat))

	// timeA
	timeA, err := timeKit.ParseStringToTime(string(timeKit.CommonFormat), "2022-10-01 00:00:00.000")
	if err != nil {
		panic(err)
	}
	logrus.Infof("timeA: [%s].", timeKit.FormatTimeToString(timeA, timeKit.EntireFormat))
	if err := timeKit.SetSystemTime(timeA, rootPwd); err != nil {
		panic(err)
	}
	logrus.Infof("System time is set to timeA(%s).", timeKit.FormatTimeToString(timeA, timeKit.EntireFormat))

	// timeB
	logrus.Info("Please enter text to continue...")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		logrus.Infof("input text: [%s].", text)
		timeB = time.Now()
		logrus.Infof("timeB: [%s].", timeKit.FormatTimeToString(timeB, timeKit.EntireFormat))
		break
	}
	d := timeB.Sub(timeA)
	logrus.Infof("timeB.Sub(timeA): [%s].", timeKit.FormatDurationToString(d))

	// time1
	time1 = time0.Add(d)
	logrus.Infof("time1: [%s].", timeKit.FormatTimeToString(time1, timeKit.EntireFormat))
	if err := timeKit.SetSystemTime(time1, rootPwd); err != nil {
		panic(err)
	}
	logrus.Infof("System time is set to time1(%s).", timeKit.FormatTimeToString(time1, timeKit.EntireFormat))
}
