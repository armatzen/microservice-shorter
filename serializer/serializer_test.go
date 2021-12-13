package serializer

import (
	"fmt"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"microservice-shorter/shortener"
)


var (
	now = time.Now().UTC().Unix()
	input = &shortener.Redirect{
		Code: "abc",
		URL: "https://",
		CreatedAt: 0,
	}
	InputRaw = []byte(fmt.Sprint(
		`{"code":"%s","URL":"%s","createdAt":"%s"}`,
		input.Code,
		input.URL,
		input.createdAt,
	))
)

func TestDecode(t *testing.T){
	redirect:= &Redirect()
	decode, err:= redirect.Decode(InputRaw)
	assert.Nil(t,err)
	assert.Equal(t, input, decode)
}

func TestEncode(t *testing.T){
	redirect:= &Redirect()
	encode, err:= redirect.Encode(input)
	assert.Nil(t,err)
	assert.Equal(t, InputRaw, raw)
}