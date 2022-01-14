package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	pb "weavelab.xyz/week2/proto/frequency"
)

func TestCalculate(t *testing.T) {
	s := server{}
	req := pb.InputRequest{
		Text: "mason mason",
	}

	out, err := s.Calculate(context.Background(), &req)
	if err != nil {
		t.Errorf("error in call %v", err)
	}
	fmt.Printf("%v", out)

	var newReq pb.InputRequest
	reqString := "{\"text\":\"oli oli\"}"
	err = json.Unmarshal([]byte(reqString), &newReq)
	if err != nil {
		t.Errorf("error unmarshalling data %v", err)
	}

}
