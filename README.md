# IOT Devices Management
Application to manage iot devices configuration.

Chips have some many sensor connected which could need some configuration. 

For example, how often the information is sent.

# Architecture and libraries
* MongoDB: NoSQL database used to store the information about chips and sensors.
* Gin-gonic: Framework used to develop the microservice.
* GoDotEnv: library that facilitates working with enviroments variables
* mongo-driver: library to work with MongoDB

# Configuration
| Variable | Default value | Description |
| ------ | ------ | ------ |
| GIN_MODE | debug | Gin gonic mode. (release for production mode) |
| MONGODB_HOST | "" | MongoDB host url. For example: mongodb://localhost:27017 |
| MONGODB_DEVICE_DB | "" | Mongo database name |
| MONGODB_DEVICE_COLLECTION_NAME | "" | Name of collection into de database |