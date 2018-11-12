package logger

import (
	"time"
	"encoding/json"
	"log"
	"github.com/utahta/go-cronowriter"
)

type Logger struct {
	OriginalPath string
	ForwardPath  string
	Steps        []Steps
	Status       bool
	StartTime    time.Time
	EndTime      time.Time
}
type Steps struct {
	Step  string
	Error string
}

var LogsInstance *Logger = nil

func GetLogInstance() *Logger {
	if LogsInstance == nil {
		LogsInstance = new(Logger);
	}
	return LogsInstance;
}

func DestroyLogInstance() {
	b, _ := json.Marshal(LogsInstance)
	log.Print(string(b))

	LogsInstance = nil
}

func (log *Logger) InitLog(originalPath string) {
	log.OriginalPath = originalPath
	log.StartTime = time.Now()
	log.Status = false
}

func (log *Logger) AddStep(step string, err string) {
	StepData := Steps{
		Step:  step,
		Error: err,
	}
	log.Steps = append(log.Steps, StepData)
}
