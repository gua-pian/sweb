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
		c.ResErr(http.StatusBadRequest, sweb.Response{"status": -1, "info": "parameter error"})
		return
	}
	c.Res(sweb.Response{"status": 0, "info": "success"})
}

func bar(c *sweb.Context) {
	name := c.Param("name")
	c.Res(sweb.Response{"name": name})
}

func logger(c *sweb.Context) {
	startTime := time.Now().Unix()
	c.Next()
	fmt.Printf("handler cost %d seconds\n", time.Now().Unix()-startTime)
}

func main() {
	s := sweb.NewSweb()
	s.Use(logger)
	s.Bind("post", "/foo/bar", foo)
	s.Any("/bar/:name", bar)
	log.Fatal(s.Run(3001))
}

```
