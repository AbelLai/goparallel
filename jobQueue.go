package goparallel

type job interface{}
var jobQueue chan job

func initJobQueue(maxJobQueueSize int) {
	jobQueue = make(chan job, maxJobQueueSize)
}

func acceptNewJob(_job interface{}) {
	jobQueue <- _job
}
