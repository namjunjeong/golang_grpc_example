/*
serverstreaming gRPC (04_serverstreaming) client side code
written by namjune jeong
*/
package main

/*
각종 패키지 import
04_serverstreaming/serverstream 		: proto파일과 protoc를 이용해 생성한 패키지
"google.golang.org/grpc" 				: 구글이 제공하는 go의 grpc 패키지
*/
import (
	pb "04_serverstreaming/serverstream"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 서버 주소 미리 상수로 선언
const (
	SERVERADDR string = "localhost:50051"
)

func main() {
	var myvalue int32
	fmt.Printf("input value : ")
	fmt.Scanf("%d", &myvalue)
	var data = pb.Req{Value: myvalue}

	// grpc.Dial을 통해 새로운 gRPC채널 생성.
	// no encryption or authentication => https://grpc.io/docs/guides/auth/
	// 필요할 경우 구글 토큰 기반 인증, SSL/TLS 서버 인증, app레벨 보안(ALTS(구글개발)) 등 사용 가능
	conn, err := grpc.Dial(SERVERADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection error : %v", err)
	}
	// 불시에 client가 종료되더라도 conn.Close()함수 실행 보장
	defer conn.Close()

	// proto파일을 통해 생성된 Serverstream패키지의 함수를 이용하여 client(stub) 생성
	client := pb.NewServerstreamClient(conn)

	recvdone := make(chan bool) //서버의 송신이 종료되었음을 확인하기위한 채널

	fmt.Println("request start")
	// Multiply gRPC 함수 호출과 동시에 stream 생성
	stream, err := client.Multiply(context.Background(), &data)
	if err != nil {
		log.Fatalf("open Multiply stream error : %v", err)
	}
	fmt.Println("request finish")

	fmt.Println("response start")
	fmt.Println("-----------------------------------")

	// recv를 위한 go routine. 생성된 stream으로부터 계속해서 정보 recv
	go func() {
		for {
			// Multiply stream으로부터 Recv
			res, err := stream.Recv()

			// 만약 서버의 송신이 완료되었을 경우 done 채널 close
			if err == io.EOF {
				close(recvdone)
				return
			}
			if err != nil {
				log.Fatalf("can not receive : %v", err)
			}
			fmt.Printf("%2d:    %3d * %3d = %6d \n", res.GetMultiplier()-1, myvalue, res.GetMultiplier(), res.GetAnswer())
		}
	}()

	<-recvdone // 수신이 완료될때까지 code block

	fmt.Println("-----------------------------------")
	fmt.Println("response finish")
}
