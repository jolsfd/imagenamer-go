// Package doc holds all constants for documentation.
package doc

// Root Command
const (
	IamgeNamerShort = "ImageNamer renames images in a directory "
	IamgeNamerLong  = `ImageNamer is a CLI to rename images in folders after a format.
You can use the rename command to rename the images. To change settings change the config file.
For more help please visit: https://github.com/jolsfd/imagenamer-go
`
	ConfigFileFlag = "Location of the configuration file (default: $HOME/.imagenamer/config.yaml)"
	DebugFlag      = "Show debug messages"
)

// Rename Command
const (
	RenameShort = "Rename images in a directory"
	RenameLong  = `Rename takes paths to folders as argument.
In this folders ImageNamer renames all Images after your config.
To exclude folders use the --exclude flag.Default path when no 
argument is given is the workdirectory.`
	ExcludeFlag    = "Exclude folders in given paths"
	SafeRenameFlag = "Disable or enable safe rename function (default: true)"
	YesFlag        = "Confirm all question with yes (default: false)"
	WaitMessage    = "Please wait..."
)

// Version of the program.
const Version = "0.0.0"
