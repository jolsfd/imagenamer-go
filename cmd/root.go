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
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/jolsfd/imagenamer-go/pkg/config"
	"github.com/jolsfd/imagenamer-go/pkg/doc"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "imagenamer-go",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init initializes configuration and flags.
func init() {
	// Init config
	cobra.OnInitialize(initConfig)

	// Init flags
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, doc.DebugFlag)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", doc.ConfigFileFlag)
	rootCmd.Version = doc.Version
}

// initConfig reads in config file.
func initConfig() {
	var err error

	// Initalize config.
	viper.SetConfigName(config.DefaultConfigName)
	viper.SetConfigType(config.DefaultConfigType)

	// Set defaults.
	config.DefaultConfig()
	configDir := config.GetConfigDir()
	configFile := config.GetConfigFile()

	// Use config file from the flag.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		err = viper.ReadInConfig()
		checkError(err)
	} else {
		viper.AddConfigPath(configDir)

		err = viper.ReadInConfig()
		if config.CheckLoadError(err).Error() == config.DefaultNoConfigError {
			err = config.WriteConfigFile(configDir, configFile)
			checkError(err)
		}
	}

	// Debug Message:
	if debug {
		color.Cyan("Configfile: %s\n", viper.ConfigFileUsed())
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
