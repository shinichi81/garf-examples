package main

import "github.com/weSolution/gapi-example/gapi"

func main() {
	r := gapi.Registry()
	r.Set("URI", "mongodb://localhost:32768")
	r.Configure()
	r.Server().Run(":3000")
}
