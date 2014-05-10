package cmds

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"veyron/lib/cmdline"

	"veyron2/idl/build"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds)
}

func checkErrors(w io.Writer, env *build.Env) {
	if !env.Errors.IsEmpty() {
		fmt.Fprintf(w, "ERROR\n%v", env.Errors.ToError())
		fmt.Fprintln(w, `   (run with "idlc -v" for verbose logging or "idlc help" for help)`)
		os.Exit(2)
	}
}

// runHelper returns a function that generates a sorted list of transitive
// targets, and calls the supplied run function.
func runHelper(run func(targets []*build.BuildPackage, env *build.Env)) func(cmd *cmdline.Command, args []string) error {
	return func(cmd *cmdline.Command, args []string) error {
		if flagVerbose {
			build.SetVerbose()
		}
		if len(args) == 0 {
			// If the user doesn't specify any targets, the cwd is implied.
			args = append(args, ".")
		}
		env := build.NewEnv(flagMaxErrors)
		targets := build.GetTransitiveTargets(args, env)
		checkErrors(cmd.Stderr(), env)
		if len(targets) == 0 {
			// The user's probably confused if we don't end up with any targets.
			return cmd.Errorf("no target packages specified")
		}
		run(targets, env)
		checkErrors(cmd.Stderr(), env)
		return nil
	}
}

const pkgDesc = `
<packages> are a list of packages to process, specified as arguments for each
command.  The format is similar to the go tool.  In its simplest form each
package is an import path; e.g. "veyron/lib/idl".  A package that is an absolute
path or that contains a "." is interpreted as a file system path and denotes the
package in that directory.  A package that ends with "..." does a wildcard match
against all directories with that prefix.  The special import path "all" expands
to all package directories found in all the GOPATH trees.

For more information use "go help packages" to see the standard go package
documentation.
`

var cmdCompile = &cmdline.Command{
	Run:   runHelper(runCompile),
	Name:  "compile",
	Short: "Compile packages and dependencies, but don't generate code",
	Long: `
Compile compiles packages and their transitive dependencies, but does not
generate code.  This is useful to sanity-check that your IDL files are valid.
`,
	ArgsName: "<packages>",
	ArgsLong: pkgDesc,
}

var cmdGen = &cmdline.Command{
	Run:   runHelper(runGen),
	Name:  "generate",
	Short: "Compile packages and dependencies, and generate code",
	Long: `
Generate compiles packages and their transitive dependencies, and generates code
in the specified languages.
`,
	ArgsName: "<packages>",
	ArgsLong: pkgDesc,
}

var cmdListInfo = &cmdline.Command{
	Run:   runHelper(runListInfo),
	Name:  "listinfo",
	Short: "List package and dependency info in transitive order",
	Long: `
Listinfo returns information about packages and their transitive dependencies,
in transitive order.  This is the same order the generate and compile commands
use for processing.  If "idlc listinfo A" is run and A depends on B, which
depends on C, the returned order will be C, B, A.  If multiple packages are
specified the ordering is over all combined dependencies.

Reminder: cyclic dependencies between packages are not allowed.  Cyclic
dependencies between IDL files within the same package are also not allowed.
This is more strict than regular Go; it makes it easier to generate code for
other languages like C++.
`,
	ArgsName: "<packages>",
	ArgsLong: pkgDesc,
}

const (
	genLangGo genLang = iota
	genLangJava
	numGenLang
)

type genLang int

func (l genLang) String() string {
	switch l {
	case genLangGo:
		return "go"
	case genLangJava:
		return "java"
		// Add other languages...
	}
	panic(fmt.Errorf("Unhandled language %d", l))
}

func genLangFromString(str string) genLang {
	switch str {
	case "go":
		return genLangGo
	case "java":
		return genLangJava
		// Add other languages...
	}
	panic(fmt.Errorf("Unknown language %s", str))
}

type genLangs map[genLang]bool

func (gls genLangs) String() string {
	ret := "["
	for gl, _ := range gls {
		if ret != "[" {
			ret += ", "
		}
		ret += gl.String()
	}
	ret += "]"
	return ret
}

func (gls genLangs) Set(value string) error {
	// We allow this flag to be repeated on the cmdline.
	for _, str := range strings.Split(value, ",") {
		gls[genLangFromString(str)] = true
	}
	return nil
}

type genJavaDirTranslation struct {
	srcSuffix, dstSuffix string
}

func (gj genJavaDirTranslation) String() string {
	return fmt.Sprintf("{src: %q, dst: %q", gj.srcSuffix, gj.dstSuffix)
}

func (gj genJavaDirTranslation) Set(value string) error {
	strs := strings.Split(value, "->")
	if len(strs) != 2 {
		return fmt.Errorf("string format must be 'srcDir->dstDir', have: %q", value)
	}
	gj.srcSuffix = strs[0]
	gj.dstSuffix = strs[1]
	return nil
}

var (
	// Common flags for the tool itself, applicable to all commands.
	flagVerbose   bool
	flagMaxErrors int

	// Options for each command.
	optCompileStatus         bool
	optGenStatus             bool
	optGenGoFmt              bool
	optGenJavaDirTranslation = genJavaDirTranslation{
		srcSuffix: "v/src",
		dstSuffix: "java/src",
	}
	optGenJavaPkgPrefix string
	optGenLangs         = genLangs{genLangGo: true}
)

// Root returns the root command for the IDL tool.
func Root() *cmdline.Command {
	idl := &cmdline.Command{
		Name:  "idlc",
		Short: "Manage veyron IDL source code",
		Long: `
The idlc tool manages veyron IDL source code.  It's similar to the go tool used
for managing Go source code.
`,
		Children: []*cmdline.Command{cmdGen, cmdCompile, cmdListInfo},
	}

	// Common flags for the tool itself, applicable to all commands.
	idl.Flags.BoolVar(&flagVerbose, "v", false, "Turn on verbose logging.")
	idl.Flags.IntVar(&flagMaxErrors, "max_errors", -1, "Stop processing after this many errors, or -1 for unlimited.")

	// Options for compile.
	cmdCompile.Flags.BoolVar(&optCompileStatus, "status", true, "Show package names while we compile")

	// Options for generate.
	var allLangs string
	for lx := 0; lx < int(numGenLang); lx++ {
		if lx > 0 {
			allLangs += ", "
		}
		allLangs += `"` + genLang(lx).String() + `"`
	}
	cmdGen.Flags.Var(&optGenLangs, "lang", "Comma-separated list of languages to generate, currently supporting "+allLangs)
	cmdGen.Flags.BoolVar(&optGenGoFmt, "go_fmt", true, "Format generated Go code")
	cmdGen.Flags.BoolVar(&optGenStatus, "status", true, "Show package names while we compile")
	cmdGen.Flags.StringVar(&optGenJavaPkgPrefix, "java_pkg_prefix", "com",
		"Package prefix that will be added to the IDL package prefixes when generating Java files. ")
	cmdGen.Flags.Var(&optGenJavaDirTranslation, "java_dir_translation",
		"Directory translation for generated Java files.  The rule must be specified in the format: "+
			"srcSuffix->dstSuffix, where srcSuffix is the suffix of the directory just preceding the "+
			"package path, and dstSuffix is the suffix that srcSuffix is replaced with. For example, "+
			"default rule: v/src->java/src will result in IDL files: $DIR/v/src/.../*.idl "+
			"having their Java generated files output in: $DIR/java/src/...")
	return idl
}

func runCompile(targets []*build.BuildPackage, env *build.Env) {
	for _, target := range targets {
		pkg := build.CompilePackage(target, env)
		if pkg != nil && optCompileStatus {
			fmt.Println(pkg.Path)
		}
	}
}

func runGen(targets []*build.BuildPackage, env *build.Env) {
	for _, target := range targets {
		pkg := build.CompilePackage(target, env)
		if pkg == nil {
			continue
		}
		// TODO(toddw): Skip code generation if the semantic contents of the
		// generated file haven't changed.
		changed := false
		for gl, _ := range optGenLangs {
			switch gl {
			case genLangGo:
				for _, file := range pkg.Files {
					opts := build.GenGoOpts{Fmt: optGenGoFmt}
					data := build.GenFileGo(file, opts)
					if writeFile(data, pkg.Dir, file.BaseName+".go", env) {
						changed = true
					}
				}
			case genLangJava:
				build.SetJavaGenPkgPrefix(optGenJavaPkgPrefix)
				files := build.GenJavaFiles(pkg)
				dir, err := javaDir(pkg)
				if err != nil {
					env.Errors.Errorf("Couldn't translate package dir %q for Java generation", pkg.Dir)
					continue
				}
				for _, file := range files {
					if writeFile(file.Data, dir, file.Name, env) {
						changed = true
					}
				}
			default:
				env.Errors.Errorf("Generating code for language %v isn't supported", gl)
			}
		}
		if changed && optGenStatus {
			fmt.Println(pkg.Path)
		}
	}
}

// writeFile writes data into the standard location for file, using the given
// suffix.  Errors are reported via env.  Returns true iff a new file was
// written; returns false if the file already exists with the given data.
func writeFile(data []byte, dirName, baseName string, env *build.Env) bool {
	// Create containing directory, if it doesn't already exist.
	if err := os.MkdirAll(dirName, os.FileMode(0777)); err != nil {
		env.Errors.Errorf("Couldn't create directory %s: %v", dirName, err)
		return false
	}
	dstName := filepath.Join(dirName, baseName)
	// Don't change anything if old and new are the same.
	if oldData, err := ioutil.ReadFile(dstName); err == nil && bytes.Equal(oldData, data) {
		return false
	}
	if err := ioutil.WriteFile(dstName, data, os.FileMode(0666)); err != nil {
		env.Errors.Errorf("Couldn't write file %s: %v", dstName, err)
		return false
	}

	return true
}

func javaDir(pkg *build.Package) (string, error) {
	d := pkg.Dir
	if !strings.HasSuffix(d, pkg.Path) {
		return "", fmt.Errorf("package directory %q must end with package pathname %q", d, pkg.Path)
	}
	d = filepath.Clean(d[:len(d)-len(pkg.Path)])
	src := optGenJavaDirTranslation.srcSuffix
	if !strings.HasSuffix(d, src) {
		return "", fmt.Errorf("package directory %q must end with translation source %q", d, src)
	}
	d = filepath.Clean(d[:len(d)-len(src)])
	return filepath.Join(d, optGenJavaDirTranslation.dstSuffix, optGenJavaPkgPrefix, pkg.Path), nil
}

func runListInfo(targets []*build.BuildPackage, env *build.Env) {
	for tx, target := range targets {
		num := fmt.Sprintf("%d", tx)
		fmt.Println(num, strings.Repeat("=", 80-len(num)))
		fmt.Printf("Name: %v\n", target.Name)
		fmt.Printf("Path: %v\n", target.Path)
		fmt.Printf("Dir:  %v\n", target.Dir)
		if len(target.IDLBaseFileNames) > 0 {
			fmt.Print("Files:\n")
			for _, file := range target.IDLBaseFileNames {
				fmt.Printf("   %v\n", file)
			}
		}
	}
}
