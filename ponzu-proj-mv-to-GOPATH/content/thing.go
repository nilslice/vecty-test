package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Thing struct {
	item.Item

	Name string `json:"name"`
}

// MarshalEditor writes a buffer of html to edit a Thing within the CMS
// and implements editor.Editable
func (t *Thing) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(t,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Thing field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", t, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Thing editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Thing"] = func() interface{} { return new(Thing) }
}
