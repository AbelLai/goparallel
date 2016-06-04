package goparallel

import (
	log "fmt"
)

type worker struct {
	workerChannel chan chan job
	jobChannel chan job
	quit chan bool
}

func (w worker) start() {
	go func() {
		for {
			w.workerChannel <- w.jobChannel

			select {
			case job := <- w.jobChannel:
				if err := jobActor(job); err != nil {
					log.Errorf("Error from job actor")
				}
			case <- w.quit:
				return
			}
		}
	}()
}

func (w worker) stop() {
	go func() {
		w.quit <- true
	}()
}







