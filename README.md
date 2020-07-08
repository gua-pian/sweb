# sweb
small step for web scaffold  

## Example code
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gua-pian/sweb"
)

type FooRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func foo(c *sweb.Context) {
	var req FooRequest
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		c.Res(http.StatusBadRequest, sweb.Response{"status": -1, "info": "parameter error"})
		return
	}
	c.Res(http.StatusOK, sweb.Response{"status": 0, "info": "success"})
}

func logger(c *sweb.Context) {
	startTime := time.Now().Unix()
	c.Next()
	fmt.Printf("handler cost %d seconds\n", time.Now().Unix()-startTime)
}

func main() {
	s := sweb.NewSweb()
	s.Use(logger)
	s.Bind("/foo/bar", foo)
	log.Fatal(s.Run(3001))
}
```
