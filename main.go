//go:generate abigen --sol=./web3/contracts/demo/SADemo.sol --pkg=demo --out=./web3/contracts/demo/SADemo.go
package main

import (
	"bufio"
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/refitor/rslog"
	"github.com/refitor/selfAuth/auth"
	"github.com/refitor/selfAuth/web3"
	"github.com/refitor/selfAuth/web3/contracts/demo"
	"github.com/skip2/go-qrcode"
)

const (
	c_host = ""

	c_env_file       = ".env"
	c_email_host     = ""
	c_email_user     = ""
	c_email_pwd      = ""
	c_email_push     = ""
	c_address_wallet = ""
)

//go:embed web
var memfs embed.FS

type rsFS struct {
	fs embed.FS
}

var (
	v_selfauth_account = ""
	v_selfauth_secret  = ""

	port    = flag.String("port", "9999", "--port=9999")
	bDeploy = flag.Bool("deploy", false, "--deploy=true")

	v_event   *web3.SelfAuthEvent
	v_client  *ethclient.Client
	v_session *demo.SADemoSession
	v_myenv   = make(map[string]string, 0)
)

func main() {
	flag.Parse()
	runForInit()
	runForAuth()
	go runForWeb()
	go runForMonitor()

	for {
		fmt.Printf(
			"Pick an option:\n" + "" +
				"1. Trigger authRequest.\n" +
				"2. Check VerifyCount.\n" +
				"3. Exit.\n",
		)

		// Reads a single UTF-8 character (rune)
		// from STDIN and switches to case.
		switch readStringStdin() {
		case "1":
			runForAuthRequest("trigger")
			break
		case "2":
			runForAuthRequest("verified")
			break
		case "3":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
			break
		}
	}
}

func runForInit() {
	loadEnv(c_env_file)
	auth.InitEmail(c_email_host, c_email_user, c_email_pwd)

	client, err := web3.Init(v_myenv["GATEWAY"], func(s string) string {
		return v_myenv[s]
	})
	fatalCheck(err)
	v_client = client

	session, err := demo.SADemoInit(context.Background(), v_client, func(s string) string {
		if s == "CONTRACTADDR" {
			if *bDeploy {
				return ""
			}
			return v_myenv["CONTRACTADDR"]
		} else if *bDeploy && s == "operate" {
			return "CONTRACTNEW"
		}
		return v_myenv[s]
	})
	fatalCheck(err)
	v_session = session

	if *bDeploy {
		updateEnvFile("CONTRACTADDR", demo.SADemoAddress(), c_env_file)
		fmt.Printf("contract address: %s\n", demo.SADemoAddress())
		os.Exit(0)
	}
}

func runForWeb() {
	http.Handle("/", http.FileServer(&rsFS{memfs}))
	http.HandleFunc("/api/auth/verify", webVerify)
	log.Printf("selfAuth ListenAndServe at %s......\n", *port)
	log.Println(http.ListenAndServe(":"+*port, nil))
}

func (p rsFS) Open(name string) (http.File, error) {
	if name == "/" {
		return http.FS(p.fs).Open("web")
	}
	if _, err := p.fs.Open("web/" + strings.TrimPrefix(name, "/")); err == nil {
		return http.FS(p.fs).Open("web/" + strings.TrimPrefix(name, "/"))
	}
	return nil, errors.New("permission denied")
}

func webVerify(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		httpResponse(w, r, []byte(err.Error()), "text/plain")
		return
	}
	code := r.PostForm.Get("code")
	authID := r.PostForm.Get("authID")
	rslog.Infof("receive selfAuth verify successed, authID: %s, code: %s", authID, code)

	if authID == c_address_wallet {
		if ok, err := auth.NewGoogleAuth().VerifyCode(v_selfauth_secret, code); err != nil {
			httpResponse(w, r, []byte(err.Error()), "text/plain")
			return
		} else if !ok {
			httpResponse(w, r, []byte("selfAuth verify failed"), "text/plain")
			return
		}
	} else {
		httpResponse(w, r, []byte("permission denied"), "text/plain")
		return
	}

	if _, err := runForAuthResponse(); err != nil {
		httpResponse(w, r, []byte(err.Error()), "text/plain")
		return
	} else {
		httpResponse(w, r, []byte("verified successful"), "text/plain")
	}
}

func runForAuth() {
	frunes := []rune(c_address_wallet)
	googleAuth := auth.NewGoogleAuth()
	v_selfauth_secret = googleAuth.GetSecret()
	v_selfauth_account = string(frunes[:4]) + string(frunes[len(frunes)-2:])
	fmt.Println("\nPlease scan the QR code to add an account through google authenticator on your mobile phone:")
	fmt.Println("")
	renderQRCode(googleAuth.GetQrcode(v_selfauth_account, v_selfauth_secret))
}

func runForAuthRequest(operate string) {
	if operate == "trigger" {
		tx, err := v_session.Trigger(common.HexToAddress(c_address_wallet))
		fatalCheck(err)
		rslog.Infof("send trigger to contract successed: %s, TX: %s", c_address_wallet, tx.Hash().String())
	} else if operate == "verified" {
		verifyCount, err := v_session.Verified(common.HexToAddress(c_address_wallet))
		fatalCheck(err)
		rslog.Infof("send verifyCount to contract successed: %s, verified: %v", c_address_wallet, verifyCount.Int64())
	}
}

func runForAuthResponse() (result string, err error) {
	if respTx, respErr := web3.AuthResponse(v_client, v_myenv["CONTRACTADDR"], c_address_wallet, v_event.Params); respErr != nil {
		rslog.Errorf("send response to contract failed, contract: %s, authID: %s, detail: %s", v_myenv["CONTRACTADDR"], c_address_wallet, respErr.Error())
		err = respErr
	} else {
		rslog.Infof("send response to contract successed, contract: %s, authID: %s, respTx: %s", v_myenv["CONTRACTADDR"], c_address_wallet, respTx)
		result = fmt.Sprintf("https://kovan.etherscan.io/tx/%s", respTx)
	}
	return
}

func runForMonitor() {
	rslog.Infof("start monitor for demo %s", v_myenv["CONTRACTADDR"])
	web3.WatchContract(v_client, v_myenv["CONTRACTADDR"], func(l types.Log) error {
		v_event = web3.GetEvent(l)
		contractAddr := l.Address.String()
		authAddr := web3.GetEventAddress(v_event)

		if contractAddr == v_myenv["CONTRACTADDR"] {
			pushLink := fmt.Sprintf("%v?account=%s&random=%s", c_host, v_selfauth_account, authAddr)
			_, err := auth.PushByEmail(c_email_push, "private authorization", "", fmt.Sprintf("[selfAuth] Please click the link and enter the dynamic google verification code, the contract chain operation can only be completed if the verification is successful: %s", pushLink), nil)
			if err != nil {
				rslog.Errorf(err.Error())
				return err
			}
			rslog.Infof("push notify successed, url: %s", pushLink)
		}
		return nil
	})
}

//// help functions
func fatalCheck(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func readStringStdin() string {
	reader := bufio.NewReader(os.Stdin)
	inputVal, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("invalid option: %v\n", err)
		return ""
	}

	output := strings.TrimSuffix(inputVal, "\n") // Important!
	return output
}

func renderQRCode(s string) error {
	q, err := qrcode.New(s, qrcode.Medium)
	if err != nil {
		return err
	}
	fmt.Println(q.ToSmallString(false))
	return nil
}

func httpResponse(w http.ResponseWriter, r *http.Request, data []byte, strType string) {
	rslog.Infof("%s %s ===> Content-Type: %s, Content-Length: %d, Content: %s", r.Method, r.RequestURI, strType, len(data), string(data))
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Type", strType)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//// Contract interaction functions
// loadEnv loads environment variables from location envLoc
// Call this at the top of every function that uses environment variables.
func loadEnv(envFile string) {
	var err error
	if v_myenv, err = godotenv.Read(envFile); err != nil {
		log.Printf("could not load env from %s: %v", envFile, err)
	}
}

// updateEnvFile saves the contract address to our .env file
func updateEnvFile(k, val, envFile string) {
	v_myenv[k] = val
	err := godotenv.Write(v_myenv, envFile)
	if err != nil {
		log.Printf("failed to update %s: %v\n", envFile, err)
	}
}
