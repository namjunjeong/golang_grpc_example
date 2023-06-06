/*
01_hello_grpc의 RPC에 사용되는 미리 정의된 데이터
proto로 정의된 타입을 사용하기위해 pb.Req타입으로 데이터 선언
*/

package sampledata

import (
	pb "01_hello_grpc/unary"
)

var TestData = []*pb.Req{
	{
		Index:  0,
		ValueA: 123,
		ValueB: 54,
	},
	{
		Index:  1,
		ValueA: 1,
		ValueB: 3,
	},
	{
		Index:  2,
		ValueA: 3,
		ValueB: 4,
	},
	{
		Index:  3,
		ValueA: 2,
		ValueB: 9,
	},
	{
		Index:  4,
		ValueA: 12,
		ValueB: 8,
	},
	{
		Index:  5,
		ValueA: 43,
		ValueB: 4,
	},
	{
		Index:  6,
		ValueA: 232,
		ValueB: 98,
	},
	{
		Index:  7,
		ValueA: 56,
		ValueB: 2,
	},
	{
		Index:  8,
		ValueA: 89,
		ValueB: 245,
	},
	{
		Index:  9,
		ValueA: 8,
		ValueB: 468,
	},
}
