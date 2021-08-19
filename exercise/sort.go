package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"baixinho", "amor", "escola", "coracao"}
	sort.Strings(ss)
	fmt.Println(ss)

}
