//go:generate abigen --sol=../contracts/demo/SADemo.sol --pkg=demo --out=../contracts/demo/SADemo.go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/refitor/selfAuth/web3"
	"github.com/refitor/selfAuth/web3/contracts/demo" // replace with go module path used in go.mod
)

var (
	myenv   map[string]string
	name    = flag.String("name", "demo", "--name=demo")
	vconf   = flag.String("conf", ".env", "--conf=.env")
	bReset  = flag.Bool("reset", false, "--reset=true")
	bDeploy = flag.Bool("deploy", false, "--deploy=true")
)

func main() {
	flag.Parse()
	loadEnv(*vconf)

	switch *name {
	case "demo":
		deployForDemo()
	}
}

func deployForDemo() {
	client, err := web3.Init(myenv["GATEWAY"], myenv["GATEWAY"], nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer client.Close()

	// Load and init variables
	if _, err := demo.SADemoInit(context.Background(), client, func(s string) string {
		if *bDeploy && s == "operate" {
			return "CONTRACTNEW"
		} else if s == "CONTRACTADDR" && *bReset {
			return ""
		}
		return myenv[s]
	}); err != nil {
		log.Fatalln(err.Error())
	}
	updateEnvFile("CONTRACTADDR", demo.SADemoAddress(), *vconf)
	fmt.Printf("contract address: %s\n", demo.SADemoAddress())
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
