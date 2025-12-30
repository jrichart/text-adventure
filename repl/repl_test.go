package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestReplGameLoop(t *testing.T) {
	startCommand := strings.NewReader("go east")
	output := &bytes.Buffer{}
	Start(startCommand, output)

	want := "A long hallway"
	got := output.String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
