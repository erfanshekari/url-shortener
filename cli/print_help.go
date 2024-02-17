package cli

import "fmt"

func printHelp() {
	fmt.Println("URL Shortener")
	fmt.Println("usage:")
	fmt.Println("url-shortener -c [config.path]")
}
