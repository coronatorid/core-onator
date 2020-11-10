package provider

// Command provide console command for core-onator
type Command interface {
	Execute() error
}
