package oauth

type IAuth interface {
	GetAccessUrl() (string, error)
	Authorize(s *Session) (*Session, error)
	Refresh(s *Session) (*Session, error)
}