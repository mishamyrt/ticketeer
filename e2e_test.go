//go:build e2e
// +build e2e

package main_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestTicketeerIntegrity(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: filepath.Join("testdata", "scripts"),
		Setup: func(env *testscript.Env) error {
			env.Vars = append(env.Vars, fmt.Sprintf("GOCOVERDIR=%s", os.Getenv("GOCOVERDIR")))
			return nil
		},
	})
}
