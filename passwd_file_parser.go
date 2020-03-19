package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//UserProp - Struct to hold user properties information
type UserProp struct {
	UID      string   `json:"uid"`
	FullName string   `json:"full_name"`
	Groups   []string `json:"groups"`
}

func main() {

	if len(os.Args) > 3 {
		log.Fatalf("Usage: ./passwd_file_parser paaswd_file_path groups_file_path")
	}

	var paaswd_file_path, groups_file_path string

	if len(os.Args) < 2 {

		paaswd_file_path = "/etc/passwd"
		groups_file_path = "/etc/group"

	} else {

		paaswd_file_path = os.Args[1]
		groups_file_path = os.Args[2]
	}

	fmt.Println(paaswd_file_path)
	fmt.Println(groups_file_path)

	var users = make(map[string]UserProp)
	users["abc"] = UserProp{UID: "AAA", FullName: "aaaa", Groups: []string{"xyz", "pqr", "jkl"}}

	jsonString, _ := json.Marshal(users)
	fmt.Println(string(jsonString))
}
