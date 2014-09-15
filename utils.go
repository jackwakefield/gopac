package gopac

import (
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/robertkrimen/otto"
)

const dnsServer = "8.8.8.8:53"

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

// isResolvable returns true if the host is resolvable.
func isResolvable(host string) bool {
	if len(host) == 0 {
		return false
	}

	if _, err := net.ResolveIPAddr("ip4", host); err != nil {
		return false
	}

	return true
}

// isInNet returns true if the IP address of the host matches the specified IP
// address pattern.
// mask is the pattern informing which parts of the IP address to match against.
// 0 means ignore, 255 means match.
func isInNet(host, pattern, mask string) bool {
	if len(host) == 0 {
		return false
	}

	address, err := net.ResolveIPAddr("ip4", host)

	if err != nil {
		return false
	}

	maskIp := net.IPMask(net.ParseIP(mask))
	return address.IP.Mask(maskIp).String() == pattern
}

// dnsResolve returns the IP address of the host.
func dnsResolve(host string) string {
	address, err := net.ResolveIPAddr("ip4", host)

	if err != nil {
		return ""
	}

	return address.String()
}

// myIpAddress returns the IP address of the host machine.
func myIpAddress() otto.Value {
	hostname, err := os.Hostname()

	if err != nil {
		return otto.UndefinedValue()
	}

	address := dnsResolve(hostname)

	if value, err := otto.ToValue(address); err == nil {
		return value
	}

	return otto.UndefinedValue()
}

// dnsDomainLevels returns the number of domain levels in the host.
func dnsDomainLevels(host string) int {
	return strings.Count(host, ".")
}

// shExpMatch returns true if the string matches the specified shell expression.
func shExpMatch(str, shexp string) bool {
	shexp = strings.Replace(shexp, ".", "\\.", -1)
	shexp = strings.Replace(shexp, "?", ".?", -1)
	shexp = strings.Replace(shexp, "*", ".*", -1)
	matched, err := regexp.MatchString(shexp, "^"+str+"$")

	return err == nil && matched
}
