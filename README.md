# xhttp

A general purpose testable HTTP client library.


## Installation

To insall the latest version of xhttp, use the following command:

```bash
go get github.com/jedi-knights/xhttp
```

Using the json client to retrieve a response from a REST API:

```go
import (
	"fmt"
	"github.com/jedi-knights/xhttp/pkg/xhttp/json"
)

func main() {
	var err error
	var statusCode int
	
    pClient := json.NewClient()
	requestUrl := "https://jsonplaceholder.typicode.com/posts/1"
    post := struct {
        UserId int    `json:"userId"`
        Id     int    `json:"id"`
        Title  string `json:"title"`
        Body   string `json:"body"`
    }{}

    if statusCode, err = pClient.Get(requestUrl, &post); err != nil {
		fmt.Printf("Error: %v\n", err)
    }
	
	fmt.Printf("Status Code: %d\n", statusCode)
	fmt.Printf("UserId: %d\n", post.UserId)
	fmt.Printf("Id: %d\n", post.Id)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)
}
```
