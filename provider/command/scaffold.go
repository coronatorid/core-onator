package command

//go:generate mockgen -source=./scaffold.go -destination=./mocks/scaffold.go -package mockCommand

// Scaffold provide console command injector for core-onator
type Scaffold interface {
	Use() string
	Example() string
	Short() string
	Run(args []string)
}
