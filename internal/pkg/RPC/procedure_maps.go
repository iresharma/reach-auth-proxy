package RPC

import "fmt"

type MessageInterface struct {
	Name    string
	Headers map[string][]string
	Query   map[string][]string
	Body    map[string][]string
	Perm    []string
}

type Error struct {
	Status  int
	Message string
}

func ProceduresMapping(input MessageInterface) (map[string]string, *Error) {
	switch input.Name {
	case "hi":
		fmt.Println("hi procedure in some service would be called")
	default:
		fmt.Println("everything")
	}
	var tempRet = map[string]string{}
	return tempRet, &Error{}
}
