package service

import (
	"context"
	"errors"
	"tester/internal/app/v1/repository"
	"tester/internal/app/v1/repository/grpc"
	arithmetical "tester/internal/app/v1/repository/grpc/protos"
	"tester/models"
)

type ArithmeticalService struct {
	rep     repository.Repository
	grpcRep repository.GrpcRepository
}

func NewArithmeticalService(rep repository.Repository, cl arithmetical.ArithmeticalClient) (*ArithmeticalService, error) {
	if rep == nil {
		return nil, errors.New("expected repository, got nil")
	}
	return &ArithmeticalService{
		rep:     rep,
		grpcRep: grpc.NewGrpcRep(cl),
	}, nil
}

func (a ArithmeticalService) Mul(ctx context.Context, mul models.MulRequest) (models.MulResponse, error) {
	resp, err := a.grpcRep.Mul(ctx, mul)
	if err != nil {
		return models.MulResponse{}, err
	}

	mulResp := models.MulResponse{
		Mul:    resp.Mul,
		Factor: resp.Factor,
		Result: resp.Result,
	}

	for _, exp := range resp.Exps {
		mulResp.Exps = append(mulResp.Exps, models.MulExp{
			Val1:             exp.Val1,
			Val2:             exp.Val2,
			IsOverflow:       exp.IsOverflow,
			IsCorrectionStep: exp.IsCorrectionStep,
		})
	}
	return mulResp, nil
}

func (a ArithmeticalService) Code(ctx context.Context, num int32) (models.Codes, error) {
	resp, err := a.grpcRep.Code(ctx, num)
	if err != nil {
		return models.Codes{}, err
	}

	return models.Codes{
		Direct:     resp.Direct,
		Reverse:    resp.Reverse,
		Additional: resp.Additional,
	}, nil
}