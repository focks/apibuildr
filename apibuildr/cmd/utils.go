package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

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
