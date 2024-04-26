package workers

// If more than 5, we should make the struct of w with funcs and workers, it should be better btw
// but now we have only 3 and we can don't think about it.
// Should we do it? I do it btw
type Worker struct {
	name     string
	function func()
}

func NewWorker(name string, function func()) *Worker {
	Worker := &Worker{
		name:     name,
		function: function,
	}
	return Worker
}

func InitWorkers() []*Worker {
	res := make([]*Worker, 0)
	res = append(res, NewWorker("time", TimeWorker))
	res = append(res, NewWorker("graph", GraphWorker))
	res = append(res, NewWorker("binary tree", BinaryTreeWorker))

	return res
}
func Workers() {
	workers := InitWorkers()
	for _, work := range workers {
		go work.function()
	}
}
