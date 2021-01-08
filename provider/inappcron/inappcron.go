package inappcron

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/coronatorid/core-onator/provider"
	"github.com/rs/zerolog/log"
)

// InAppCron stands for in application cronjob
type InAppCron struct {
	cronList   []provider.InAppCronAdapter
	ctx        context.Context
	cancelFunc context.CancelFunc
	wg         *sync.WaitGroup
}

// Fabricate new inappcron struct
func Fabricate() *InAppCron {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &InAppCron{
		cronList:   []provider.InAppCronAdapter{},
		ctx:        ctx,
		cancelFunc: cancelFunc,
		wg:         &sync.WaitGroup{},
	}
}

// Inject new in application cronjob adapter
func (i *InAppCron) Inject(inAppCronAdapter ...provider.InAppCronAdapter) {
	i.cronList = append(i.cronList, inAppCronAdapter...)
}

// Run inapp cron
func (i *InAppCron) Run() {
	for _, cron := range i.cronList {
		i.wg.Add(1)
		cronRunner := cron
		go func() {
			for {
				select {
				case <-i.ctx.Done():
					log.Info().Msg(fmt.Sprintf("Stopping %s in application cronjob", cronRunner.Name()))
					cronRunner.Close()
					i.wg.Done()
					return
				case <-time.After(cronRunner.Delay()):
					if err := cronRunner.Run(); err != nil {
						log.Error().Stack().Err(err).Msg(fmt.Sprintf("Failed to run %s in application cronjob", cronRunner.Name()))
					} else {
						log.Info().Msg(fmt.Sprintf("Success run %s in application cronjob", cronRunner.Name()))
					}
				}
			}
		}()
	}
}

// Close inapp cron
func (i *InAppCron) Close() {
	i.cancelFunc()
	i.wg.Wait()
}
