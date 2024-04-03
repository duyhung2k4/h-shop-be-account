package grpc

import (
	"app/grpc/api"
	"app/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunServerGRPC() {
	listenerGRPC, err := net.Listen("tcp", ":20001")

	if err != nil {
		log.Fatalln(listenerGRPC)
	}

	creds, errKey := credentials.NewServerTLSFromFile(
		"keys/server-account/public.pem",
		"keys/server-account/private.pem",
	)

	if errKey != nil {
		log.Fatalln(errKey)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterProfileServiceServer(grpcServer, api.NewProfileGRPC())
	log.Fatalln(grpcServer.Serve(listenerGRPC))
}
