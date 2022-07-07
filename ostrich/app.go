package ostrich

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strings"
)

type RunFunc func(basename string) error

type App struct {
	basename    string
	name        string
	description string
	runFunc     RunFunc
	noConfig    bool
	args        cobra.PositionalArgs
	cmd         *cobra.Command
}

func WithRunFunc(fn RunFunc) Option {
	return func(a *App) {
		a.runFunc = fn
	}
}

type Option func(*App)

func New(name string, basename string, opts ...Option) *App {
	a := &App{
		name:     name,
		basename: basename,
	}

	for _, o := range opts {
		o(a)
	}

	a.buildCommand()

	return a
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:   FormatBaseName(a.basename),
		Short: a.name,
		Long:  a.description,
		// stop printing usage when the command errors
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          a.args,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true

	cmd.Flags().SetNormalizeFunc(WordSepNormalizeFunc)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	a.cmd = &cmd
}

//
func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	// 显示所有的 flags
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("FLAG: --%s=%q \n", flag.Name, flag.Value)
	})

	if !a.noConfig {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		//if err := viper.Unmarshal(a.options); err != nil {
		//	return err
		//}
	}

	// run application
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}

	return nil
}

func (a *App) Run() error {
	return a.cmd.Execute()
}

// WordSepNormalizeFunc changes all flags that contain "_" separators.
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

func FormatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}
