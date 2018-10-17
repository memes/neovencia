package main

import pb "github.com/neudesic/neovencia/contexts/claim/proto/claim"

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
