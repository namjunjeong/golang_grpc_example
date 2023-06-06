/*
clientstreaming gRPC (03_clientstreaming) client side code
written by namjune jeong
*/
package main

/*
각종 패키지 import
03_clientstreaming/sampledata 			: gRPC요청을 위한 예시 데이터
03_clientstreaming/clientstream 		: proto파일과 protoc를 이용해 생성한 패키지
"google.golang.org/grpc" 				: 구글이 제공하는 go의 grpc 패키지
*/
import (
	pb "03_clientstreaming/clientstream"
	"context"
	"fmt"
	"log"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 서버 주소 미리 상수로 선언
const (
	SERVERADDR string = "localhost:50051"
)

func main() {
	// grpc.Dial을 통해 새로운 gRPC채널 생성.
	// no encryption or authentication => https://grpc.io/docs/guides/auth/
	// 필요할 경우 구글 토큰 기반 인증, SSL/TLS 서버 인증, app레벨 보안(ALTS(구글개발)) 등 사용 가능
	conn, err := grpc.Dial(SERVERADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection error : %v", err)
	}
	// 불시에 client가 종료되더라도 conn.Close()함수 실행 보장
	defer conn.Close()

	// proto파일을 통해 생성된 clientstreaming 패키지의 함수를 이용하여 client(stub) 생성
	client := pb.NewClientstreamClient(conn)
	// Multiply 함수를 수행하는 stream 생성
	stream, err := client.Multiply(context.Background())
	if err != nil {
		log.Fatalf("open Multiply stream error : %v", err)
	}

	fmt.Println("request start")
	fmt.Println("-----------------------------------")

	var data pb.Req
	for i := 0; i < 10; i++ {
		// 총 10회, 1~10사이에서 랜덤하게 생성된 수 전송
		data = pb.Req{Value: rand.Int31n(10) + 1}
		if err := stream.Send(&data); err != nil {
			log.Fatalf("Send error : %v", err)
		}
		fmt.Printf("%d request sent \n", data.GetValue())
	}
	fmt.Println("-----------------------------------")
	fmt.Printf("request finish\n\n")

	// 송신이 모두 완료되었으면 CloseSend 호출
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("CloseAndRecv error : %v", err)
	}
	fmt.Printf("multiply all value = %d", res.GetAnswer())

}
