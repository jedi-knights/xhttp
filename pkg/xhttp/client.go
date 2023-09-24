package xhttp

//go:generate mockgen -destination=mocks/mock_client.go -package=mocks github.com/jedi-knights/xhttp/pkg/xhttp ClientInterface

// ClientInterface is an interface for the provided HTTP clients.
type ClientInterface interface {
	Get(requestUrl string, out interface{}) (int, error)
}
