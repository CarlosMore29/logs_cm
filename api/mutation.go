package api

var Mutation string = `
mutation($input: influx_input!){
	insert_logs(input: $input) {
	  affects_rows
	}  
}
`
