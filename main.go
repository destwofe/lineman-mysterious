package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	fmt.Println(string(sd))
	message := []byte{}
	for i := len(sd) - 1; i >= 0; i-- {
		message = append(message, sd[i])
	}

	fmt.Println(string(message))
}
