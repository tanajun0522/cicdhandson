package main

import (
	"testing"
)

func TestMakeGreeing(t *testing.T) {
	want := "Hello,Taro"
	got := makerGreeting("Taro")
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}
}
