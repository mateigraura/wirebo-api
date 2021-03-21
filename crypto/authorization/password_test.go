package authorization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPasswordCorrectInput_ShouldMatch(t *testing.T) {
	psw := "testPassword"

	hashedPsw, err := HashPassword(psw)
	assert.Nil(t, err)

	assert.True(t, CheckEqual(psw, hashedPsw))
}

func TestHashPasswordIncorrectInput_ShouldFail(t *testing.T) {
	psw := "testPassword"

	hashedPsw, err := HashPassword(psw)
	assert.Nil(t, err)

	incorrectPsw := "incorrectPassword"

	assert.False(t, CheckEqual(incorrectPsw, hashedPsw))
}
