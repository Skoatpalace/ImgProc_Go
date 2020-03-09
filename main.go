package main

import (
	"fmt"
	"time"
	"training.go/imgproc/task"
	"training.go/imgproc/filter"
)

func main()  {
	var f filter.Filter = filter.Grayscale{}
	//t := task.NewWaitGrpTask("./imgs", "output", f)
	t := task.NewChanTask("./imgs", "output", f, 16)

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)
}