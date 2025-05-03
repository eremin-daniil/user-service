package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	"user_service/adapter/controller"
	"user_service/adapter/gateway"
	"user_service/adapter/repository"
	"user_service/domain/usecase"
	jwtService "user_service/infrastructure/jwt"
	"user_service/infrastructure/logger"
	fileLogger "user_service/infrastructure/logger/file"
	mainLogger "user_service/infrastructure/logger/main"
	"user_service/infrastructure/logger/model"
	mongoRepository "user_service/infrastructure/mongo"
	"user_service/infrastructure/rest"
)

const (
	AppID       = "MarketplaceFeedbacksApplication"
	ServicePort = ":8080"

	JWTSecretKeyEnv = "JWT_SECRET_KEY"
	JWTTokenTTL     = 60 * 60 * 24 * 2
)

const (
	NameLogFile          = "app.log"
	LoggerBufferSize     = 100
	FileLoggerBufferSize = 100
	MainLoggerBufferSize = 100
)

const mongoDBName = "user_service"
const mongoURI = "mongodb://localhost:27017"

func main() {
	ctx := context.Background()
	loggerService := logger.NewService(LoggerBufferSize)

	mainLog := mainLogger.NewMainLogger(MainLoggerBufferSize)
	loggerService.RegisterLogger(mainLogger.LoggerID, mainLog)

	file, err := os.OpenFile(NameLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		mainLog.WriteLog(model.ErrorLogData(ctx, "failed run file logger"))
	} else {
		fileLog := fileLogger.NewFileLogger(file, AppID, FileLoggerBufferSize)
		loggerService.RegisterLogger(fileLogger.LoggerID, fileLog)
	}

	loggerService.Start()
	log := logger.NewLogger(loggerService.InputChan())

	mongoDatabase, _ := MongoDatabase(mongoURI, mongoDBName)
	mongoRepo := mongoRepository.NewMongoRepository(mongoDatabase, log)

	userRepo := repository.NewUserRepository(mongoRepo)
	userGateway := gateway.UserEventDrivenGateway{}
	profile := gateway.ProfileEventDrivenGateway{}
	userUseCase := usecase.NewUserUseCase(&userGateway, &profile, userRepo)
	userController := controller.NewUserController(log, userUseCase)
	profileController := controller.NewProfileController(log, userUseCase)

	jwt := jwtService.NewJWTService(JWTSecretKeyEnv, JWTTokenTTL)
	server := rest.NewMuxServer(log, jwt)

	server.RegisterPrivateRoute("GET", "/users", userController.GetCurrentUser)
	server.RegisterPrivateRoute("GET", "/users/profile", profileController.GetProfileCurrentUser)
	server.RegisterPrivateRoute("PUT", "/users/profile", profileController.UpdateProfile)
	server.RegisterPrivateRoute("POST", "/users/profile", profileController.CreateProfile)

	server.RegisterPublicRoute("POST", "/users", userController.RegisterUser)
	server.RegisterPublicRoute("GET", "/users/{login}", userController.GetByLogin)

	server.Start(ctx, ServicePort)
}

func MongoDatabase(url, mongoDBName string) (*mongo.Database, error) {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	pingTimeout := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), pingTimeout)
	defer cancelFunc()
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return mongoClient.Database(mongoDBName), nil
}
