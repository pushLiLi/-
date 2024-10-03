package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

// 构造参数
func NewWorker(internal time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(internal * time.Second),
		runner: r,
	}
}

func (w *Worker) StartWorker() {
	//啥时候开启执行
	//该类的执行函数
	for {
		select {
		//读取到周期执行的时间，开始执行
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.StartWorker()
}
