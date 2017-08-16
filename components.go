package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type Home struct {
	vecty.Core

	ctx *js.Object
}

type Nav struct {
	vecty.Core

	links map[string]string
}

func (n *Nav) Render() *vecty.HTML {
	var listItems []vecty.MarkupOrComponentOrHTML
	for display, path := range n.links {
		listItems = append(listItems, elem.ListItem(
			vecty.Tag(
				"a",
				prop.Href(path),
				vecty.Text(display),
			),
		))
	}

	return elem.UnorderedList(listItems...)
}

type App struct {
	vecty.Core

	child vecty.MarkupOrComponentOrHTML
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
