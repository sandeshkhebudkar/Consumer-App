package database

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connect is use to add connection with db-client
func Connect() (*mongo.Client, error) {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("config") // path to look for the config file in
	err1 := viper.ReadInConfig()  // Find and read the config file
	if err1 != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err1))
	}
	uri := viper.GetString("mongodb.URI")

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil

}

//Disconnect is use to break connection
func Disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
