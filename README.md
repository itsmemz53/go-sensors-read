# Read Sensors Data REST API using Gin and Gorm

How to use:

```
$ go run main.go
```

POST Sensor Data
```
POST /sensor

Body {"name": "Sensor 3", "mac_address": "3"}
```

GET Sensor Data
```
GET /sensors

Response   
{
    "data": [
        {
            "id": 1,
            "name": "Sensor 1",
            "mac_address": "1"
        },
        {
            "id": 2,
            "name": "Sensor 2",
            "mac_address": "2"
        },
        {
            "id": 3,
            "name": "Sensor 3",
            "mac_address": "3"
        }
    ]
}
```

POST Sensors Reading
```
POST /sensor-reading

Body {"request_id": 3, "sensor_id": 1, "data": "some data"}
```

GET Sensors Data
```
GET /sensors/1/

Response
{
    "data": [
        {
            "request_id": 1,
            "sensor_id": 1,
            "data": "some data"
        },
        {
            "request_id": 2,
            "sensor_id": 1,
            "data": "some data"
        },
        {
            "request_id": 3,
            "sensor_id": 1,
            "data": "some data"
        }
    ]
}
```


Why Go?
https://www.toptal.com/back-end/server-side-io-performance-node-php-java-go
https://medium.com/innovation-incubator/rest-api-performance-comparison-python-vs-golang-dc4decbd0543
