package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMentions(t *testing.T) {
	assert := assert.New(t)

	res, err := parseMentions("@chris you around?")

	assert.Nil(err)
	assert.Equal(1, len(res))
	assert.Equal("chris", res[0])

}
