package core

import "fmt"

type Generator struct {
	Version string
	Name    string
	Opts    []string
	Src     struct {
		Local  []string
		Remote []Remote
	}
	ProtoPath struct {
		Local  []string
		Remote []Remote
	} `mapstructure:"proto_path"`
}

func GenV1() {
	generator := parseConfig()
	protoPaths := []string{}
	srcPaths := []string{}
	fmt.Println(generator)

	for _, local := range generator.ProtoPath.Local {
		protoPaths = append(protoPaths, local)
	}

	for _, remote := range generator.ProtoPath.Remote {
		if err := remote.clone(); err != nil {
			panic(err)
		}
		subPaths := remote.getSubPaths()

		for _, subPath := range subPaths {
			protoPaths = append(protoPaths, subPath)
		}
	}

	for _, local := range generator.Src.Local {
		srcPaths = append(srcPaths, local)
	}

	for _, remote := range generator.Src.Remote {
		if err := remote.clone(); err != nil {
			panic(err)
		}

		subPaths := remote.getSubPaths()

		for _, subPath := range subPaths {
			srcPaths = append(srcPaths, subPath)
		}
	}

	fmt.Println("deps path = ", protoPaths)
	fmt.Println("src path = ", srcPaths)

	execProtoc(protoPaths, generator.Opts, srcPaths)
}

func (r *Remote) getSubPaths() []string {
	if len(r.Path) < 1 {
		return []string{*r.getClonePath()}
	} else {
		result := []string{}
		for _, p := range r.Path {
			result = append(result, *r.getClonePath()+"/"+p)
		}
		return result
	}
}
