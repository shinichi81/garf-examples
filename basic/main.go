package main

import "github.com/backenderia/garf-example/garf"

func main() {
	r := garf.Registry()
	r.Set("URI", "mongodb://localhost:32768")
	r.Configure()
	r.Server().Run(":3000")
}
