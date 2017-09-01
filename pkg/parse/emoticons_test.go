package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmotes(t *testing.T) {
	assert := assert.New(t)

	res, err := parseEmoticons("Good morning! (megusta) (coffee)")

	assert.Nil(err)
	assert.Equal(2, len(res))
	assert.Equal("megusta", res[0])
}
