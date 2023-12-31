/*
Copyright © 2023 Peeanio

*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os/exec"
	"log"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	REDIRECT = "http%3A%2F%2Flocalhost%3A8080%2Foauth%2Fredirect"
)

var code string

type LoginResponse struct {
	AccessToken	string `json:"access_token"`
	ExpiresIn	int `json:"expires_in"`
	TokenType	string `json:"token_type"`
	RefreshToken	string `json:"refresh_token"`
}

func init() {
  rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
  Use:   "login",
  Short: "Login and store a token for EVE API",
  Long:  `Login and store a token for EVE API, along with the expiry and refresh token data. returns a success description or api error code`,
  Run: func(cmd *cobra.Command, args []string) {
    check_token()
  },
}

func randomBytesInHex(count int) (string, error) {
	buf := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", fmt.Errorf("Could not generate %d random bytes: %v", count, err)
	}

	return hex.EncodeToString(buf), nil
}

func authorizationURL() (string, string, string) {
	//generates and returns secure session vars
	codeVerifier, verifierErr := randomBytesInHex(32) // 64 character string here
	if verifierErr != nil {
		fmt.Fprintln(os.Stderr, "Could not create a code verifier:", verifierErr)
	}
	sha2 := sha256.New()
	_, err := io.WriteString(sha2, codeVerifier)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not write string", err)
	}
	codeChallenge := base64.RawURLEncoding.EncodeToString(sha2.Sum(nil))

	state, stateErr := randomBytesInHex(24)
	if stateErr != nil {
		fmt.Fprintln(os.Stderr, "Could not generate random state:", stateErr)
	}
	return codeVerifier, codeChallenge, state
}

func check_token() {
	//checks if token in config or if valid
	if viper.IsSet("access_token") && viper.IsSet("expires_at"){
		now := time.Now().Unix()
		if viper.GetInt("expires_at") > (int(now) + 120) {
			log.Printf("token exists and is valid, exiting")
		} else {
			log.Printf("first if\n")
			log.Printf("token exists but expired, logging in again")
			login()
		}
	} else {
		log.Printf("token/expiry not configured or invalid, logging in")
		login()
	}
}

func login() {
	//preps login secrets
	verifier, challenge, state := authorizationURL()
	clientId := viper.GetString("client_id")
	scope := viper.GetStringSlice("scope")
	scope_string := strings.Join(scope, "+")
	fmt.Println(scope_string)
	URL := fmt.Sprintf("https://login.eveonline.com/v2/oauth/authorize/?response_type=code&redirect_uri=%s&client_id=%s&scope=%s&code_challenge=%s&code_challenge_method=S256&state=%s", REDIRECT, clientId, scope_string, challenge, state)
	switch runtime.GOOS {
	case "linux":
		go func() {
			cmd :=  exec.Command("xdg-open", URL)
			err := cmd.Run()
			if err != nil {log.Fatal(err)}
		}()
	case "windows":
			go func() {
			cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", URL)
			err := cmd.Run()
			if err != nil {log.Fatal(err)}
		}()
	case "darwin":
		go func() {
			cmd := exec.Command("open", URL)
			err := cmd.Run()
			if err != nil {log.Fatal(err)}
		}()
	}

	//starts an http server to redirect to during OAUTH2
	srv := http.Server{
        Addr:    "localhost:8080",
	}
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		code = r.URL.Query().Get("code")
		state = r.URL.Query().Get("state")
		_, err := io.WriteString(w, "Authorized, you can close this page\n")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not write string", err)
		}
		//shuts down http server, continues process
		go srv.Shutdown(context.Background())
	})
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        panic(err)
	}

	//preps the next stage of login, gets http content and headers set
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", clientId)
	data.Set("code", code)
	data.Set("code_verifier", verifier)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, "https://login.eveonline.com/v2/oauth/token/",  strings.NewReader(data.Encode()))
	if err != nil{
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "login.eveonline.com")
	//calls out to get access token
	resp, _ := client.Do(req)
	now := time.Now().Unix()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response_body LoginResponse
	err = json.Unmarshal(resBody, &response_body)
	if err != nil {
		fmt.Println(err)
	}
	token := response_body.AccessToken
	expires_at := int64(response_body.ExpiresIn) + now
	viper.Set("access_token", token)
	viper.Set("refresh_token", response_body.RefreshToken)
	viper.Set("expires_at", expires_at)
	err = viper.WriteConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Write config failed:", err)
	}
	log.Printf("Successfully logged in and saved access token\n")

}

func refresh_token() {
	clientId := viper.GetString("client_id")
	refresh_token := viper.GetString("refresh_token")
	grant_type := "refresh_token"
	data := url.Values{}
	data.Set("grant_type", grant_type)
	data.Set("client_id", clientId)
	data.Set("refresh_token", refresh_token)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, "https://login.eveonline.com/v2/oauth/token/",  strings.NewReader(data.Encode()))
	if err != nil{
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "login.eveonline.com")
	//calls out to get access token
	resp, _ := client.Do(req)
	now := time.Now().Unix()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response_body LoginResponse
	err = json.Unmarshal(resBody, &response_body)
	if err != nil {
		fmt.Println(err)
	}
	token := response_body.AccessToken
	expires_at := int64(response_body.ExpiresIn) + now
	viper.Set("access_token", token)
	viper.Set("refresh_token", response_body.RefreshToken)
	viper.Set("expires_at", expires_at)
	err = viper.WriteConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Write config failed:", err)
	}
	//fmt.Printf("Successfully logged in and refreshed saved access token\n")
}
