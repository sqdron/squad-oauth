package oauthApi

import (
	"github.com/sqdron/squad-oauth/oauth"
	"github.com/sqdron/squad/util"
	"golang.org/x/oauth2"
	"fmt"
	"errors"
)

type RequestRegisterOAuthProvider struct {
	Name string
}

type oauthApi struct {
	providers map[string]oauth.IProvider
}

type RequestSession struct {
	Provider string
}

type RequestAuthorize struct {
	oauth.Session
	Provider string
	Code     string
}

type RequestRefresh struct {
	Provider string
	oauth.Session
}

type IAuthApi interface {
	GetAccessUrl(r RequestSession) (string, error)
	Authorize(RequestAuthorize) (string, error)
	Refresh(r RequestRefresh) *oauth2.Token
}

func (a *oauthApi) GetAccessUrl(r RequestSession) (string, error) {
	fmt.Println(r)
	fmt.Println(r.Provider)
	p := a.providers[r.Provider]
	if (p == nil) {
		return "", errors.New("Unknown OAuth provider " + r.Provider )
	}
	fmt.Println(p)
	return p.GetAccessUrl(util.GenerateString(10))
}

func (a *oauthApi) Authorize(r RequestAuthorize) (string, error) {
	fmt.Println(r)
	p := a.providers[r.Provider]
	if (p == nil) {
		return "", errors.New("Invalid OAuth provider")
	}
	return p.Authorize(&r.Session, r.Code)
}

func (a *oauthApi) Refresh(r RequestRefresh) *oauth2.Token {
	p := a.providers[r.Provider]
	res, _ := p.RefreshToken(r.Session.RefreshToken)
	return res
}

func NewApi(providers ...oauth.IProvider) IAuthApi {
	pmap := map[string]oauth.IProvider{}
	for _, v := range providers {
		pmap[v.Name()] = v
	}
	return &oauthApi{providers:pmap}
}