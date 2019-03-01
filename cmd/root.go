package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/ddosakura/gklang"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	// root is the root of the command
	root = &cobra.Command{
		Use:   "gInstaller",
		Short: "Program Installer",
		Long:  `The installer for programs.`,
		Run: func(cmd *cobra.Command, args []string) {
			if install {
				// TODO: install pkg = load plugin, then do the same as prog mod
				return
			}
			// TODO:
			// 1. call `func generate()` in core.go in sblocker
			// 2. save the prog or pkg code in cacheDir
			// 3. add meta data
			// 4. compile the code
			// 5. copy the result to current path
		},
	}

	pkg     bool
	dist    string
	meta    map[string]string
	install bool
)

func init() {
	cobra.OnInitialize(initCLI)

	root.PersistentFlags().BoolVarP(&pkg, "pkg", "p", false, "use pkg mode")
	root.PersistentFlags().StringVarP(&dist, "dist", "d", "./dist", "the dist includes the program need to be installed")
	root.PersistentFlags().StringToStringVarP(&meta, "meta", "m", make(map[string]string), "custom meta data of package")
	root.PersistentFlags().BoolVarP(&install, "install", "i", false, "install the pkg created in the pkg mode")
}

// Execute the cli
func Execute() {
	gklang.Init("gInstaller")
	fmt.Println(logo)

	root.Version = Version
	root.Execute()
}

func initCLI() {
	home, err := homedir.Dir()
	if err != nil {
		gklang.Er(err)
	}
	HomeDir = path.Join(home, ".ddosakura", ".gInstaller")
	CacheDir = path.Join(HomeDir, "cache")
	err = os.MkdirAll(CacheDir, 0755)
	if err != nil {
		gklang.Er(err)
	}
}
