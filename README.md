# IOT Devices Management
Application to manage iot devices configuration.

Chips have some many sensor connected which could need some configuration. 

For example, how often the information is sent.

# Architecture and libraries
* MongoDB: NoSQL database used to store the information about chips and sensors.
* Gin-gonic: Framework used to develop the microservice.
* GoDotEnv: library that facilitates working with enviroments variables

# Configuration
| Variable | Default value | Description |
| ------ | ------ | ------ |
| GIN_MODE | debug | Gin gonic mode. (release for production mode) |