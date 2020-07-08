package sweb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

type Context struct {
	req    *http.Request
	res    http.ResponseWriter
	values map[reflect.Type]reflect.Value
	params map[string]interface{}
	hander []Handler
	index  int
}

func (c *Context) Body() (b []byte) {
	b, _ = ioutil.ReadAll(c.req.Body)
	return
}

func (c *Context) Args() url.Values {
	return c.req.URL.Query()
}

func (c *Context) Res(code int, m Response) {
	c.res.Header().Set("Content-Type", "application/json")
	c.res.WriteHeader(code)
	b, _ := json.Marshal(m)
	c.res.Write(b)
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.hander) {
		c.hander[c.index](c)
		c.index++
	}
}

func (c *Context) SetParams(key string, value interface{}) {
	c.params[key] = value
}

func (c *Context) GetParams(key string) interface{} {
	return c.params[key]
}
