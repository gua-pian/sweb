package sweb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	Req    *http.Request
	Resp   http.ResponseWriter
	param  httprouter.Params
	values map[reflect.Type]reflect.Value

	handlers []Handler
	index    int
}

func (c *Context) Body() (b []byte) {
	b, _ = ioutil.ReadAll(c.Req.Body)
	return
}

func (c *Context) Args() url.Values {
	return c.Req.URL.Query()
}

func (c *Context) Param(key string) string {
	return c.param.ByName(key)
}

func (c *Context) Res(m Response) {
	c.writeBody(http.StatusOK, m)
}

func (c *Context) ResErr(code int, m Response) {
	c.writeBody(code, m)
}

func (c *Context) writeBody(code int, m Response) {
	c.Resp.Header().Set(TYPE, JSON)
	c.Resp.WriteHeader(code)
	b, _ := json.Marshal(m)
	c.Resp.Write(b)
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}
