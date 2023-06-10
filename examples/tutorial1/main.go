package main

import (
	"fmt"
	"os"

	"gitea.linesip.com/libraries/gopipewire/pipewire"
)

func main() {
	pipe := pipewire.Init(os.Args)
	defer pipe.DeInit()
	fmt.Printf(
		"Compiled with libpipewire: %s\nLinked with libpipewire %s\n",
		pipe.GetHeaderVersion(), pipe.GetLibraryVersion(),
	)
}
