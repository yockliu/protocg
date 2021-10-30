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
	Deps struct {
		Local  []string
		Remote []Remote
	}
}

func GenV1() {
	generator := parseConfig()
	depsPaths := []string{}
	srcPaths := []string{}
	fmt.Println(generator)

	for _, remote := range generator.Src.Remote {
		if err := remote.clone(); err != nil {
			panic(err)
		}

		subPaths := remote.getSubPaths()

		for _, subPath := range subPaths {
			depsPaths = append(depsPaths, subPath)
			srcPaths = append(srcPaths, subPath)
		}
	}

	for _, local := range generator.Src.Local {
		depsPaths = append(depsPaths, local)
		srcPaths = append(srcPaths, local)
	}

	for _, remote := range generator.Deps.Remote {
		if err := remote.clone(); err != nil {
			panic(err)
		}
		subPaths := remote.getSubPaths()

		for _, subPath := range subPaths {
			depsPaths = append(depsPaths, subPath)
		}
	}

	for _, local := range generator.Deps.Local {
		depsPaths = append(depsPaths, local)
	}

	fmt.Println("deps path = ", depsPaths)
	fmt.Println("src path = ", srcPaths)

	execProtoc(depsPaths, generator.Opts, srcPaths)
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
