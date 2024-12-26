package ticketeer

import (
	"github.com/mishamyrt/ticketeer/pkg/log"
	"github.com/mishamyrt/ticketeer/pkg/log/color"
)

// LogOptions represent application log options
type LogOptions struct {
	Verbose bool
	NoColor bool
}

// App represent ticketeer application
type App struct {
	log log.Logger
}

// New creates new ticketeer application
func New() *App {
	return &App{
		log: log.New(),
	}
}

// SetupLog configures logger
func (a *App) SetupLog(opts LogOptions) {
	var level log.Level
	if opts.Verbose {
		level = log.LevelDebug
	} else {
		level = log.LevelInfo
	}
	a.log.SetLevel(level)
	color.SetNoColor(opts.NoColor)
}

func (a *App) handleError(err error, message string) error {
	a.log.Errorf("%s: %v", message, err)
	return wrapHandledError(err)
}
