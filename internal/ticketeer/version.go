package ticketeer

const version = "0.1.5"

var commitHash = "unknown"

// Version prints application version
func (a *App) Version(full bool) {
	if !full {
		a.log.Info(version)
		return
	}
	a.log.Infof("%s (%s)", version, commitHash)
}
