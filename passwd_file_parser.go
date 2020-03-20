package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

	passwd_file, err := os.Open(paaswd_file_path)
	check(err)
	defer passwd_file.Close()

	groups_file, err := os.Open(groups_file_path)
	check(err)
	defer groups_file.Close()

	var users = make(map[string]UserProp)

	passwd_reader := bufio.NewReader(passwd_file)
	for {
		line, _, err := passwd_reader.ReadLine()

		if err == io.EOF {
			break
		}

		passwd_splits := strings.Split(string(line), ":")

		users[passwd_splits[0]] = UserProp{UID: passwd_splits[2], FullName: passwd_splits[4], Groups: make([]string, 0)}

	}

	groups_reader := bufio.NewReader(groups_file)
	for {
		line, _, err := groups_reader.ReadLine()

		if err == io.EOF {
			break
		}

		groups_splits := strings.Split(string(line), ":")

		if len(groups_splits[3]) > 0 {

			usernames := strings.Split(groups_splits[3], ",")
			for _, user := range usernames {

				if thisUser, ok := users[user]; ok {
					thisUser.Groups = append(thisUser.Groups, groups_splits[0])
					users[user] = thisUser
				}

			}

		}

	}

	jsonString, err := json.MarshalIndent(users, "", "\t")
	check(err)
	fmt.Println(string(jsonString))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
