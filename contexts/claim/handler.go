package main

import (
	"context"

	pb "github.com/neudesic/neovencia/contexts/claim/proto/claim"
)

type service struct {
	repo Repository
}

func (s *service) FileClaim(ctx context.Context, req *pb.Claim, res *pb.Response) error {
	claim, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Claim = claim
	return nil
}

func (s *service) ClaimStatus(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	claim, err := s.repo.Get(req)
	if err != nil {
		return err
	}
	res.Claim = claim
	return nil
}

func (s *service) AdjustClaim(ctx context.Context, req *pb.Adjustment, res *pb.Response) error {
	//TODO Add adjustment to adjustments table
	return nil
}

func (s *service) SumbitEstimate(ctx context.Context, req *pb.Estimate, res *pb.Response) error {
	//TODO Add estimate to estimate table
	return nil
}

func (s *service) IssuePayment(ctx context.Context, req *pb.Claim, res *pb.Response) error {
	//TODO Call payment service
	return nil
}
