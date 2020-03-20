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

//UserProp - Struct to hold user properties information. Also, specifying the corresponding JSON field names
type UserProp struct {
	UID      string   `json:"uid"`
	FullName string   `json:"full_name"`
	Groups   []string `json:"groups"`
}

func main() {

	//If more than expected arguments provided, throw an error
	if len(os.Args) > 3 {
		log.Fatalf("Usage: ./passwd_file_parser paaswd_file_path groups_file_path")
	}

	var paaswd_file_path, groups_file_path string

	//If file paths are not provided with arguments, use default LINUX locations
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

	//Initialize map to hold data in the expected format
	var users = make(map[string]UserProp)

	passwd_reader := bufio.NewReader(passwd_file)
	//Read line by line
	for {
		line, _, err := passwd_reader.ReadLine()

		if err == io.EOF {
			break
		}

		passwd_splits := strings.Split(string(line), ":")

		//Strat populating UID and FullName fields in the map, initialize empty slice for Groups
		users[passwd_splits[0]] = UserProp{UID: passwd_splits[2], FullName: passwd_splits[4], Groups: make([]string, 0)}

	}

	groups_reader := bufio.NewReader(groups_file)
	//Read line by line
	for {
		line, _, err := groups_reader.ReadLine()

		if err == io.EOF {
			break
		}

		groups_splits := strings.Split(string(line), ":")

		//If last column in the groups file is not empty
		if len(groups_splits[3]) > 0 {

			usernames := strings.Split(groups_splits[3], ",")

			//For each username in the group, populate the Groups field in the map
			for _, user := range usernames {

				thisUser := users[user]
				thisUser.Groups = append(thisUser.Groups, groups_splits[0])
				users[user] = thisUser

			}

		}

	}

	//Convert map to JSON. Pretty printing with tabs
	jsonString, err := json.MarshalIndent(users, "", "\t")
	check(err)
	fmt.Println(string(jsonString))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
