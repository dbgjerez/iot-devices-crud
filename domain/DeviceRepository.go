package domain

import (
	"context"
	"os"
	"time"

	"github.com/dbgjerez/iot-devices-crud/config"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var server = os.Getenv("MONGODB_HOST")
var dbName = os.Getenv("MONGODB_DEVICE_DB")
var collectionName = os.Getenv("MONGODB_DEVICE_COLLECTION_NAME")
var client = config.NewConnection(server)
var collection = client.Database(dbName).Collection(collectionName)

const (
	idKeyName string = "_id"
)

type DeviceRepository struct {
}

func (repository *DeviceRepository) FindById(idDevice string) (*Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoDefaultTimeout*time.Second)
	defer cancel()
	query := bson.D{primitive.E{Key: idKeyName, Value: idDevice}}
	var device Device
	cur := collection.FindOne(ctx, query)
	if cur.Err() != nil {
		if cur.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, cur.Err()
	}
	err := cur.Decode(&device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (repository *DeviceRepository) FindAll(page int64, size int64) ([]Device, error) {
	filter := bson.D{}
	paginatedDate, err := mongopagination.New(collection).Limit(size).Page(page).Filter(filter).Find()
	if err != nil {
		return nil, err
	}
	var devices []Device
	var device *Device
	for _, raw := range paginatedDate.Data {
		if marshallErr := bson.Unmarshal(raw, &device); marshallErr == nil {
			devices = append(devices, *device)
		}
	}
	return devices, nil
}

func (repository *DeviceRepository) CreateDevice(device Device) (*Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoDefaultTimeout*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (repository *DeviceRepository) DeleteDevice(idDevice string) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoDefaultTimeout*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.D{primitive.E{Key: idKeyName, Value: idDevice}})
	if err != nil {
		return err
	}
	return nil
}

func (repository *DeviceRepository) UpdateDevice(idDevice string, device Device) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoDefaultTimeout*time.Second)
	defer cancel()
	_, err := collection.ReplaceOne(ctx, bson.D{primitive.E{Key: idKeyName, Value: idDevice}}, device)
	if err != nil {
		return err
	}
	return nil
}
