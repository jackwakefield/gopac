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

func TestLocalHostOrDomainIs(t *testing.T) {
	assert.True(t, localHostOrDomainIs("www.netscape.com", "www.netscape.com"), "'www.netscape.com' should be valid as it equals the domain 'www.netscape.com'")
	assert.True(t, localHostOrDomainIs("www", "www.netscape.com"), "'www' should be valid as it contains no domain part")
	assert.False(t, localHostOrDomainIs("www.mcom.com", "wwww.netscape.com"), "'www.mcom.com' should not be as it contains a domain part")
	assert.False(t, localHostOrDomainIs("home.netscape.com", "wwww.netscape.com"), "'home.netscape.com' should not be as it contains a domain part")
}

func TestIsResolvable(t *testing.T) {
	assert.True(t, isResolvable("www.netscape.com"), "'www.netscape.com' should be resolvable")
	assert.False(t, isResolvable("foobar.foobar.foobar"), "'foobar.foobar.foobar' should not be resolvable")
}
