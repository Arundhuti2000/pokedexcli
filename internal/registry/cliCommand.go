package registry

type CliCommand struct {
	Name            string
	Description     string
	Callback        func() error
	CallbackWithArg func(string) error
}