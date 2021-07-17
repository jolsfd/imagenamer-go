/*
Copyright © 2021 jolsfd

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jolsfd/imagenamer-go/pkg/config"
	"github.com/jolsfd/imagenamer-go/pkg/doc"
	"github.com/jolsfd/imagenamer-go/pkg/question"
	"github.com/jolsfd/imagenamer-go/pkg/rename"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: doc.RenameShort,
	Long:  doc.RenameLong,
	Run: func(cmd *cobra.Command, args []string) {
		renameCommand(cmd, args)
	},
}

// renameCommand rename images in directories.
func renameCommand(cmd *cobra.Command, paths []string) {
	var filesInDirs [][]rename.FileInformation

	// Set Values
	excludes, err := cmd.Flags().GetStringSlice("exclude")
	checkError(err)

	safeRename, err := cmd.Flags().GetBool("safe")
	checkError(err)

	confirm, err := cmd.Flags().GetBool("yes")
	checkError(err)

	workdir, err := os.Getwd()
	checkError(err)

	templateString := viper.GetString(config.Template)
	extensions := viper.GetStringSlice(config.Extensions)
	safePrefixes := viper.GetStringSlice(config.SafeStrings)
	separator := viper.GetString(config.Separator)

	fmt.Printf("%s\n", doc.WaitMessage)

	// Debug:
	if debug {
		color.Cyan("Excludes: %v\nSafeRename: %v\nTemplate: %v\nExtensions: %v\nSafePrefixes: %s\n", excludes, safeRename, templateString, extensions, safePrefixes)
	}

	if len(paths) == 0 {
		paths = append(paths, workdir)
	}

	// Init table.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"SourceName", "TargetName", "Status"})

	// Get Files.
	for _, dir := range paths {
		// Get source names from directory.
		sourceNames, err := rename.ListImagesInDir(dir, extensions, excludes, safeRename, safePrefixes)
		checkError(err)

		// Get files informations from source names.
		files, tableData, err := rename.GetFileInformation(sourceNames, templateString, separator, debug)
		checkError(err)

		// Append to sclice.
		filesInDirs = append(filesInDirs, files)

		// Append bulk.
		table.AppendBulk(tableData)

		// Output.
		fmt.Printf("%v images in %s detected\n", len(sourceNames), dir)
	}

	// Render table.
	table.Render()

	if !confirm {
		// Get number of files.
		numberOfFiles := 0
		for i := range filesInDirs {
			numberOfFiles += len(filesInDirs[i])
		}

		// Ask user.
		confirm = question.YesNo(fmt.Sprintf("Rename %v images?", numberOfFiles))
	}

	// Rename.
	if confirm {
		for _, filesInDir := range filesInDirs {
			err = rename.RenameImages(filesInDir)
			checkError(err)
		}
		fmt.Println(color.GreenString("All images succesfully renamed."))
	} else {
		fmt.Println(color.RedString("No images were renamed."))
	}
}

// init initiates flags and commands.
func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.Flags().BoolP("yes", "y", false, doc.YesFlag)
	renameCmd.Flags().BoolP("safe", "s", true, doc.SafeRenameFlag)
	renameCmd.Flags().StringSliceP("exclude", "e", []string{}, doc.ExcludeFlag)
}
