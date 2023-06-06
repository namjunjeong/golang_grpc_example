/*
serverstreaming gRPC (04_serverstreaming) server side code
written by namjune jeong
*/
package main

/*
각종 패키지 import
04_serverstreaming/serverstream			: proto파일과 protoc를 이용해 생성한 패키지
"google.golang.org/grpc" 				: 구글이 제공하는 go의 grpc 패키지
*/
import (
	pb "04_serverstreaming/serverstream"
	"log"
	"net"

	"google.golang.org/grpc"
)

// pb에 정의되어있는 interface를 만족시키기 위해 구조체 선언
type ServerstreamServer struct {
	pb.ServerstreamServer
}

// interface를 만족시키기 위해 Multiply 메소드 구현
func (s ServerstreamServer) Multiply(in *pb.Req, server pb.Serverstream_MultiplyServer) error {
	log.Printf("Received: %4d", in.GetValue())

	// 입력된 값에 대한 구구단 전송
	for i := 1; i < 10; i++ {
		if err := server.Send(&pb.Res{Answer: in.GetValue() * int32(i), Multiplier: int32(i)}); err != nil {
			log.Printf("Send error %v", err)
			return err
		}
		log.Printf("data sended")
	}
	return nil
}

func main() {
	// net 모듈로 소켓 열기
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	// gRPC 서버 생성 후 proto로 생성된 ServerstreamServer 인터페이스를 만족시킨 메소드를 등록
	grpcServer := grpc.NewServer()
	pb.RegisterServerstreamServer(grpcServer, &ServerstreamServer{})

	log.Println("****************************************")
	log.Println("*        🟢 grpc server started        *")
	log.Println("*        listening on port : 50051     *")
	log.Println("****************************************")
	// 서버 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
