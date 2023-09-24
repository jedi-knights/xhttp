package xhttp

type ClientInterface interface {
	Get(requestUrl string, out interface{}) (int, error)
}
