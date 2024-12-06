package random

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	"strings"
	"time"
)

type Log []api.Log

// NewLog generates a random log.
func NewLog(count int, logLevels map[string]float32) Log {
	logs := make([]api.Log, count)

	if logLevels == nil || len(logLevels) == 0 {
		logLevels = map[string]float32{
			"debug": 1,
			"info":  5,
			"warn":  2,
			"error": 1,
		}
	}

	// Convert levels to gofakeit's format.
	var levels []any
	var weights []float32

	for level, weight := range logLevels {
		levels = append(levels, level)
		weights = append(weights, weight)
	}

	for i := range logs {
		level, _ := gofakeit.Weighted(levels, weights)
		levelString := fmt.Sprint(level)
		now := time.Now()
		message := gofakeit.HackerPhrase()

		logs[i] = api.Log{
			Level:     &levelString,
			Message:   &message,
			Timestamp: &now,
		}
	}

	return logs
}

func (l Log) String() string {
	var sb strings.Builder
	for _, log := range l {
		j, _ := json.Marshal(log)
		sb.WriteString(string(j))
		sb.WriteString("\n")
	}

	return sb.String()
}
