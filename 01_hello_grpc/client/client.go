/*
Unary gRPC (01_hello_grpc) client side code
written by namjune jeong
*/
package main

/*
각종 패키지 import
01_hello_grpc/sampledata 	: gRPC요청을 위한 예시 데이터
01_hello_grpc/unary 		: proto파일과 protoc를 이용해 생성한 패키지
"google.golang.org/grpc" 	: 구글이 제공하는 go의 grpc 패키지
*/
import (
	"01_hello_grpc/sampledata"
	pb "01_hello_grpc/unary"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 서버 주소 미리 상수로 선언
const (
	SERVERADDR string = "localhost:50051"
)

func main() {
	var origin *pb.Req
	// grpc.Dial을 통해 새로운 gRPC채널 생성.
	// no encryption or authentication => https://grpc.io/docs/guides/auth/
	// 필요할 경우 구글 토큰 기반 인증, SSL/TLS 서버 인증, app레벨 보안(ALTS(구글개발)) 등 사용 가능
	conn, err := grpc.Dial(SERVERADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection error : %v", err)
	}
	// 불시에 client가 종료되더라도 conn.Close()함수 실행 보장
	defer conn.Close()

	// proto파일을 통해 생성된 unary패키지의 함수를 이용하여 client(stub) 생성
	client := pb.NewUnaryClient(conn)

	// request 시작. 미리 만들어놓은 TestData를 이용
	fmt.Println("request start")
	fmt.Println("-----------------------------------")
	for _, data := range sampledata.TestData {
		fmt.Printf("%d request\n", data.GetIndex())
		// 정의된 Multiply 함수를 이용하여 RPC 수행
		// request deadline, cancellation등도 여기서 설정 가능.
		res, err := client.Multiply(context.Background(), data)
		if err != nil {
			log.Fatalf("Multiply request failed : %v", err)
		}
		origin = res.GetOrigin()
		fmt.Printf("%2d:    %3d * %3d = %6d \n", origin.GetIndex(), origin.GetValueA(), origin.GetValueB(), res.GetAnswer())
	}
	fmt.Println("-----------------------------------")
	fmt.Println("request finish")
	conn.Close()
}
