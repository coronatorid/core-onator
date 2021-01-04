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
	cronList   map[string]provider.InAppCronAdapter
	ctx        context.Context
	cancelFunc context.CancelFunc
	wg         *sync.WaitGroup
}

// Fabricate new inappcron struct
func Fabricate() *InAppCron {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &InAppCron{
		cronList:   map[string]provider.InAppCronAdapter{},
		ctx:        ctx,
		cancelFunc: cancelFunc,
		wg:         &sync.WaitGroup{},
	}
}

// Inject new in application cronjob adapter
func (i *InAppCron) Inject(inAppCronAdapter ...provider.InAppCronAdapter) {
	i.wg.Add(len(inAppCronAdapter))
	for _, cron := range inAppCronAdapter {
		i.cronList[cron.Name()] = cron
	}
}

// Run inapp cron
func (i *InAppCron) Run() {
	for _, cron := range i.cronList {
		cron := cron
		go func() {
			for {
				select {
				case <-i.ctx.Done():
					log.Info().Msg(fmt.Sprintf("Stopping %s in application cronjob", cron.Name()))
					cron.Close()
					i.wg.Done()
					return
				default:
					if err := cron.Run(); err != nil {
						log.Info().Msg(fmt.Sprintf("Failed to run %s in application cronjob", cron.Name()))
					} else {
						log.Info().Msg(fmt.Sprintf("Success run %s in application cronjob", cron.Name()))
					}
					time.Sleep(cron.Delay())

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
