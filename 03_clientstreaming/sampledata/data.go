/*
01_hello_grpc의 RPC에 사용되는 미리 정의된 데이터
proto로 정의된 타입을 사용하기위해 pb.Req타입으로 데이터 선언
*/

package sampledata

import (
	pb "03_clientstreaming/clientstream"
)

var TestData = []*pb.Req{
	{
		Value: 3,
	},
	{
		Value: 4,
	},
	{
		Value: 7,
	},
	{
		Value: 1,
	},
	{
		Value: 6,
	},
	{
		Value: 9,
	},
	{
		Value: 13,
	},
	{
		Value: 6,
	},
	{
		Value: 5,
	},
	{
		Value: 2,
	},
}
