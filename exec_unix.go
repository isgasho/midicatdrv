// +build !windows

package midicatdrv

import (
	"fmt"
	"os/exec"
)

func _execCommand(c string) *exec.Cmd {
	return exec.Command("/bin/sh", "-c", "exec "+c)
}

func midiCatOutCmd(index int) *exec.Cmd {
	cmd := exec.Command("midicat", "out", fmt.Sprintf("--index=%v", index))
	//	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return cmd
}

func midiCatInCmd(index int) *exec.Cmd {
	cmd := exec.Command("midicat", "in", fmt.Sprintf("--index=%v", index))
	//	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return cmd
}

func midiCatCmd(args string) *exec.Cmd {
	cmd := _execCommand("midicat " + args)
	// important! prevents that signals such as interrupt send to the main program gets passed
	// to midicat (which would not allow us to shutdown properly, e.g. stop hanging notes)
	//cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return cmd
}
