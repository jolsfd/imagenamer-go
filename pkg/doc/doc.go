package doc

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
