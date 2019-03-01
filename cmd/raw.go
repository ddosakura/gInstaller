package cmd

const (
	// Version is the verison of the core
	Version = "v0.0.1"

	// TODO: gk-figlet project
	logo = `
 ____  ____       ____        _                    
|  _ \|  _ \  ___/ ___|  __ _| | ___   _ _ __ __ _ 
| | | | | | |/ _ \___ \ / _\ | |/ / | | | '__/ _\ |
| |_| | |_| | (_) |__) | (_| |   <| |_| | | | (_| |
|____/|____/ \___/____/ \__,_|_|\_\\__,_|_|  \__,_|

gInstaller ` + Version + `
`
)

var (
	// HomeDir for program
	// in Linux: ~/.ddosakura/.gInstaller
	HomeDir string
	// CacheDir for git
	// in Linux: ~/.ddosakura/.gInstaller/cache
	CacheDir string
)
