package pokemon

type sem struct {
	c chan struct{}
}

func (s *sem) Aquire() {
	s.c <- struct{}{}
}

func (s *sem) Release() {
	<-s.c
}

//NewSem creates a new semaphore to control the level
//of concurrency
func NewSem(limit int) *sem {
	return &sem{
		c: make(chan struct{}, limit),
	}
}
