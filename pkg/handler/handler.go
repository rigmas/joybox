package handler

import (
	"os"
	"os/signal"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Handler interface {
	Start() error
	Stop()
}

type handlerError struct {
	err error
	id  string
}

type Registry struct {
	handlers  map[string]Handler
	wg        *sync.WaitGroup
	errorChan chan handlerError
}

func NewRegistry() *Registry {
	return &Registry{
		handlers:  make(map[string]Handler),
		wg:        new(sync.WaitGroup),
		errorChan: make(chan handlerError),
	}
}

func (r *Registry) Register(id string, s Handler) {
	r.handlers[id] = s
}

func (r *Registry) run(k string, s Handler) {
	go func() {
		defer r.wg.Done()
		log.Println("Starting handler", k)
		err := s.Start()
		if err != nil {
			r.errorChan <- handlerError{
				id:  k,
				err: err,
			}
		}
		log.Println(k, "started successfully")
	}()
}

func (r *Registry) wait() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	log.Println("waiting for signal")

	select {
	case <-signalCh:
		log.Println("interrupted")
	case err := <-r.errorChan:
		log.Errorln("fatal error for service:", err.id)
		log.Errorln(err.err)
	}
}

func (r *Registry) StopAll() {
	for k, s := range r.handlers {
		log.Println("Stopping handler", k)
		s.Stop()
		log.Println(k, "stopped succesfully")
	}
	r.wg.Wait()
	log.Println("All handler stopped")
}

func (r *Registry) StartAll() {
	for k, s := range r.handlers {
		r.wg.Add(1)
		r.run(k, s)
	}
	r.wait()
}
