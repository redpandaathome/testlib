package testlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Greet(test *testing.T) {
	actual := Greet("yum")
	expected := "hello~~~~~~~ yum"
	assert.Equal(test, expected, actual)
}
