package main

import (
	"fmt"
	fb "github.com/huandu/facebook/v2"
)

type (
	M = map[string]interface{}
)

func main() {

	// Create a global App var to hold app id and secret.
	var globalApp = fb.New("your-app-id", "your-app-secret")
	var session *fb.Session

	session = globalApp.Session("token")

	// This validates the access token by ensuring that the current user ID is properly returned. err is nil if the token is valid.
    err := session.Validate()
	
	if err != nil {
		panic(err)
	}

	res, _ := session.Get("/me/posts", fb.Params{
		"fields":       "created_time,comments,message",
		"access_token": "EAADhQj7y68YBAB9oZCUb9Vmvbem07apBszKdrLQBCVSx6Pw7o1r3dsDJGI2Pf7mC8WHUaenKsAnisdXgBDRiandw1nyib6UgEDCij12krW0lmbm2ESjQMI1G2erZBkco9diazFEWfypT3Sxncat7btLg9Kk5xWlxQIEboCiEv4UBJKmvzJZCg2ObrHRLwbrsqPJZCyCZBxE1auRRT44fkPZBZCeopjKG7i4ksZAoFCeAuekrk0qnMedRwyImkYrWo18ZD",
	})
	fmt.Println("Here is my Facebook first name:\n", res)
}
