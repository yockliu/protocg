package core

import (
	"fmt"
	"os/exec"
	"strings"
)

func execProtoc(deps []string, opts []string, src []string) {
	cmdOpts := []string{}

	cmdOpts = append(cmdOpts, "protoc")

	for _, pp := range deps {
		cmdOpts = append(cmdOpts, "--proto_path="+pp)
	}

	for _, opt := range opts {
		cmdOpts = append(cmdOpts, opt)
	}

	for _, src := range src {
		if strings.HasSuffix(src, ".proto") {
			cmdOpts = append(cmdOpts, src)
		} else {
			cmdOpts = append(cmdOpts, src+"/*.proto")
		}
	}

	cmd := exec.Command("sh", "-c", strings.Join(cmdOpts, " "))
	fmt.Println(cmd)

	out, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", string(out))
	if err != nil {
		panic(err)
	}
}
