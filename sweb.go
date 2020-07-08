package sweb

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type sweb struct {
	values      map[reflect.Type]reflect.Value
	methods     map[string]Handler
	middlewares []Handler
	methodMaps  map[string][]reflect.Type
}

func NewSweb() *sweb {
	return &sweb{
		values:      map[reflect.Type]reflect.Value{},
		methods:     map[string]Handler{},
		methodMaps:  map[string][]reflect.Type{},
		middlewares: make([]Handler, 0),
	}
}

func (s *sweb) Bind(path string, method Handler) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Bind path %v error:%v\n", path, err)
		}
	}()

	if _, ok := s.methods[path]; ok {
		panic(fmt.Sprintf("Duplicate Bind for path:%v", path))
	}

	s.methods[path] = method
}

func (s *sweb) Use(method Handler) {
	s.middlewares = append(s.middlewares, method)
}

func (s *sweb) MapTo(t interface{}) {
	s.values[reflect.TypeOf(t)] = reflect.ValueOf(t)
}

func (s *sweb) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if _, ok := s.methods[path]; !ok {
		notFound(res)
		return
	}
	c := &Context{
		req:    req,
		res:    res,
		values: s.values,
		params: map[string]interface{}{},
		hander: append(s.middlewares, s.methods[path]),
		index:  -1,
	}
	c.Next()
}

func (s *sweb) Run(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s)
}
