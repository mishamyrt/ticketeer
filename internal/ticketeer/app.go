package ticketeer

import (
	"github.com/mishamyrt/ticketeer/pkg/log"
	"github.com/mishamyrt/ticketeer/pkg/log/color"
)

// App represent ticketeer application
type App struct{}

// Options represent command line options
type Options struct {
	Verbose bool
	NoColor bool
}

// New creates new ticketeer application
func New() *App {
	return &App{}
}

// Setup configures application
func (a *App) Setup(opts *Options) {
	if opts.Verbose {
		log.SetLevel(log.LevelDebug)
	}
	if opts.NoColor {
		color.SetNoColor(true)
	}
}
