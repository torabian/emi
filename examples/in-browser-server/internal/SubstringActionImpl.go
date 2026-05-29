package internal

// This file is manual and actual implementation of the action, which needs to be routed
// both in wasm and gin server

import "github.com/torabian/emi/examples/in-browser-server/internal/defs"

func SubstringAction(c defs.SubstringActionRequest) (*defs.SubstringActionResponse, error) {

	return (&defs.SubstringActionResponse{
		Payload: defs.SubstringActionRes{
			Result: c.Body.Input[c.Body.Start:c.Body.End],
		},
	}), nil
}
