package main

import (
	_ "embed"

	"github.com/siderolabs/talos/pkg/machinery/overlay"
	"github.com/siderolabs/talos/pkg/machinery/overlay/adapter"
)

func main() {
	adapter.Execute(&BoardInstaller{})
}

type BoardInstaller struct{}

type boardExtraOptions struct {
	Console    []string `json:"console"`
	ConfigFile string   `json:"configFile"`
}

func (i *BoardInstaller) GetOptions(extra boardExtraOptions) (overlay.Options, error) {
	kernelArgs := []string{
		"console=tty0",
		"sysctl.kernel.kexec_load_disabled=1",
		"talos.dashboard.disabled=1",
	}

	kernelArgs = append(kernelArgs, extra.Console...)

	return overlay.Options{
		Name:       "orangepi-5",
		KernelArgs: kernelArgs,
	}, nil
}

func (i *BoardInstaller) Install(options overlay.InstallOptions[boardExtraOptions]) error {
	return nil
}
