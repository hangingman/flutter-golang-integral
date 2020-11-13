package service

import (
	"context"

	"github.com/hangingman/flutter-golang-integral/pb"
)

// EchoService 大文字始まりじゃないと公開されないというGoの暗黙ルール
type EchoService struct{}

// Echo はうんたらかんたら
func (s *EchoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}
