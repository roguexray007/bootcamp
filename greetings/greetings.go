package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v . Welcome!", name)
	return message
}

func HelloFromMaster(name string) string {
	message := fmt.Sprintf("%v's soul is mine ", name)
	return message
}
