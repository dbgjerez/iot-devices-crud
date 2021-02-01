package domain

import (
	"context"
	"iot-devices-crud/config"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var server = os.Getenv("MONGODB_HOST")
var dbName = os.Getenv("MONGODB_DEVICE_DB")
var collectionName = os.Getenv("MONGODB_DEVICE_COLLECTION")
var client = config.NewConnection(server)
var collection = client.Database(dbName).Collection(collectionName)

const id string = "_id"

type DeviceRepository struct {
}

func (repository *DeviceRepository) FindById(idDevice string) *Device {
	ctx, cancel := context.WithTimeout(context.Background(), config.MONGODB_TIMEOUT_SEC*time.Second)
	defer cancel()
	query := bson.D{primitive.E{Key: id, Value: idDevice}}
	var device Device
	err := collection.FindOne(ctx, query).Decode(&device)
	if err != nil {
		return nil
	}
	return &device
}

func (repository *DeviceRepository) FindAll() []Device {
	ctx, cancel := context.WithTimeout(context.Background(), config.MONGODB_TIMEOUT_SEC*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var devices []Device
	for cur.Next(ctx) {
		var device Device
		err := cur.Decode(&device)
		devices = append(devices, device)
		if err != nil {
			log.Fatal(err)
		}
	}
	return devices
}