package taskrunner

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longLived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		dataSize:   size,
		longLived:  longlived,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	//关闭实例化的runner
	defer func() {
		if r.longLived == false {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()

	for {
		//持续监听
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_DISPATCH {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}

		case e := <-r.Error:
			if e == CLOSE {
				return
			}

		default:

		}

	}
}

func (r *Runner) StartAll() {
	//输入执行信息
	r.Controller <- READY_TO_DISPATCH
	//持续监听
	r.startDispatch()
}
