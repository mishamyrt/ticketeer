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

var goCoverDir = fmt.Sprintf("GOCOVERDIR=%s", os.Getenv("GOCOVERDIR"))

func TestTicketeerIntegrity(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: filepath.Join("testdata", "testscripts"),
		Setup: func(env *testscript.Env) error {
			env.Vars = append(env.Vars, goCoverDir)
			return nil
		},
	})
}
