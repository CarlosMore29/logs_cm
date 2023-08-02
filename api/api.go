package api

import (
	"encoding/json"
	"fmt"

	"github.com/CarlosMore29/logs_cm/api/graphql"
)

func SaveLogs(APIQL_URL, APIQL_SECRET string, inputObj Input) error {
	input := []graphql.InputVariable{
		{Key: "input", Value: inputObj},
	}

	response, errGLobal := graphql.MutationRequest(APIQL_URL, APIQL_SECRET, Mutation, input)

	if errGLobal == nil {
		_, errGLobal = json.Marshal(response)
	}

	if errGLobal != nil || response["insert_logs"].AffectedRows == 0 {
		errGLobal = fmt.Errorf("no se pudo guardar el log | %s", errGLobal)
	}
	return errGLobal
}
