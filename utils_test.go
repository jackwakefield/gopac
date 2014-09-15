package gopac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPlainHostName(t *testing.T) {
	assert.True(t, isPlainHostName("www"), "'www' should be a valid plain host")
	assert.False(t, isPlainHostName("www.netscape.com"), "'www.netscape.com' should not be a valid plain host")
}

func TestDnsDomainIs(t *testing.T) {
	assert.True(t, dnsDomainIs("www.netscape.com", ".netscape.com"), "'www.netscape.com' should be a valid host for domain '.netscape.com'")
	assert.False(t, dnsDomainIs("www", ".netscape.com"), "'www' should not be a valid host for domain '.netscape.com'")
	assert.False(t, dnsDomainIs("www.mcom.com", ".netscape.com"), "'www.mcom.com' should not be a valid host for domain '.netscape.com'")
}
