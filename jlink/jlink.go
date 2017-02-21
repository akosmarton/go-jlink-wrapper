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
	serial  string
	iface   string
	speed   string
	device  string
	cmd     string
}

func NewJLink(exePath, serial, iface, speed, device string) *JLink {
	j := JLink{
		exePath: exePath,
		serial:  serial,
		iface:   iface,
		speed:   speed,
		device:  device,
	}

	return &j
}

func (j *JLink) LoadBin(binPath string, addr int) error {
	var err error

	var args []string

	if j.serial != "" {
		args = append(args, "-SelectEmuBySN")
		args = append(args, j.serial)
	}
	if j.iface != "" {
		args = append(args, "-if")
		args = append(args, j.iface)
	}
	if j.speed != "" {
		args = append(args, "-speed")
		args = append(args, j.speed)
	}
	if j.device != "" {
		args = append(args, "-device")
		args = append(args, j.device)
	}
	cmd := exec.Command(j.exePath, args...)

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

	var args []string

	if j.serial != "" {
		args = append(args, "-SelectEmuBySN")
		args = append(args, j.serial)
	}
	if j.iface != "" {
		args = append(args, "-if")
		args = append(args, j.iface)
	}
	if j.speed != "" {
		args = append(args, "-speed")
		args = append(args, j.speed)
	}
	if j.device != "" {
		args = append(args, "-device")
		args = append(args, j.device)
	}
	cmd := exec.Command(j.exePath, args...)

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

	var args []string

	if j.serial != "" {
		args = append(args, "-SelectEmuBySN")
		args = append(args, j.serial)
	}
	if j.iface != "" {
		args = append(args, "-if")
		args = append(args, j.iface)
	}
	if j.speed != "" {
		args = append(args, "-speed")
		args = append(args, j.speed)
	}
	if j.device != "" {
		args = append(args, "-device")
		args = append(args, j.device)
	}
	cmd := exec.Command(j.exePath, args...)

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

func (j *JLink) GetEmuList() ([]string, error) {
	var err error

	var args []string

	if j.serial != "" {
		args = append(args, "-SelectEmuBySN")
		args = append(args, j.serial)
	}
	if j.iface != "" {
		args = append(args, "-if")
		args = append(args, j.iface)
	}
	if j.speed != "" {
		args = append(args, "-speed")
		args = append(args, j.speed)
	}
	if j.device != "" {
		args = append(args, "-device")
		args = append(args, j.device)
	}
	cmd := exec.Command(j.exePath, args...)

	iow, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	ior, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	iors := bufio.NewReader(ior)

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	defer cmd.Process.Kill()

	_, err = io.WriteString(iow, "ShowEmuList\n")
	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(iow, "qc\n")
	if err != nil {
		return nil, err
	}

	var serials []string

	for {
		var l []byte

		l, _, err = iors.ReadLine()
		if err != nil {
			return nil, err
		}

		if containsError(string(l)) {
			return nil, errors.New(string(l))
		}

		if strings.Contains(string(l), "Serial number: ") {
			s := strings.Split(string(l), " ")
			if s[3] == "Serial" && s[4] == "number:" {
				serials = append(serials, strings.Trim(s[5], ","))
			}
			break
		}
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	if len(serials) == 0 {
		return nil, errors.New("No emulator found")
	}

	return serials, nil
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
