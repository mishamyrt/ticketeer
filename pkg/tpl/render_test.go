package tpl_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/pkg/tpl"
)

func TestRender(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		tpl       tpl.Template
		variables tpl.Variables
		want      string
		wantErr   bool
	}{
		{
			"hello, {name}",
			tpl.Variables{"name": "world"},
			"hello, world",
			false,
		},
		{
			"hello, {name}",
			tpl.Variables{},
			"hello, undefined",
			true,
		},
		{
			"{foo}\n\n{bar}",
			tpl.Variables{"foo": "foo", "bar": "bar"},
			"foo\n\nbar",
			false,
		},
		{
			"{foo}\n\n{bar}",
			tpl.Variables{"foo": "foo"},
			"foo\n\nundefined",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.tpl.String(), func(t *testing.T) {
			got, err := tt.tpl.Render(tt.variables)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Render() got = %v, want %v", got, tt.want)
			}
		})
	}
}
