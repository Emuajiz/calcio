package main

import (
	"calcio/config"
	"calcio/internal/group"
	"calcio/internal/logger"
	"calcio/internal/mongo"
	"calcio/internal/rest"
	"context"
	"os"
	"os/signal"
	"time"
)

func main() {

	logger := logger.NewLogger()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	config, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config", "error", err)
	}

	mongoClient := mongo.NewMongoClient()
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		logger.Error("Failed to connect to MongoDB", "error", err)
		return
	}

	mongoDB := mongo.NewMongoDB(config, mongoClient)

	groupMongoRepo := group.NewGroupMongoRepo(mongoDB)
	groupService := group.NewService(groupMongoRepo)
	groupController := group.NewController(logger, groupService)

	echo := rest.NewEcho()
	controller := rest.NewController(echo, logger, groupController)
	controller.RegisterRoutes()

	logger.Info("Starting server on :8080")
	controller.Start(":8080")

	<-ch
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := controller.Stop(ctx); err != nil {
		logger.Error("Error during shutdown", "error", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mongoClient.Disconnect(ctx); err != nil {
		logger.Error("Failed to disconnect MongoDB client", "error", err)
	}

	logger.Info("Server gracefully stopped")
}
