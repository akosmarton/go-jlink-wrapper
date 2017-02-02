package jlink

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

type JLink struct {
	exePath string
	iface   string
	speed   string
	device  string
}

func NewJLink(exePath, iface, speed, device string) *JLink {
	j := JLink{
		exePath: exePath,
		iface:   iface,
		speed:   speed,
		device:  device,
	}

	return &j
}

func (j *JLink) LoadBin(binPath string, addr int) error {
	var err error

	cmd := exec.Command(j.exePath, "-if", j.iface, "-speed", j.speed, "-device", j.device)

	iow, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	ior, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	iors := bufio.NewReader(ior)

	err = cmd.Start()
	if err != nil {
		return err
	}
	defer cmd.Process.Kill()

	_, err = io.WriteString(iow, "loadbin "+binPath+","+strconv.Itoa(addr)+"\n")
	if err != nil {
		return err
	}

	_, err = io.WriteString(iow, "qc\n")
	if err != nil {
		return err
	}

	success := false

	for {
		var l []byte

		l, _, err = iors.ReadLine()
		if err != nil {
			return err
		}

		if containsError(string(l)) {
			return errors.New(string(l))
		}

		if strings.Contains(string(l), "Verifying") && strings.Contains(string(l), "100%") && strings.Contains(string(l), "Done") {
			success = true
			break
		}
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	if success == false {
		return errors.New("Unknown error")
	}

	return nil
}

func (j *JLink) Erase() error {
	var err error

	cmd := exec.Command(j.exePath, "-if", j.iface, "-speed", j.speed, "-device", j.device)

	iow, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	ior, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	iors := bufio.NewReader(ior)

	err = cmd.Start()
	if err != nil {
		return err
	}
	defer cmd.Process.Kill()

	_, err = io.WriteString(iow, "erase\n")
	if err != nil {
		return err
	}

	_, err = io.WriteString(iow, "qc\n")
	if err != nil {
		return err
	}

	success := false

	for {
		var l []byte

		l, _, err = iors.ReadLine()
		if err != nil {
			return err
		}

		if containsError(string(l)) {
			return errors.New(string(l))
		}

		if strings.Contains(string(l), "Verifying") && strings.Contains(string(l), "100%") && strings.Contains(string(l), "Done") {
			success = true
			break
		}
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	if success == false {
		return errors.New("Unknown error")
	}

	return nil
}

func (j *JLink) Reset() error {
	var err error

	cmd := exec.Command(j.exePath, "-if", j.iface, "-speed", j.speed, "-device", j.device)

	iow, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	ior, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	iors := bufio.NewReader(ior)

	err = cmd.Start()
	if err != nil {
		return err
	}
	defer cmd.Process.Kill()

	_, err = io.WriteString(iow, "r\n")
	if err != nil {
		return err
	}

	_, err = io.WriteString(iow, "qc\n")
	if err != nil {
		return err
	}

	success := false

	for {
		var l []byte

		l, _, err = iors.ReadLine()
		if err != nil {
			return err
		}

		if containsError(string(l)) {
			return errors.New(string(l))
		}

		if strings.Contains(string(l), "Reset type") && strings.Contains(string(l), "Resets") {
			success = true
			break
		}
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	if success == false {
		return errors.New("Unknown error")
	}

	return nil
}

func containsError(s string) bool {
	errorStrings := []string{
		"FAILED",
		"Failed",
		"failed",
		"ERROR",
		"Error",
		"error",
		"WARNING",
		"Warning",
		"warning",
	}

	for _, v := range errorStrings {
		if strings.Contains(s, v) {
			return true
		}
	}

	return false
}
