package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	pb "github.com/neudesic/neovencia/contexts/claim/proto/claim"
)

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
