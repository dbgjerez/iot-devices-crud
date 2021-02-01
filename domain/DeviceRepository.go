package domain

import (
	"context"
	"iot-devices-crud/config"
	"log"
	"os"
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var server = os.Getenv("MONGODB_HOST")
var dbName = os.Getenv("MONGODB_DEVICE_DB")
var collectionName = os.Getenv("MONGODB_DEVICE_COLLECTION")
var client = config.NewConnection(server)
var collection = client.Database(dbName).Collection(collectionName)

const (
	idKeyName string = "_id"
)

type DeviceRepository struct {
}

func (repository *DeviceRepository) FindById(idDevice string) *Device {
	ctx, cancel := context.WithTimeout(context.Background(), config.MONGODB_TIMEOUT_SEC*time.Second)
	defer cancel()
	query := bson.D{primitive.E{Key: idKeyName, Value: idDevice}}
	var device Device
	err := collection.FindOne(ctx, query).Decode(&device)
	if err != nil {
		return nil
	}
	return &device
}

func (repository *DeviceRepository) FindAll(page int64, size int64) []Device {
	filter := bson.D{}
	paginatedDate, err := mongopagination.New(collection).Limit(size).Page(page).Filter(filter).Find()
	if err != nil {
		log.Fatal(err)
	}
	var devices []Device
	var device *Device
	for _, raw := range paginatedDate.Data {
		if marshallErr := bson.Unmarshal(raw, &device); marshallErr == nil {
			devices = append(devices, *device)
		}

	}
	return devices
}

func (repository *DeviceRepository) CreateDevice(device Device) *Device {
	ctx, cancel := context.WithTimeout(context.Background(), config.MONGODB_TIMEOUT_SEC*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, device)
	log.Println(result.InsertedID, " created")
	if err != nil {
		return nil
	}
	return &device
}

func (repository *DeviceRepository) DeleteDevice(idDevice string) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MONGODB_TIMEOUT_SEC*time.Second)
	defer cancel()
	collection.DeleteOne(ctx, bson.D{primitive.E{Key: idKeyName, Value: idDevice}})
}

func (repository *DeviceRepository) UpdateDevice(idDevice string, device Device) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MONGODB_TIMEOUT_SEC*time.Second)
	defer cancel()
	collection.ReplaceOne(ctx, bson.D{primitive.E{Key: idKeyName, Value: idDevice}}, device)
}
