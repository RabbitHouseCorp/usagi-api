package rest

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
) 




type Data struct {
	Url string
}


type DataAPI struct {
	Dance string
	Feed string
	Hug string
	Kiss string
	Pat string
	Poke string
	Slap string
	Tickle string
}
 

func Api(endpoint string) Data {
	api  := "https://usagiapi.danielagc.repl.co" + endpoint
	res, err := http.Get(api)

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
	}
	
	return data
}