package provider

//go:generate mockgen -source=./command.go -destination=./mocks/command_mock.go -package mockProvider

// Command CLI command bearer of coronator
type Command interface {
	InjectCommand(scaffolc ...CommandScaffold)
}

// CommandScaffold use for standard of creating CLI command
type CommandScaffold interface {
	Use() string
	Example() string
	Short() string
	Run(args []string)
}
