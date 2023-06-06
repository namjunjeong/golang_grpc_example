/*
Bidirectional gRPC (02_bidirectional-streaming) client side code
written by namjune jeong
*/
package main

/*
각종 패키지 import
02_bidirectional-streaming/sampledata 			: gRPC요청을 위한 예시 데이터
02_bidirectional-streaming/bidirectional 		: proto파일과 protoc를 이용해 생성한 패키지
"google.golang.org/grpc" 						: 구글이 제공하는 go의 grpc 패키지
*/
import (
	pb "02_bidirectional-streaming/bidirectional"
	"02_bidirectional-streaming/sampledata"
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
	// grpc.Dial을 통해 새로운 gRPC채널 생성.
	// no encryption or authentication => https://grpc.io/docs/guides/auth/
	// 필요할 경우 구글 토큰 기반 인증, SSL/TLS 서버 인증, app레벨 보안(ALTS(구글개발)) 등 사용 가능
	conn, err := grpc.Dial(SERVERADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection error : %v", err)
	}
	// 불시에 client가 종료되더라도 conn.Close()함수 실행 보장
	defer conn.Close()

	// proto파일을 통해 생성된 bidirectional패키지의 함수를 이용하여 client(stub) 생성
	client := pb.NewBidirectionalClient(conn)
	// Multiply 함수를 수행하는 stream 생성
	stream, err := client.Multiply(context.Background())
	if err != nil {
		log.Fatalf("open Multiply stream error : %v", err)
	}

	recvdone := make(chan bool) //서버의 송신이 종료되었음을 확인하기위한 채널

	fmt.Println("request start")
	fmt.Println("-----------------------------------")

	// send를 위한 go routine. 계속해서 서버에게 정보 send
	go func() {
		for _, data := range sampledata.TestData {
			if err := stream.Send(data); err != nil {
				log.Fatalf("Send error : %v", err)
			}
			fmt.Printf("%d request sent \n", data.GetIndex())
		}

		// 송신이 모두 완료되었으면 CloseSend 호출
		if err := stream.CloseSend(); err != nil {
			log.Printf("CloseSend error : %v", err)
		}
	}()

	// recv를 위한 go routine. 계속해서 서버로부터 정보 recv
	go func() {
		var origin *pb.Req
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

			origin = res.GetOrigin()
			fmt.Printf("%2d:    %3d * %3d = %6d \n", origin.GetIndex(), origin.GetValueA(), origin.GetValueB(), res.GetAnswer())
		}
	}()

	<-recvdone // 수신이 완료될때까지 code block
	fmt.Println("-----------------------------------")
	fmt.Println("request finish")
}
