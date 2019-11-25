package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/samtech09/grpc-vs-rest/grpc/service"
	"github.com/samtech09/grpc-vs-rest/server/handlers"

	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func main() {
	go startGrpcServer()
	startRestServer()
}

func startGrpcServer() {
	//listen on local port 8888
	port := ":8888"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s", port)
	}

	// create server instance
	s := handlers.Server{}

	// create grpc server
	grpcServer = grpc.NewServer()

	// register service to grpc server
	service.RegisterReportServer(grpcServer, &s)

	// start the server
	fmt.Println("Grpc server listening on port ", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func stopGrpcServer() {
	grpcServer.GracefulStop()
}

func startRestServer() {
	//listen on local port 8888
	port := ":8080"
	//http.HandleFunc("/getDetail", handlers.GetDetailRest)
	http.HandleFunc("/getDetailbyPost", handlers.GetDetailRestByPost)
	http.HandleFunc("/getDetailbyPostJsoniter", handlers.GetDetailRestByPostJsoniter)
	http.HandleFunc("/getDetailbyPostMsgpack", handlers.GetDetailRestByPostMsgpack)
	fmt.Printf("Server Init")

	fmt.Println("REST server listening on port ", port)
	http.ListenAndServe(port, nil)
}
