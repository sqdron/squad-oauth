package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad-oauth/oauthApi"
	"github.com/sqdron/squad-digitalocean/digitalocean"
	"github.com/sqdron/squad/configurator"
	"fmt"
	"io/ioutil"
)

type AuthOptions struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type Options struct {
	DigitalOcean *AuthOptions
}

func main() {
	o := &Options{}
	fmt.Println("Debug files...")
	files, _ := ioutil.ReadDir("./env")
	for _, f := range files {
		fmt.Println(f.Name())
	}
	configurator.New().ReadFromFile("./env/providers.json", &o)
	fmt.Println(o.DigitalOcean.ClientID)
	api := oauthApi.NewApi(digitalocean.DigitalOcean(o.DigitalOcean.ClientID, o.DigitalOcean.ClientSecret, o.DigitalOcean.RedirectURL))
	client := squad.Client()
	client.Api.Route("oauth_open").Action(api.OpenSession)
	client.Api.Route("oauth_authorize").Action(api.Authorize)
	client.Api.Route("oauth_refresh").Action(api.Refresh)
	client.Activate()
}
