package hades

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalWords(t *testing.T) {
	assert.Equal(t, 4, TotalWords("你好世界"))
	assert.Equal(t, 2, TotalWords("hello, playground"))
	//assert.Equal(t, 2, TotalWords("hello,playground"))
	assert.Equal(t, 5, TotalWords("Hello 你好世界"))
	//assert.Equal(t, 5, TotalWords("Hello你好世界"))
}
