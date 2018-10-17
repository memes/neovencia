package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/neudesic/neovencia/contexts/claim/proto/claim"
	"golang.org/x/net/context"
)

const (
	defaultFilename = "claim.json"
)

func parseFile(file string) (*pb.Claim, error) {
	var claim *pb.Claim
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &claim)
	return claim, err
}

func main() {

	cmd.Init()
	client := pb.NewClaimsServiceClient("neovencia.claim", microclient.DefaultClient)

	file := defaultFilename
	claim, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	newClaim, err := client.FileClaim(context.TODO(), claim)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("created: %v", newClaim.Claim)

}
