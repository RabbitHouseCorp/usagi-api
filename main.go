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


func dance() {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/dance")

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
		return
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
		return
	}

	fmt.Println(data.Url)
}


func feed() {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/feed")

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
		return
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
		return
	}

	fmt.Println(data.Url)
}


func hug() {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/hug")

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
		return
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
		return
	}

	fmt.Println(data.Url)
}


func kiss() {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/kiss")

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
		return
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
		return
	}

	fmt.Println(data.Url)
}


func pat() {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/pat")

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
		return
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
		return
	}

	fmt.Println(data.Url)
}



func poke() {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/poke")

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
		return
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
		return
	}

	fmt.Println(data.Url)
}



func slap() {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/slap")

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
		return
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
		return
	}

	fmt.Println(data.Url)
}



func tickle() {
	api  := "https://usagiapi.danielagc.repl.co/"
	res, err := http.Get(api + "api/tickle")

	if err != nil {
		fmt.Println("Failed of communication with API ~> Usagi API")
		return
	}

	defer res.Body.Close()

	var data Data
	
	error := json.NewDecoder(io.Reader(res.Body)).Decode(&data)

	if error != nil {
		fmt.Println("Failed with decoding of JSON ~> Usagi API")
		return
	}

	fmt.Println(data.Url)
}