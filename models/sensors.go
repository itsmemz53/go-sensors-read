package models

type Sensors struct {
	Id     int64   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	MacAddress string `json:"mac_address"`
}


type Readings struct {
	RequestId  int64 `json:"request_id" gorm:"primary_key"`
	SensorId  int64 `json:"sensor_id" sql:"type:int64 REFERENCES sensors(id)"`
	Data 	string `json:"data"`
}