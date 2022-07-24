package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

func (a *Api) realizeApiDirectories() error {
	apiMod := fmt.Sprintf("%s/pkg/api", a.ProjectDirectory)
	if _, err := os.Stat(apiMod); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(apiMod, 0754); err != nil {
			return err
		}
	}

	internalMod := fmt.Sprintf("%s/internal", a.ProjectDirectory)
	if _, err := os.Stat(internalMod); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(internalMod, 0754); err != nil {
			return err
		}
	}

	return nil
}

func (a *Api) makeFiles(files *[]*ApiFile) {
	for i, f := range *files {
		file, err := os.Create(f.path)
		if err != nil {
			f.created = false
			f.err = err
		}

		tpl := template.Must(template.New(fmt.Sprintf("tpl-%v", i)).Parse(string(f.template)))
		if err = tpl.Execute(file, a); err != nil {
			f.created = true
			f.err = err
		}

		file.Close()

	}
}

func CheckError(msg interface{}) {
	if msg != nil {
		fmt.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}

func goGet(mod string) error {
	return exec.Command("go", "get", mod).Run()
}

func getModImportPath() string {
	mod, cd := parseModInfo()
	return path.Join(mod.Path, fileToURL(strings.TrimPrefix(cd.Dir, mod.Dir)))
}

func fileToURL(in string) string {
	i := strings.Split(in, string(filepath.Separator))
	return path.Join(i...)
}

func parseModInfo() (Mod, CurDir) {
	var mod Mod
	var dir CurDir

	m := modInfoJSON("-m")
	cobra.CheckErr(json.Unmarshal(m, &mod))

	// Unsure why, but if no module is present Path is set to this string.
	if mod.Path == "command-line-arguments" {
		cobra.CheckErr("Please run `go mod init <MODNAME>` before `apibuildr init`")
	}

	e := modInfoJSON("-e")
	cobra.CheckErr(json.Unmarshal(e, &dir))

	return mod, dir
}

type Mod struct {
	Path, Dir, GoMod string
}

type CurDir struct {
	Dir string
}

func modInfoJSON(args ...string) []byte {
	cmdArgs := append([]string{"list", "-json"}, args...)
	out, err := exec.Command("go", cmdArgs...).Output()
	cobra.CheckErr(err)
	return out
}
