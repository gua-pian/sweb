package sweb

type Handler func(ctx *Context)

type Response map[string]interface{}
