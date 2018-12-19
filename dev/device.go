package dev

/*
*/

import (
	"strconv"
	"os/exec"
	"bytes"
	"strings"
)

type Device struct {
     FileSystem string
     TotalKb    int64
     UsedKb     int64
     AvailKb    int64
     MountPoint string
}

func UpdateDeviceList(devs map[string]*Device) {
	
	cmd := exec.Command("df", "-t", "ext4")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	
	str := out.String()
	str2 := strings.Split(str, "\n")

	for line_index, line_str := range str2 {

		if line_index == 0 { // skip the first line
			continue
		}
		
		fields := strings.Fields(line_str)
		
		if len(fields) == 0 { // skip the last brank line
			continue
		}
		
		fs_str := fields[0]
		totalk, _ := strconv.ParseInt(fields[1], 10, 64) // base10, int64
		usedk, _   := strconv.ParseInt(fields[2], 10, 64)
		availk, _  := strconv.ParseInt(fields[3], 10, 64)
		// usedp   := fields[4]
		mount_str  := fields[5]
		
		p := devs[fs_str]
		
		if p == nil {
			// is a new device
			p = &Device{
				FileSystem: fs_str,
				TotalKb: totalk,
				UsedKb: usedk,
				AvailKb: availk,
				MountPoint: mount_str,
			}
		} else {
			// just update
			p.TotalKb = totalk
			p.UsedKb = usedk
			p.AvailKb = availk
		}
		
		devs[fs_str] = p
	}

	for key, v := range devs {
		if v.TotalKb == 0 {
			delete(devs, key)
		}
	}
}

type ByUsage []Device

func (a ByUsage) Len() int      { return len(a) }
func (a ByUsage) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByUsage) Less(i, j int) bool {
	return a[i].FileSystem > a[j].FileSystem
}

func MarkDirtyDeviceList(devs map[string]*Device) {
	for k, _ := range devs {
		devs[k].TotalKb = 0
	}
}
