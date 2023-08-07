package runner

type E interface {
	Before(fn func()) // Before is called before the command is executed.
	Do(fn func())     // Do executes the command.
	After(fn func())  // After is called after the command is executed.
	Run()             // Run executes the command.
}

type R struct {
	beforeFn func()
	doFn     func()
	afterFn  func()
}

func New() *R { return &R{} }

func (r *R) Before(fn func()) { r.beforeFn = fn }
func (r *R) Do(fn func())     { r.doFn = fn }
func (r *R) After(fn func())  { r.afterFn = fn }
func (r *R) Run() {
	if r.doFn != nil {
		if r.beforeFn != nil {
			r.beforeFn()
		}
		r.doFn()
		if r.afterFn != nil {
			r.afterFn()
		}
	}
}
