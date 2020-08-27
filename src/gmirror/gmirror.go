package gmirror

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"strings"
)

//GMirror
type GMirror struct {
}

//GMirrorStatus
type GMirrorStatus struct {
	Mirrors map[string]GMirrorMirror
}

type GMirrorMirror struct {
	Name    string
	State   string
	Devices map[string]GMirrorDevice
}

type GMirrorDevice struct {
	Name  string
	State string
}

func NewGMirror() GMirror {
	return GMirror{}
}

//Status returns gmirror status
func (g *GMirror) Status() (*GMirrorStatus, error) {
	/*
		root  COMPLETE  ada0p2 (ACTIVE)
		root  COMPLETE  ada1p2 (ACTIVE)
		swap  COMPLETE  ada0p3 (ACTIVE)
		swap  COMPLETE  ada1p3 (ACTIVE)
	*/
	out, err := sh.Command("gmirror", "status", "-gs").Output()
	result := string(out)
	fmt.Printf("%s\n", result)

	if err != nil {
		return nil, err
	}

	status := GMirrorStatus{
		Mirrors: map[string]GMirrorMirror{},
	}

	lines := strings.Split(result, "\n")
	fmt.Printf("lines %v\n", lines)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		fmt.Printf("parts %v\n", parts)
		var cols []string
		for _, part := range parts {
			if part != "" {
				cols = append(cols, part)
			}
		}
		fmt.Printf("cols %v\n", cols)

		if len(cols) != 4 {
			continue
		}

		mirror := cols[0]
		mirrorState := cols[1]
		device := cols[2]
		deviceState := cols[3]

		if _, ok := status.Mirrors[mirror]; !ok {
			status.Mirrors[mirror] = GMirrorMirror{
				Name:    mirror,
				State:   mirrorState,
				Devices: map[string]GMirrorDevice{},
			}
		}

		if _, ok := status.Mirrors[mirror].Devices[device]; !ok {
			status.Mirrors[mirror].Devices[device] = GMirrorDevice{
				Name:  device,
				State: deviceState,
			}
		}
	}

	fmt.Printf("%v\n", status)

	return &status, nil
}
