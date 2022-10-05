package timerutils

import "time"

func ImmediatelyTick(execution func(), duration time.Duration) {

	execution()

	DelayTick(execution, duration)

}
func DelayTick(execution func(), duration time.Duration) {

	t := time.NewTicker(duration)

	go func() {
		for {
			<-t.C

			execution()

		}

		t.Stop()
	}()

}
