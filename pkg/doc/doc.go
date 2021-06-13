// doc holds all constants for documentation.
package doc

// Root Command
const (
	IamgeNamerShort = "ImageNamer renames images in a directory "
	IamgeNamerLong  = `ImageNamer is a CLI to rename images in folders after a format.
You can use the rename command to rename the images. To change settings change the config file.
For more help please visit: https://github.com/jolsfd/imagenamer-go
`
	ConfigFileFlag = "config file (default is $HOME/.config/imagenamer/config.yaml)"
	DebugFlag      = "show debug messages"
)

// Rename Command
const (
	RenameShort = "Rename images in a directory"
	RenameLong  = `Rename takes paths to folders as argument.
In this folders ImageNamer renames all Images after your config.
To exclude folders use the --exclude flag.Default path when no 
argument is given is the workdirectory.`
	ExcludeFlag    = "exclude folders in your given path"
	SafeRenameFlag = "disable safe rename function"
	YesFlag        = "confirmed all questions with yes"
)

// Version
const Version = "0.0.0"
