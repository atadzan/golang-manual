package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	id := RandStringRunes(31)
	fmt.Println(id)
}

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("g28ERTM5fds14ASDFLKJ7ahjk9POI63lpoiuyQWNBVZqwertGUYmnbvzxcHXC")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	id := "PLB" + string(b)
	return id
}
