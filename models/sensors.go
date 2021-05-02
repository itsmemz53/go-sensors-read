package models

type Sensors struct {
	Id     int64   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	MacAddress string `json:"mac_address"`
}


type Readings struct {
	Sensor  *Sensors `json:"sensor"`
	RequestId  int64 `json:"requestId" gorm:"primary_key"`
	Data 	string `json:"data"`
}