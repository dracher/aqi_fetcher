package libs

import (
	"strings"
	"time"
)

const atLayout = "2006-01-02 15:04:05"

var loc, _ = time.LoadLocation("Asia/Shanghai")

// MyTime is
type MyTime struct {
	time.Time
}

// UnmarshalJSON is
func (at *MyTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		at.Time = time.Time{}
		return
	}
	at.Time, err = time.ParseInLocation(atLayout, s, loc)
	return
}

// IaqiData is
type IaqiData struct {
	Date MyTime `json:"date"`
	Val  int    `json:"val"`
}

// AqiData is
type AqiData struct {
	Aqi struct {
		Date   MyTime `json:"date"`
		Impact string `json:"impact"`
		Val    int    `json:"val"`
	}
	Iaqi struct {
		Co   IaqiData
		No2  IaqiData
		O3   IaqiData
		Pm10 IaqiData
		Pm25 IaqiData
		So2  IaqiData
	}
	Weather struct {
		Tempnow   int
		Temptoday string
	}
	City struct {
		Name string
	}
}

// AqiFlatData is
type AqiFlatData struct {
	Date     string
	Impact   string
	Aqi      int
	Co       int
	No2      int
	O3       int
	Pm10     int
	Pm25     int
	So2      int
	Temp     int
	CityName string
}

// Flat make data flat easy to read and store
func (a AqiData) Flat() *AqiFlatData {
	d := new(AqiFlatData)
	d.Date = a.Aqi.Date.String()
	d.Impact = a.Aqi.Impact
	d.Aqi = a.Aqi.Val
	d.Co = a.Iaqi.Co.Val
	d.No2 = a.Iaqi.No2.Val
	d.O3 = a.Iaqi.O3.Val
	d.Pm10 = a.Iaqi.Pm10.Val
	d.Pm25 = a.Iaqi.Pm25.Val
	d.So2 = a.Iaqi.So2.Val
	d.Temp = a.Weather.Tempnow
	d.CityName = strings.ToLower(a.City.Name)

	return d
}
