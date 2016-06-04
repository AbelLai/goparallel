package goparallel

type dispatcher struct {
	workerChannel chan chan job
	maxWorker     int
	workers       *[]worker
}

func (d *dispatcher) run() {
	workers := make([]worker, d.maxWorker)

	for i := 0; i < d.maxWorker; i++ {
		worker := newWorker(d.workerChannel)
		workers[i] = worker
		worker.start()
	}

	d.workers = &workers

	go d.dispatch()
}

func (d *dispatcher) stop() {
	for _, worker := range *d.workers {
		worker.stop()
	}
}

func (d *dispatcher) dispatch() {
	for {
		select {
		case _job := <-jobQueue:
			go func(_job job) {
				jobChannel := <-d.workerChannel

				jobChannel <- _job
			}(_job)
		}
	}
}

func newDispatcher(maxWorker int) *dispatcher {
	pool := make(chan chan job, maxWorker)
	return &dispatcher{workerChannel: pool, maxWorker: maxWorker}
}

func newWorker(workerChannel chan chan job) worker {
	return worker{
		workerChannel: workerChannel,
		jobChannel:    make(chan job),
		quit:          make(chan bool),
	}
}
