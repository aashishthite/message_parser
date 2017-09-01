package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLinks(t *testing.T) {
	assert := assert.New(t)

	res, err := parseLinks("Olympics are starting soon; http://www.nbcolympics.com")

	assert.Nil(err)
	assert.Equal(1, len(res))
	assert.Equal("http://www.nbcolympics.com", res[0].URL)
	assert.Equal("2018 PyeongChang Olympic Games | NBC Olympics", res[0].Title)
}
