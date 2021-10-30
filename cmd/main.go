package main

import (
	"github.com/joho/godotenv"
	"github.com/krupyansky/user-manager/internal/handler"
	"github.com/krupyansky/user-manager/internal/repository"
	"github.com/krupyansky/user-manager/internal/service"
	pb "github.com/krupyansky/user-manager/pkg"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Could not listen on port")
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserApiServer(grpcServer, handlers)

	log.Println("Starting server")
	grpcServer.Serve(listener)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
