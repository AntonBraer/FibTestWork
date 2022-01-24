package main

import (
	"fbsTest/config"
	"fbsTest/internal/api"
	fibgrpc "fbsTest/internal/api/grpc"
	"fbsTest/internal/repository"
	"fbsTest/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "fbsTest/docs"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

// @title        Fibonacci
// @version      1.0
// @description  FBS test work.

// @host      localhost:8080
// @BasePath  /

func main() {
	config, err := config.LoadConfig("config")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"main": config.MainRedis,
		},
	})

	redCache := cache.New(&cache.Options{
		Redis: ring,
	})

	repo := repository.NewFibRepository(redCache, ring)
	fibService := service.NewFibService(*repo)

	http_server := api.NewServer(fibService, config)
	http_server.InitRouter()
	go func() {
		if err := http_server.Run(config.HttpUrl); err != nil {
			log.Println("HTTP server run error: ", err)
		}
	}()
	log.Println("HTTP run on 8080")
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	fibGrpcServer := fibgrpc.NewServer(*fibService)
	go func() {
		if err := fibGrpcServer.Run(grpcServer); err != nil {
			log.Println("GRPC server run error:", err)
		}
	}()
	log.Println("GRPC run on 8081")

	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
	log.Println("shutdown server")
}
