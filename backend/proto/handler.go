package proto

import (
	"context"
	"log"
)

type Server struct {
	// 加算される合計値を保持する
	totalNum int32
}

func (s *Server) AddNum(ctx context.Context, addingNum *AddNumParams) (*TotalNum, error) {
	log.Println("add number.")
	s.totalNum += addingNum.Number
	total := &TotalNum{Total: s.totalNum}
	return total, nil
}

func (s *Server) GetTotalNum(ctx context.Context, _ *GetTotalNumParams) (*TotalNum, error) {
	log.Println("return total number.")
	total := &TotalNum{Total: s.totalNum}
	return total, nil
}


