package main

import (
	"log"
	"net"
	"os"

	"videostreaming/internal/service"
	pb "videostreaming/proto"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVideoStreamingServer(grpcServer, &service.Server{})

	log.Printf("gRPC server started on :%s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

/*todo: log: | Жағдай                       | Лог |
| ---------------------------- | --- |
| Видео жүктелмеді             | ✅   |
| Видео mp4 емес               | ✅   |
| ffmpeg қате берді            | ✅   |
| Бөлімдер MinIO-ға жүктелмеді | ✅   |
| Уақытша файлдар өшірілмеді   | ✅   |
*/
