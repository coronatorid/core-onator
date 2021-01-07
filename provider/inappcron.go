package provider

import "time"

//go:generate mockgen -source=./inappcron.go -destination=./mocks/inappcron_mock.go -package mockProvider

// InAppCron provider handle all in application cronjob
type InAppCron interface {
	Inject(inAppCronAdapter ...InAppCronAdapter)
	Close()
	Run()
}

// InAppCronAdapter ...
type InAppCronAdapter interface {
	Delay() time.Duration
	Close()
	Run() error
	Name() string
}
