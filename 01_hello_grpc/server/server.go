/*
Unary gRPC (01_hello_grpc) server side code
written by namjune jeong
*/
package main

/*
각종 패키지 import
01_hello_grpc/unary 		: proto파일과 protoc를 이용해 생성한 패키지
"google.golang.org/grpc" 	: 구글이 제공하는 go의 grpc 패키지
*/

import (
	pb "01_hello_grpc/unary"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// pb에 정의되어있는 interface를 만족시키기 위해 구조체 선언
type UnaryServer struct {
	pb.UnaryServer
}

// interface를 만족시키기 위해 Multiply 메소드 구현
func (s *UnaryServer) Multiply(ctx context.Context, in *pb.Req) (*pb.Res, error) {
	log.Printf("Received: %4d %8d %8d", in.GetIndex(), in.GetValueA(), in.GetValueB())
	ans := in.GetValueA() * in.GetValueB()
	return &pb.Res{Answer: ans, Origin: in}, nil
}

func main() {
	// net 모듈로 소켓 열기
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	// gRPC 서버 생성 후 proto로 생성된 UnaryServer 인터페이스를 만족시킨 메소드를 등록
	grpcServer := grpc.NewServer()
	pb.RegisterUnaryServer(grpcServer, &UnaryServer{})

	log.Println("****************************************")
	log.Println("*        🟢 grpc server started        *")
	log.Println("*        listening on port : 50051     *")
	log.Println("****************************************")
	log.Println("           idx|    valA|    valB|")
	// 서버 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
