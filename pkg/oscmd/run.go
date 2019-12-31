package oscmd

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

type OsCmd struct {
	name  string
	args  []string
	stdin string
}

func OsRun(oscmd OsCmd) error {
	cmd := exec.Command(oscmd.name, oscmd.args...)
	if oscmd.stdin != "" {
		cmd.Stdin = strings.NewReader(oscmd.stdin)
	}
	var outbuf bytes.Buffer
	var errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Run()
	log.Printf("out:\n%v\n", outbuf.String())
	if errbuf.Len() > 0 {
		log.Printf("err:\n%v\n", errbuf.String())
	}
	return err
}
