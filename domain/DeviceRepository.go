package domain

import (
	"context"
	"iot-devices-crud/config"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var server = os.Getenv("MONGODB_HOST")
var dbName = os.Getenv("MONGODB_DEVICE_DB")
var collectionName = os.Getenv("MONGODB_DEVICE_COLLECTION_NAME")
var client = config.NewConnection(server)

const id string = "_id"

type DeviceRepository struct {
}

func (repository *DeviceRepository) FindById(idDevice string) *Device {
	ctx, cancel := context.WithTimeout(context.Background(), config.MONGODB_TIMEOUT_SEC*time.Second)
	defer cancel()
	query := bson.D{primitive.E{Key: id, Value: idDevice}}
	var device Device
	err := client.Database(dbName).Collection(collectionName).FindOne(ctx, query).Decode(&device)
	if err != nil {
		return nil
	}
	return &device
}
