package core

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
)

var remoteCloneMap = map[string]*string{}

type Remote struct {
	Repo string
	Refs string
	Path []string
}

func (r *Remote) getRefs() string {
	if r.Refs != "" {
		return r.Refs
	} else {
		return "master"
	}
}

func (r *Remote) sig() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(r.Repo+r.Refs)))[:8]
}

func (r *Remote) clone() error {
	sig := r.sig()

	if remoteCloneMap[sig] != nil {
		return nil
	}

	path := cachePathOf(sig)

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	cmd := exec.Command("git", "clone", "-b", r.getRefs(), r.Repo, path)
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", string(out))
	if err != nil {
		fmt.Println("err = ", err)
		return err
	}

	remoteCloneMap[sig] = &path
	return nil
}

func (r *Remote) getClonePath() *string {
	return remoteCloneMap[r.sig()]
}
