package config

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


// MongoClient is the shared global MongoDB client such that we can access the data in whatever way we want
var MongoClient *mongo.Client

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

// InitMongo initializes the MongoDB client and assigns it to MongoClient

func InitMongo() {
    LoadEnv()

    uri := os.Getenv("MONGODB_URI")
    opts := options.Client().ApplyURI(uri)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var err error
    MongoClient, err = mongo.Connect(ctx, opts)
    if err != nil {
        log.Fatalf("MongoDB connection error: %v", err)
    }

    if err := MongoClient.Ping(ctx, nil); err != nil {
        log.Fatalf("MongoDB ping error: %v", err)
    }

    log.Println("Connected to MongoDB")

}

func GetDatabase(dbName string) *mongo.Database {
    return MongoClient.Database(dbName)
}