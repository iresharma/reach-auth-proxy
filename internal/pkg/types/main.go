package types

type MessageInterface struct {
	Name    string
	Headers map[string][]string
	Query   map[string][]string
	Body    map[string]string
	Perm    []string
	Method  string
}

type Error struct {
	Status  int
	Message string
}
