package util

import (
	"time"
	"fmt"
	"strconv"
	"math"
)

func ToDateTimeString(year int,month int,day int,hour int,minute int,second int) string {
	var strMonth,strDay string
	strMonth = strconv.Itoa(month)
	if month <= 9 {
		strMonth = "0" + strconv.Itoa(int(month))
	}
	strDay = strconv.Itoa(day)
	if day <= 9 {
		strDay = "0" + strconv.Itoa(day)
	}


	var strHour,strMinute,strSecond string
	if hour <= 9 {
		strHour = "0" + strconv.Itoa(hour)
	}else{
		strHour = strconv.Itoa(hour)
	}

	if minute <= 9 {
		strMinute = "0" + strconv.Itoa(minute)
	}else {
		strMinute = strconv.Itoa(minute)
	}

	if second <= 9 {
		strSecond = "0" + strconv.Itoa(second)
	}else {
		strSecond = strconv.Itoa(second)
	}

	return fmt.Sprintf("%d-%s-%s %s:%s:%s",year,strMonth,strDay,strHour,strMinute,strSecond)
}

func ToLocaleDateTimeString() string{
	now := time.Now()
	year,month,day := now.Date()

	var strMonth,strDay string
	if month <= 9 {
		strMonth = "0" + strconv.Itoa(int(month))
	}else {
		strMonth =  strconv.Itoa(int(month))
	}
	if day <= 9 {
		strDay = "0" + strconv.Itoa(day)
	}else {
		strDay = strconv.Itoa(day)
	}

	hour := now.Hour()
	var strHour,strMinute,strSecond string
	if hour <= 9 {
		strHour = "0" + strconv.Itoa(hour)
	}else {
		strHour = strconv.Itoa(hour)
	}
	minute := now.Minute()
	if minute <= 9 {
		strMinute = "0" + strconv.Itoa(minute)
	}else {
		strMinute = strconv.Itoa(minute)
	}
	second := now.Second()
	if second <= 9 {
		strSecond = "0" + strconv.Itoa(second)
	}else {
		strSecond = strconv.Itoa(second)
	}

	return fmt.Sprintf("%d-%s-%s %s:%s:%s",year,strMonth,strDay,strHour,strMinute,strSecond)
}

func TimeSecDiff(one string,another string) int64 {

	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, err := time.LoadLocation("Asia/Shanghai")                            //重要：获取时区
	if err != nil {
		panic(err)
	}
	tA, _ := time.ParseInLocation(timeLayout, one, loc)
	A := tA.Unix()

	tB, _ := time.ParseInLocation(timeLayout, another, loc)
	B := tB.Unix()

	return int64(math.Abs(float64(A - B)))
}

func ToLocaleDateString() string {
	now := time.Now()
	year,month,day := now.Date()

	var strMonth,strDay string
	if month <= 9 {
		strMonth = "0" + strconv.Itoa(int(month))
	}else {
		strMonth =  strconv.Itoa(int(month))
	}
	if day <= 9 {
		strDay = "0" + strconv.Itoa(day)
	}else {
		strDay = strconv.Itoa(day)
	}

	return fmt.Sprintf("%d-%s-%s",year,strMonth,strDay)
}