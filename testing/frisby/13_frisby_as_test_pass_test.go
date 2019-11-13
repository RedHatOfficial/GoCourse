package main

import (
	"github.com/verdverm/frisby"
	"testing"
)

func TestHttpGet(t *testing.T) {
	f := frisby.Create("Simplest test").Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(200)
	err := f.Error()
	if err != nil {
		t.Error(err)
	}
}
