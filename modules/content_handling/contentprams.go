package content_handling

import (
	"fmt"
	"regexp"
)

func UrlHandling(body string) {
	b, _ := regexp.Compile(`([\w-]+.php)+([\?])?([\w]+)=([\w-]+)`)
	prams := b.FindAllString(body, -1)

	for _, php := range prams {
		if php != "" {
			fmt.Println("ContentPrams >> " + php)
		}
	}
}
