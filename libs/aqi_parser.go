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
