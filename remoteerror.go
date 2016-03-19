package gomarathon

type RemoteError interface {
  Type() string
  Error() string
}

var (
  HttpServerError400 = NewRemoteError("401", "Bad Request")
  HttpServerError401 = NewRemoteError("401", "Unauthorized")
  HttpServerError403 = NewRemoteError("403", "Forbidden")
  HttpServerError404 = NewRemoteError("404", "Not Found")
)

type MarathonError struct {
  errType string
  errText string
}

func NewRemoteError(errType string, errText string) *MarathonError {
  return &MarathonError{errType, errText}
}

func (e *MarathonError) Type() string {
  return e.errType
}

func (e *MarathonError) Error() string {
  return e.errText
}
