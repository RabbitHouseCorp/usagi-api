package usagiapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
) 
 
type Data struct {
	Url string
}

func dance() Data {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/dance")

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


func feed() Data {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/feed")

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


func hug() Data {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/hug")

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


func kiss() Data {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/kiss")

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


func pat() Data {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/pat")

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



func poke() Data {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/poke")

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



func slap() Data {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/slap")

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



func tickle() Data {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/tickle")

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