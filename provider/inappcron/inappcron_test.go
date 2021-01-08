package inappcron_test

import (
	"errors"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/provider/inappcron"
)

type dummyInAppCron struct {
	Err error
}

func (*dummyInAppCron) Delay() time.Duration {
	return time.Millisecond * 200
}

func (*dummyInAppCron) Close() {

}

func (d *dummyInAppCron) Run() error {
	return d.Err
}

func (*dummyInAppCron) Name() string {
	return "dummy"
}

func TestInAppCron(t *testing.T) {
	t.Run("Can run and close", func(t *testing.T) {
		inAppCron := inappcron.Fabricate()
		inAppCron.Run()
		inAppCron.Close()
	})

	t.Run("Can run with injected cron adapter", func(t *testing.T) {
		inAppCron := inappcron.Fabricate()
		inAppCron.Inject(&dummyInAppCron{}, &dummyInAppCron{}, &dummyInAppCron{}, &dummyInAppCron{}, &dummyInAppCron{}, &dummyInAppCron{}, &dummyInAppCron{})
		inAppCron.Inject(&dummyInAppCron{
			Err: errors.New("unexpected"),
		})
		inAppCron.Run()
		time.Sleep(time.Second)
		inAppCron.Close()
	})
}
