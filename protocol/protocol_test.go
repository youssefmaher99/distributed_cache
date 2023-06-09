package proto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSetCommand(t *testing.T) {
	cmd := &CommandSet{
		Key:   []byte("hello"),
		Value: []byte("world"),
		TTL:   2,
	}
	r := bytes.NewReader(cmd.Bytes())
	pcmd, err := ParseCommand(r)
	assert.Nil(t, err)
	assert.Equal(t, cmd, pcmd)
}

func TestParseGetCommand(t *testing.T) {
	cmdGet := &CommandGet{Key: []byte("hello")}
	r := bytes.NewReader(cmdGet.Bytes())
	pcmd, err := ParseCommand(r)
	assert.Nil(t, err)

	// fmt.Println(result)
	assert.Equal(t, cmdGet, pcmd)
}

// func BenchmarkParseCommand(b *testing.B) {
// 	cmd := &CommandSet{
// 		Key:   []byte("hello"),
// 		Value: []byte("world"),
// 		TTL:   2,
// 	}
// 	for i := 0; i < b.N; i++ {
// 		r := bytes.NewReader(cmd.Bytes())
// 		ParseCommand(r)
// 	}

// }
