package tests

import (
	"fmt"
	"testing"

	"github.com/Jorik887/bankAPI/types"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := types.NewAccount("a", "b", "jopik")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}
