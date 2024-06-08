package cli

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func (d *Date[T]) String() string {
	return d.Time.Format(inputDateFormat)
}

func (d verboseFlag) BeforeApply() error {
	log.SetLevel(log.DebugLevel)
	return nil
}

func setDefaultLogs() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}
