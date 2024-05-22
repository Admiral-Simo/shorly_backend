package tools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUrlHash(t *testing.T) {
	for i := 0; i < 100; i++ {
		hash := CreateUrlHash()
		fmt.Println("hash:", hash)
		assert.Equal(t, 8, len(hash), "should be equal to 6")
	}
}
