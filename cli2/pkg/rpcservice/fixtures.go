package rpcservice

import (
	"context"

	"github.com/blcksec/stripe-cli/pkg/fixtures"
	"github.com/blcksec/stripe-cli/rpc"
)

// Fixture returns the default fixture of given event in string format
func (srv *RPCService) Fixture(ctx context.Context, req *rpc.FixtureRequest) (*rpc.FixtureResponse, error) {
	fixtureFilename := fixtures.Events[req.Event]
	f, err := fixtures.NewFixtureFromFile(nil, "", "", "", fixtureFilename, []string{}, []string{}, []string{}, []string{})
	if err != nil {
		return &rpc.FixtureResponse{Fixture: ""}, err
	}
	return &rpc.FixtureResponse{Fixture: f.GetFixtureFileContent()}, nil
}
