package handling

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func UrlCheck(website string) (bool, error) {
	var port string
	
	if strings.Contains(website, "https") {
		website = strings.Trim(website, "https://")
		port = "443"
	}
	if strings.Contains(website, "http") {
		website = strings.Trim(website, "http://")
		port = "80"
	}
	check, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", website, port), 1*time.Second)
	if err != nil {
		return false, ErrCheck(err)
	}
	defer check.Close()
	return true, nil
}
