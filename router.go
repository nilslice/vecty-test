package main

//go:generate gopherjs build . -o js/app.js -m

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"

	"github.com/ponzu-cms/go-client"
)

type Router struct {
	*js.Object
}

type handler func(ctx *js.Object, next handler)

func NewRouter() *Router {
	return &Router{js.Global.Get("page")}
}

func (r *Router) Start() {
	opts := js.Global.Get("Object").New()
	r.Call("start", opts)
}

func (r *Router) On(path string, fn handler) {
	r.Invoke(path, fn)
}

func main() {
	router := NewRouter()

	router.On("/", indexHandler)
	router.On("/:name", allHandler)

	router.Start()
}

func indexHandler(ctx *js.Object, next handler) {
	vecty.RenderBody(
		&App{
			child: &Home{
				ctx: ctx,
			},
		},
	)

	ponzu := client.New(client.Config{
		Host: "http://0.0.0.0:8080",
	})

	go func() {
		resp, err := ponzu.Content("Thing", 1)
		if err != nil {
			println("Error", err)
		}

		el := js.Global.Get("document").Call("querySelector", "h2")
		el.Set("innerText", resp.Data[0]["name"].(string))
	}()

}

func allHandler(ctx *js.Object, next handler) {
	vecty.RenderBody(
		&App{
			child: &Home{
				ctx: ctx,
			},
		},
	)
}
