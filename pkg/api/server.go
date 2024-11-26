package api

import (
	"log"
	"net"
	"user-service/pkg/api/handler"
	"user-service/pkg/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewServerHttp(userHandler *handler.UserHandler) *ServerHttp {
	engine := gin.New()

	go NewGrpcServer(userHandler, "8889")

	engine.Use(gin.Logger())

	return &ServerHttp{
		Engine: engine,
	}
}

func NewGrpcServer(userHandler *handler.UserHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	
	if err != nil {
		log.Fatalln("Failed to listen to the GRPC Port", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Printf("gRPC server listening on port: %v\n", grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Could not serve the GRPC Server: ", err)
	}
}

func (s *ServerHttp) Start() {
	s.Engine.Run(":3333")
}