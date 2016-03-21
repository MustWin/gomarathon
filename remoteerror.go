package gomarathon

type RemoteError interface {
  Code() int
  Error() string
}

type marathonError struct {
  errCode int
  errText string
}

func newRemoteError(errCode int, errText string) RemoteError {
  return &marathonError{errCode, errText}
}

func (e *marathonError) Code() int {
  return e.errCode
}

func (e *marathonError) Error() string {
  return e.errText
}
