package main

import (
	"training.go/imgproc/task"
	"training.go/imgproc/filter"
)

func main()  {
	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./imgs", "output", f)
	t.Process()
}