package oauth

import (
	"golang.org/x/oauth2"
)

type Provider struct {
	realization interface{}

}

type IProvider interface {
	Name() string
	GetAccessUrl(state string) (string, error)
	Authorize(*Session, string) (string, error)
	GetAccount(*Session) (*Account, error)
	RefreshToken(refreshToken string) (*oauth2.Token, error)
	RefreshTokenAvailable() bool
}

func (p *Provider) Name() string {
	return p.realization.(IProvider).Name();
}

func (p *Provider) GetAccessUrl(state string) (string, error) {
	return p.realization.(IProvider).GetAccessUrl(state);
}

func (p *Provider) Authorize(session *Session, code string) (string, error) {
	return p.realization.(IProvider).Authorize(session, code);
}

func (p *Provider) GetAccount(s *Session) (*Account, error) {
	return p.realization.(IProvider).GetAccount(s);
}

func (p *Provider) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	return p.realization.(IProvider).RefreshToken(refreshToken);
}

func (p *Provider) RefreshTokenAvailable() bool {
	return p.realization.(IProvider).RefreshTokenAvailable();
}

func NewProvider(provider interface{}) IProvider {
	return &Provider{realization:provider}
}