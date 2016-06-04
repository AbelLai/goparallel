package goparallel

var innerDispatcher *dispatcher

func Go(maxQueueSize, maxWorker int, actor func(interface{}) error) {
	initJobQueue(maxQueueSize)

	attachJobActor(actor)

	innerDispatcher = newDispatcher(maxWorker)
	innerDispatcher.run()
}

func Accept(_job interface{}) {
	acceptNewJob(_job)
}

func Stop() {
	innerDispatcher.stop()
}
