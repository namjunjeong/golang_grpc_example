/*
Bidirectional streaming gRPC (02_bidirectional-streaming) server side code
written by namjune jeong
*/
package main

/*
각종 패키지 import
02_bidirectional-streaming/bidirectional		: proto파일과 protoc를 이용해 생성한 패키지
"google.golang.org/grpc" 						: 구글이 제공하는 go의 grpc 패키지
*/
import (
	pb "02_bidirectional-streaming/bidirectional"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

// pb에 정의되어있는 interface를 만족시키기 위해 구조체 선언
type BidirectionalServer struct {
	pb.BidirectionalServer
}

// interface를 만족시키기 위해 Multiply 메소드 구현
func (s BidirectionalServer) Multiply(server pb.Bidirectional_MultiplyServer) error {
	var ans int32

	for {

		// 데이터 recv
		in, err := server.Recv()

		// 클라이언트의 송신이 끝난 경우의 예외 확인
		if err == io.EOF {
			log.Println("client send EOF")
			return nil
		}
		if err != nil {
			log.Printf("Recv error : %v", err)
			continue
		}

		log.Printf("Received: %4d %8d %8d", in.GetIndex(), in.GetValueA(), in.GetValueB())
		ans = in.GetValueA() * in.GetValueB()

		// 데이터 송신
		if err := server.Send(&pb.Res{Answer: ans, Origin: in}); err != nil {
			log.Printf("Send error %v", err)
		}
	}
}

func main() {
	// net 모듈로 소켓 열기
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	// gRPC 서버 생성 후 proto로 생성된 BidirectionalServer 인터페이스를 만족시킨 메소드를 등록
	grpcServer := grpc.NewServer()
	pb.RegisterBidirectionalServer(grpcServer, &BidirectionalServer{})

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
