package scheduler

// Emitter emits the state of the CFF scheduler.
type Emitter interface {
	Emit(State)
}

// State describes the status of jobs managed by the CFF scheduler.
type State struct {
	// Pending is total number of jobs, including jobs being executed and
	// waiting to be executed.
	Pending int
	// Ready is number of jobs to be executed but awaiting a free worker.
	// If this number is consistently high, increase the concurrency for this
	// flow.
	Ready int
	// Waiting is number of jobs waiting for other jobs to be finished before
	// being scheduled. If this number is consistently high, your flow has a
	// task that bottlenecks its performance, consider analyzing and
	// restructuring dependencies.
	Waiting int

	// TODO(rhang): Add fields to track number of tasks currently executed
	// and an estimate of idle workers.
}
