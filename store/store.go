package store

import (
	"github.com/gopherjs/vecty/storeutil"
	"github.com/nilslice/gopherjs/routing/actions"
	"github.com/nilslice/gopherjs/routing/dispatcher"
)

var (
	Listeners = storeutil.NewListenerRegistry()

	InputValue    string
	GetInputValue = func() string {
		return InputValue
	}
)

func init() {
	dispatcher.Register(onAction)
}

func onAction(action interface{}) {
	switch a := action.(type) {
	case *actions.SetInputValue:
		InputValue = a.Value

	default:
		return
	}

	Listeners.Fire()
}
