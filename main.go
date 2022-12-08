package main

import (
	"C"
	"fmt"
	"net/url"
	"uscan/modules/exploits"

	"github.com/fatih/color"
)

func main() {
	var url string

	fmt.Println(banner())

	fmt.Printf(color.GreenString("Url: "))

	fmt.Scanln(&url)

	if IsUrl(url) {
		exploits.StartingProcess(url)
	} else {
		fmt.Printf("Invalid url, provide something like (https://google.com/)")
	}

}

// Validate if string is a url
func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// uscan banner
func banner() string {
	banner := `
		╦ ╦┌─┐┌─┐┌─┐┌┐┌
		║ ║└─┐│  ├─┤│││
		╚═╝└─┘└─┘┴ ┴┘└┘
			` +
		color.RedString("スキャナー")

	return banner
}
