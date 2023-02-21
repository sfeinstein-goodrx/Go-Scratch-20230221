package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("github_from_file/resp.json")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	var r reply // only properties in the reply struct will be read; r.Name = "", r.Public_Repos = 0
	dec := json.NewDecoder(file)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	//fmt.Println(r)
	fmt.Printf("v: %v\n", r)   // just value
	fmt.Printf("+v: %+v\n", r) // key:value
	fmt.Printf("#v: %#v\n", r) // key:formatted value [? I need to catch up on this more...it renders as quoted for strings, not for ints]
}

type reply struct {
	Name string
	//Public_Repos int
	NumRepos int `json:"public_repos,omitempty"` // field tag
}
