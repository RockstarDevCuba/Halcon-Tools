package command

import (
	"strings"

	utils "github.com/RockstarDevCuba/Halcon-Tools/utils"
)

func ipInfo(in string) {
	ipaddress := strings.TrimSpace(in)
	if !utils.IsValidIpv4(ipaddress) {

	}

}

var IpTracker = &Command{}
