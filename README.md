[![Go Report Card](https://goreportcard.com/badge/github.com/dbgjerez/iot-devices-crud)](https://goreportcard.com/report/github.com/dbgjerez/iot-devices-crud)

# IOT Devices Management
Application to manage iot devices configuration.

Chips have some many sensor connected which could need some configuration. 

For example, how often the information is sent.

# Build
## Docker
To build as docker image: 
```bash
docker build -t $user/iot-device-manager:$version .
# $docker login (if you're not logged)
docker push 
```
*$user: change with your own Docker Hub user*

*$version: change with yout own microservice version*

# Run
## Golang
```bash
go run app.go
```
*Go to "configuration section" to configure MongoDB connection*

## Docker
In this case, MongoDB and the microservice run with Docker
```bash
# create dedicated network
$ docker network create --driver bridge iot
# run MongoDB container
$ docker run -dit --name iot-mongo -p 27017:27017 --network iot mongo:4.0.22
# run microservice container
$ docker run -dit --env MONGODB_HOST=mongodb://iot-mongo:27017 --name device-crud -p 8080:8080 --network iot b0rr3g0/iot-device-manager:0.4
```
*The example above use my own Docker user. If you want use your own user change b0rr3g0 for it and 0.4 for your own version*

# Configuration
| Variable | Default value | Description |
| ------ | ------ | ------ |
| GIN_MODE | debug | Gin gonic mode. (release for production mode) |
| MONGODB_HOST | "" | MongoDB host url. For example: mongodb://localhost:27017 |
| MONGODB_DEVICE_DB | iot | Mongo database name |
| MONGODB_DEVICE_COLLECTION_NAME | device | Name of collection into de database |

# Architecture and libraries
* MongoDB: NoSQL database used to store the information about chips and sensors.
* Gin-gonic: Framework used to develop the microservice.
* GoDotEnv: library that facilitates working with enviroments variables
* mongo-driver: library to work with MongoDB
* mongo-go-pagination: library helper to paginate MongoDB queries
