/*
clientstreaming gRPC (03_clientstreaming) server side code
written by namjune jeong
*/
package main

/*
각종 패키지 import
03_clientstreaming/clientstream		: proto파일과 protoc를 이용해 생성한 패키지
"google.golang.org/grpc" 			: 구글이 제공하는 go의 grpc 패키지
*/
import (
	pb "03_clientstreaming/clientstream"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

// pb에 정의되어있는 interface를 만족시키기 위해 구조체 선언
type ClientstreamServer struct {
	pb.ClientstreamServer
}

// interface를 만족시키기 위해 Multiply 메소드 구현
func (s ClientstreamServer) Multiply(server pb.Clientstream_MultiplyServer) error {
	var ans int32 = 1

	for {
		// 데이터 recv
		in, err := server.Recv()

		// 클라이언트의 송신이 끝난 경우 stream을 닫고 지금까지 곱한 값을 전송
		if err == io.EOF {
			log.Println("client send EOF")
			server.SendAndClose(&pb.Res{Answer: ans})
			return nil
		}
		if err != nil {
			log.Printf("Recv error : %v", err)
			continue
		}

		log.Printf("Received: %4d", in.GetValue())
		ans = ans * in.GetValue()
	}
}

func main() {
	// net 모듈로 소켓 열기
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	// gRPC 서버 생성 후 proto로 생성된 clientstreamServer 인터페이스를 만족시킨 메소드를 등록
	grpcServer := grpc.NewServer()
	pb.RegisterClientstreamServer(grpcServer, &ClientstreamServer{})

	log.Println("****************************************")
	log.Println("*        🟢 grpc server started        *")
	log.Println("*        listening on port : 50051     *")
	log.Println("****************************************")
	// 서버 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
