package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	pb "github.com/neudesic/neovencia/contexts/claim/proto/claim"
)

// Repository is the interface for the ClaimRepository
type Repository interface {
	Create(*pb.Claim) (*pb.Claim, error)
	Get(*pb.GetRequest) (*pb.Claim, error)
}

// ClaimRepository is the persistence layer of claims
type ClaimRepository struct {
	claims []*pb.Claim
}

// Create a claim and add it to the repository
func (repo *ClaimRepository) Create(claim *pb.Claim) (*pb.Claim, error) {
	updated := append(repo.claims, claim)
	repo.claims = updated
	return claim, nil
}

// Get returns all claims in the repository
func (repo *ClaimRepository) Get(request *pb.GetRequest) (*pb.Claim, error) {
	return nil, nil
}

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

func main() {
	repo := &ClaimRepository{}

	srv := micro.NewService(
		micro.Name("neovencia.claim"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterClaimsServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
