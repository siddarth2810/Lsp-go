package rpc_test

import (
	"lsp/rpc"
	"testing"

	"github.com/magiconair/properties/assert"
)

type EncodingExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	assert.Equal(t, expected, actual)
}

func TestDecode(t *testing.T) {
	incomingMsg := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMsg))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, method, "hi")
	assert.Equal(t, contentLength, 15)
}
