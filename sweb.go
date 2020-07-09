package sweb

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type sweb struct {
	router      *httprouter.Router
	values      map[reflect.Type]reflect.Value
	methods     map[string]Handler
	middlewares []Handler
	methodMaps  map[string][]reflect.Type
}

func NewSweb() *sweb {
	return &sweb{
		router:      httprouter.New(),
		values:      map[reflect.Type]reflect.Value{},
		methods:     map[string]Handler{},
		methodMaps:  map[string][]reflect.Type{},
		middlewares: make([]Handler, 0),
	}
}

func (s *sweb) Bind(method, path string, handler Handler) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Bind method:(%v), path (%v) error:(%v)\n", method, path, err)
		}
	}()
	method = strings.ToUpper(method)
	fc := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		c := &Context{
			Req:      r,
			Resp:     w,
			param:    p,
			values:   s.values,
			handlers: append(s.middlewares, handler),
			index:    -1,
		}
		c.Next()
	}
	s.router.Handle(method, path, fc)
}

func (s *sweb) Any(path string, handler Handler) {
	for _, method := range Methods {
		s.Bind(method, path, handler)
	}
}

func (s *sweb) Use(method Handler) {
	s.middlewares = append(s.middlewares, method)
}

func (s *sweb) MapTo(t interface{}) {
	s.values[reflect.TypeOf(t)] = reflect.ValueOf(t)
}

func (s *sweb) Run(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
}
