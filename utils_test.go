package gopac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPlainHostName(t *testing.T) {
	assert.True(t, isPlainHostName("www"), "'www' should be a valid plain host")
	assert.False(t, isPlainHostName("www.netscape.com"), "'www.netscape.com' should not be a valid plain host")
}
