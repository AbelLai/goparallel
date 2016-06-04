package goparallel

var jobActor (func(interface{}) error)

func attachJobActor(actor func(interface{}) error) {
	jobActor = actor
}
