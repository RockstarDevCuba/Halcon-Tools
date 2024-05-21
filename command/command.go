package command

type Command struct {
	State       int8
	Key         int8
	Name        string
	Description string
	Function    func(in string) any
}
