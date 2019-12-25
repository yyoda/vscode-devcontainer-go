package main

// Job .
type Job struct {
	Name string
}

func worker(jobChan chan Job, done chan struct{}) {

	a := make(map[string]struct {
		x int
		y int
	})

	a["aa"] = struct {
		x int
		y int
	}{
		1, 2,
	}

	_ = a["aa"].x

	if a != nil {
		println("aaaaaaa")
	}

	for {
		select {
		case job, ok := <-jobChan:
			if ok {
				println("call:" + job.Name)
			}
		case <-done:
			println("done")
			return
		}
	}
}

func run(jobs []Job) {
	jobChan := make(chan Job)
	done := make(chan struct{})
	go worker(jobChan, done)

	for _, j := range jobs {
		jobChan <- j
	}

	done <- struct{}{}
}

func main() {
	run([]Job{
		Job{Name: "a"},
		Job{Name: "b"},
	})

	println("was completed.")
}

var _ I = (*foo)(nil)

type I interface {
	doSomething()
}

type foo struct {
}

func (f *foo) doSomething() {
}
