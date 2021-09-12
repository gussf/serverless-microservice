package main

import (
	"fmt"

	"github.com/gussf/serverless-microservice/infra/aws/api_gateway"
)

func main() {
	fmt.Println("MAIN")
	api_gateway.Run()
}
