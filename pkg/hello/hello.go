package hello

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func PrintHello(name string) {
	fmt.Println(Hello(name))
}
