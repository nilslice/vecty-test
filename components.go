package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/gopherjs/vecty/storeutil"
	"github.com/nilslice/gopherjs/routing/actions"
	"github.com/nilslice/gopherjs/routing/dispatcher"
	"github.com/nilslice/gopherjs/routing/store"
)

var (
	Listeners = storeutil.NewListenerRegistry()
)

type Home struct {
	vecty.Core

	ctx *js.Object
}

type Nav struct {
	vecty.Core

	links map[string]string
}

func (n *Nav) onClick(ev *vecty.Event) {
	print(ev.Target.Get("innerText"))
}

func (n *Nav) Render() *vecty.HTML {
	var listItems []vecty.MarkupOrComponentOrHTML
	for display, path := range n.links {
		listItems = append(listItems, elem.ListItem(
			vecty.Tag(
				"a",
				prop.Href(path),
				vecty.Text(display),
				event.Click(n.onClick),
			),
		))
	}

	return elem.UnorderedList(listItems...)
}

type App struct {
	vecty.Core

	counter int
	child   vecty.MarkupOrComponentOrHTML
}

func (a *App) handleChange(ev *vecty.Event) {
	dispatcher.Dispatch(&actions.SetInputValue{
		Value: ev.Target.Get("value").String(),
	})

	vecty.Rerender(a)
}

func (a *App) Render() *vecty.HTML {
	return elem.Body(
		&Nav{
			links: map[string]string{
				"Home":    "/",
				"Music":   "/music",
				"About":   "/about",
				"Contact": "/contact",
			},
		},
		elem.Input(
			prop.Value(store.GetInputValue()),
			event.Input(a.handleChange),
		),
		vecty.Tag(
			"p",
			vecty.Text(store.GetInputValue()),
		),
		vecty.Tag(
			"p",
			vecty.Text(fmt.Sprintf("%d", a.counter)),
		),
		a.child,
	)
}

func (h *Home) Render() *vecty.HTML {
	name := h.ctx.Get("params").Get("name")
	if name == js.Undefined {
		name = js.MakeWrapper("home")
	}

	return vecty.Tag(
		"h2",
		vecty.Text("Visiting: "+name.String()),
	)
}
