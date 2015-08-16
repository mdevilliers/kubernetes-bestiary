Kubernetes Scheduler
--------------------

Code lives at plugin/cmd/kube-scheduler and plugin/pkg/scheduler

High level view
---------------

plugin/pkg/scheduler/scheduler.go

// Run begins watching and scheduling. It starts a goroutine and returns immediately.
func (s *Scheduler) Run() {
	go util.Until(s.scheduleOne, 0, s.config.StopEverything)
}

scheduleOne() {
	// gets next pod request
	// pod := s.config.NextPod()

	// schedules it by finding a host 
	// dest, err := s.config.Algorithm.Schedule(pod, s.config.MinionLister) <-- where the magic happens

	// creates a binding obj
	<!-- &api.Binding{
		ObjectMeta: api.ObjectMeta{Namespace: pod.Namespace, Name: pod.Name},
		Target: api.ObjectReference{
			Kind: "Node",
			Name: dest,
		},  -->

	// sends the information back to K8
	// err := s.config.Binder.Bind(b)
}

s.config.Algorithm implements ScheduleAlgorithm interface

```
type ScheduleAlgorithm interface {
	Schedule(*api.Pod, MinionLister) (selectedMachine string, err error)
}
```

implementation is the generic_scheduler.go which wires up various predicate and prorities implementations


