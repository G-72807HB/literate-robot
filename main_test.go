package main

import (
	"testing"
)

func TestTryit(t *testing.T) {
	// Working
  if(tryit("Fedora") != "Fedora!\n") {
		// Not Working
  // if(tryit("Fedora") != "Hello Fedora!\n") {
    t.Errorf("Found Defect!")
  }
}
