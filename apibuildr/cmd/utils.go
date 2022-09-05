package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

func makeFiles(files *[]*ApiFile, v interface{}) error {
	for i, f := range *files {
		file, err := os.Create(f.path)
		if err != nil {
			f.created = false
			f.err = err
			log.Println(err)
			return err
		}
		f.created = true
		log.Println("created file", f.path)

		tpl := template.Must(template.New(fmt.Sprintf("tpl-%v", i)).Parse(string(f.template)))
		if err = tpl.Execute(file, v); err != nil {
			f.err = err
			return err
		}

		file.Close()

	}

	return nil
}

func reverseFiles(files *[]*ApiFile) {
	for _, f := range *files {
		if f.created == true {
			if err := os.Remove(f.path); err != nil {
				CheckError(err)
			}
		}
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

func makeUriPath(p string) string {
	// "//{ less:less(?:\\/)?}",
	path := strings.Trim(p, "/")
	if path == "" {
		return "/"
	}
	parts := strings.Split(path, "/")

	if len(parts) == 0 {
		return "/"
	}
	var end string
	var varRegex = regexp.MustCompile(`^{.*}$`)
	if varRegex.MatchString(parts[len(parts)-1]) {
		end = parts[len(parts)-1]
	} else {
		end = fmt.Sprintf(`{%s:%s(?:\\/)?}`, parts[len(parts)-1], parts[len(parts)-1])
	}

	if len(parts) == 1 {
		return end
	}
	apiPath := strings.Join(parts[:len(parts)-1], "/")
	return fmt.Sprintf("/%s/%s", apiPath, end)
}
