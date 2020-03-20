# parsing_passwd_file_golang

This script parses /etc/passwd and /etc/group files and combines the information to output the following:

```json
{
	"bin": {
		"uid": "1",
		"full_name": "",
		"groups": [
			"sys",
			"daemon"
		]
	},
	"daemon": {
		"uid": "2",
		"full_name": "",
		"groups": [
			"adm",
			"bin"
		]
	}
}
```

The key is the username and the value is another map with UID, Full Name and Groups which the user is a member of.

Script is built with Go version 1.13.4 </br>
To run the project: </br>
**go build passwd_file_parser.go** </br>

Pass the passwd file and group file as arguments: </br>
**./passwd_file_parser /etc/passwd /etc/group** </br>

Alternatively, you can choose to not give file paths and the script will default to using /etc/passwd and /etc/group values: </br>
**./passwd_file_parser**
