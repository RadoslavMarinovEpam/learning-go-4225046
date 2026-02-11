package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["CA"] = "California"
	states["NY"] = "New York"
	states["TX"] = "Texas"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	// delete(states, "NY")
	// fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
