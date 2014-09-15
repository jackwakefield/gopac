package gopac

import "strings"

// isPlainHostName return true if there is no domain name in the host.
func isPlainHostName(host string) bool {
	return strings.Index(host, ".") == -1
}

// dnsDomainIs return true if the host is valid for the domain.
func dnsDomainIs(host, domain string) bool {
	if len(host) < len(domain) {
		return false
	}

	return strings.HasSuffix(host, domain)
}

// localHostOrDomainIs returns true if the host matches the specified hostdom,
// or if there is no domain name part in the host, but the unqualified hostdom
// matches.
func localHostOrDomainIs(host, hostdom string) bool {
	if host == hostdom {
		return true
	}

	return strings.LastIndex(hostdom, host+".") == 0
}
