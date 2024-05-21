package utils

import (
	"net"
	"regexp"
)

const ipv4Pattern = `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`

var ipv4CompileRegex = regexp.MustCompile(ipv4Pattern)

func IsValidIpv4(ipaddress string) bool {
	return ipv4CompileRegex.MatchString(ipaddress)
}

func ConvertIpAddressIpv6toIpvc4(ipaddres *string) {
	ip := net.ParseIP(*ipaddres).To16().To4()
	*ipaddres = net.IPv4(ip[0], ip[1], ip[2], ip[3]).String()
}
