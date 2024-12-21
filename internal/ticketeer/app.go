package ticketeer

import (
	"github.com/mishamyrt/ticketeer/internal/log"
)

// App represent ticketeer application
type App struct {
	log *log.Logger
}

// Options represent command line options
type Options struct {
	Verbose bool
}

// New creates new ticketeer application
func New() *App {
	l := log.New()
	return &App{
		log: l,
	}
}

// Setup configures application
func (a *App) Setup(opts *Options) {
	if opts.Verbose {
		a.log.SetLevel(log.LevelDebug)
	}
}
