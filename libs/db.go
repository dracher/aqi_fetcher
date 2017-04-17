package libs

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/go-xorm/xorm"
)

// TableData is
type TableData struct {
	Date     string `xorm:"unique"`
	Impact   string
	Aqi      int
	Co       int
	No2      int
	O3       int
	Pm10     int
	Pm25     int
	So2      int
	CityName string `xorm:"-"`
	Temp     int
}

// TableName is
func (a TableData) TableName() string {
	return a.CityName
}

// SaveToDB is
func SaveToDB(aqiData []AqiData, orm *xorm.Engine) {
	for _, aqi := range aqiData {
		row := TableData{}
		row.Date = aqi.Aqi.Date.String()
		row.Impact = aqi.Aqi.Impact
		row.Aqi = aqi.Aqi.Val
		row.Co = aqi.Iaqi.Co.Val
		row.No2 = aqi.Iaqi.No2.Val
		row.O3 = aqi.Iaqi.O3.Val
		row.Pm10 = aqi.Iaqi.Pm10.Val
		row.Pm25 = aqi.Iaqi.Pm25.Val
		row.So2 = aqi.Iaqi.So2.Val
		row.CityName = strings.ToLower(aqi.City.Name)
		row.Temp = aqi.Weather.Tempnow
		r, err := orm.Insert(row)
		log.Info(r, err)
	}
}
