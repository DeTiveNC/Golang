package rpc_test

import (
	"github.com/detivenc/lsp-go/rpc"
	"testing"
)

func TestEncodeMessage(t *testing.T) {
	msg := map[string]interface{}{"method": "echo", "params": "hello"}
	expected := "Content-Length: 34\r\n\r\n{\"method\":\"echo\",\"params\":\"hello\"}"
	if expected != rpc.EncodeMessage(msg) {
		t.Fatal("Expected:", expected, "Got:", rpc.EncodeMessage(msg))
	} else {
		t.Log("TestEncodeMessage passed")
	}
}

func TestDecodeMessage(t *testing.T) {
	data := []byte("Content-Length: 34\r\n\r\n{\"Method\":\"echo\",\"params\":\"hello\"}")
	expectedMsg, expectedLength := "echo", 34
	if method, contentLength, _ := rpc.DecodeMessage(data); expectedMsg != method || expectedLength != len(contentLength) {
		t.Fatal("Expected:", expectedMsg, expectedLength, "Got:", method, contentLength)
	} else {
		t.Log("TestDecodeMessage passed")
	}
}
