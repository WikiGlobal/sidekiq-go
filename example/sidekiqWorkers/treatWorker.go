package sidekiqWorkers

import (
	"fmt"
	"time"

	sidekiq "github.com/oldfritter/sidekiq-go"
)

func CreateTreatWorker(w *sidekiq.Worker) sidekiq.WorkerI {
	return &TreatWorker{*w}
}

type TreatWorker struct {
	sidekiq.Worker
}

func (worker *TreatWorker) Work() (err error) {
	start := time.Now().UnixNano()
	if err != nil {
		worker.LogError("payload: ", *worker.Payload, ", time:", (time.Now().UnixNano()-start)/1000000, " ms")
		return
	}
	fmt.Println("payload: ", *worker.Payload)
	// panic("test panic")

	// err = fmt.Errorf("test")

	// 这里完成此worker的功能

	worker.LogInfo("payload: ", *worker.Payload, ", time:", (time.Now().UnixNano()-start)/1000000, " ms")
	return
}
