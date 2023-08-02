package graphql

import (
	"context"

	"github.com/machinebox/graphql"
)

type InputVariable struct {
	Key   string
	Value interface{}
}

type MutationResponse struct {
	AffectedRows int `json:"affects_rows"`
}

func MutationRequest(APIQL_URL, APIQL_SECRET, query string, variables []InputVariable) (map[string]MutationResponse, error) {
	// Type Request
	var respData map[string]MutationResponse
	var errorReq error

	// Client
	client := graphql.NewClient(APIQL_URL)

	// make a request
	req := graphql.NewRequest(query)

	// Variables
	// Variables
	for _, input := range variables {
		req.Var(input.Key, input.Value)
	}

	// set header fields
	req.Header.Set("x-hasura-admin-secret", APIQL_SECRET)

	// define a Context for the request
	ctx := context.Background()

	errorReq = client.Run(ctx, req, &respData)

	return respData, errorReq
}
