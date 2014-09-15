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
