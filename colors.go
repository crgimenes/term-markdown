package markdown

var (
	// we need a bunch of escape code for manual formatting
	boldOn = "\x1b[1m"
	// boldOff       = "\x1b[21m" --> use resetAll + snapshot with bold off instead
	italicOn      = "\x1b[3m"
	italicOff     = "\x1b[23m"
	crossedOutOn  = "\x1b[9m"
	crossedOutOff = "\x1b[29m"
	greenOn       = "\x1b[32m"

	resetAll = "\x1b[0m"
	colorOff = "\x1b[39m"
)
