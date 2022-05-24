package orchestra

type Scheduler interface {
	SelectNode()
	Count()
	Task()
}
