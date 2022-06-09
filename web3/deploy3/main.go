//go:generate abigen --sol=../contracts_demo/sademo.sol --pkg=contracts_demo --out=../contracts_demo/sademo.go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"rshub/rsapi/service/web3"
	"rshub/rsapi/service/web3/contracts_demo" // replace with go module path used in go.mod
)

var (
	myenv   map[string]string
	vconf   = flag.String("conf", ".env", "--conf=.env")
	bReset  = flag.Bool("reset", false, "--reset=true")
	bDeploy = flag.Bool("deploy", false, "--deploy=true")
)

func main() {
	flag.Parse()
	loadEnv(*vconf)

	if *bReset {
		updateEnvFile("CONTRACTADDR", "", *vconf)
		fmt.Println("Cleared contract address. Bye!")
	} else {
		client, err := web3.Init(myenv["GATEWAY"], myenv["GATEWAY"], nil)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer client.Close()

		// Load and init variables
		if _, err := contracts_demo.Init(context.Background(), client, func(s string) string {
			if *bDeploy && s == "operate" {
				return "CONTRACTNEW"
			}
			return myenv[s]
		}); err != nil {
			log.Fatalln(err.Error())
		}
		updateEnvFile("CONTRACTADDR", contracts_demo.ContractAddress(), *vconf)
		fmt.Printf("contract address: %s\n", contracts_demo.ContractAddress())
	}
}

//// Contract interaction functions

// loadEnv loads environment variables from location envLoc
// Call this at the top of every function that uses environment variables.
func loadEnv(envFile string) {
	var err error
	if myenv, err = godotenv.Read(envFile); err != nil {
		log.Printf("could not load env from %s: %v", envFile, err)
	}
}

// updateEnvFile saves the contract address to our .env file
func updateEnvFile(k, val, envFile string) {
	myenv[k] = val
	err := godotenv.Write(myenv, envFile)
	if err != nil {
		log.Printf("failed to update %s: %v\n", envFile, err)
	}
}
