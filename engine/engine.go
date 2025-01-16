package engine

import (
	"fmt"
	"sync"
	"time"
)

type Middleware func(input string) string
type EventHandler func(event string, data map[string]interface{})

type Engine struct {
	Logger      *Logger
	Router      *Router
	Utils       *Utils
	Validator   *Validator
	Middlewares []Middleware
	TaskQueue   chan func()
	Events      map[string][]EventHandler
	Running     bool
	mu          sync.Mutex
}

func NewEngine() *Engine {
	return &Engine{
		Logger:      NewLogger(),
		Router:      NewRouter(),
		Utils:       NewUtils(),
		Validator:   NewValidator(),
		Middlewares: []Middleware{},
		TaskQueue:   make(chan func(), 100),
		Events:      make(map[string][]EventHandler),
		Running:     false,
	}
}

func (e *Engine) Start() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.Running {
		e.Logger.Warn("Engine is already running")
		return
	}

	e.Logger.Info("Starting the engine...")
	e.Running = true

	go e.processTaskQueue()
	e.Logger.Info("Engine started successfully")
}

func (e *Engine) Stop() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.Running {
		e.Logger.Warn("Engine is not running")
		return
	}

	e.Logger.Info("Stopping the engine...")
	e.Running = false

	// Close task queue safely
	close(e.TaskQueue)
	e.Logger.Info("Engine stopped successfully")
}

func (e *Engine) Restart() {
	e.Logger.Info("Restarting the engine...")
	e.Stop()
	time.Sleep(1 * time.Second)
	e.Start()
	e.Logger.Info("Engine restarted successfully")
}

func (e *Engine) AddMiddleware(middleware Middleware) {
	e.Middlewares = append(e.Middlewares, middleware)
	e.Logger.Info("Middleware added to the engine")
}

func (e *Engine) RegisterRoute(command string, handler func(string)) {
	e.Router.RegisterRoute(command, handler)
}

func (e *Engine) Run(input string) {
	e.Logger.Info("Engine received input for processing")

	// Apply middleware
	for _, middleware := range e.Middlewares {
		input = middleware(input)
	}

	if !e.Validator.ValidateInput(input) {
		e.Logger.Error("Invalid input provided. Terminating processing.")
		return
	}

	handler, exists := e.Router.GetHandler(input)
	if !exists {
		e.Logger.Error("No handler found for input")
		return
	}

	e.Logger.Info("Processing input...")
	handler(input)
	e.Logger.Info("Input processing completed")
}

func (e *Engine) QueueTask(task func()) {
	if !e.Running {
		e.Logger.Error("Cannot queue task: Engine is not running")
		return
	}

	select {
	case e.TaskQueue <- task:
		e.Logger.Info("Task successfully added to the queue")
	default:
		e.Logger.Error("Task queue is full")
	}
}

func (e *Engine) processTaskQueue() {
	e.Logger.Info("Task queue processor started")
	for task := range e.TaskQueue {
		e.Logger.Info("Executing a task from the queue")
		task()
	}
	e.Logger.Info("Task queue processor stopped")
}

func (e *Engine) On(event string, handler EventHandler) {
	e.Events[event] = append(e.Events[event], handler)
	e.Logger.Info(fmt.Sprintf("Event handler registered for event: %s", event))
}

func (e *Engine) Trigger(event string, data map[string]interface{}) {
	e.Logger.Info(fmt.Sprintf("Triggering event: %s", event))
	if handlers, exists := e.Events[event]; exists {
		for _, handler := range handlers {
			go handler(event, data)
		}
	} else {
		e.Logger.Warn(fmt.Sprintf("No handlers registered for event: %s", event))
	}
}

func LoggingMiddleware(logger *Logger) Middleware {
	return func(input string) string {
		logger.Info(fmt.Sprintf("Middleware: Logging input '%s'", input))
		return input
	}
}

func PrintEventHandler(event string, data map[string]interface{}) {
	fmt.Printf("Event '%s' triggered with data: %v\n", event, data)
}