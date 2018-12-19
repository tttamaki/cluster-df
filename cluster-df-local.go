package main

import (
	"flag"
	"time"
	"github.com/tttamaki/cluster-df/dev"
)

// hash map of all devices
var work_devices map[string]*dev.Device

var clus Cluster

func main() {

	showTimePtr := flag.Bool("t", false, "show time of events")
	flag.Parse()

	work_devices = make(map[string]*dev.Device)

	// load ports and ip-address
	cfg := LoadConfig()

	node := Node{}
	InitNode(&node)
	clus.Nodes = append(clus.Nodes, node)

	for {
		// reset most devices
		dev.MarkDirtyDeviceList(work_devices)

		dev.UpdateDeviceList(work_devices)

		clus.Nodes[0].Devices = GetDevices(work_devices, cfg.MaxDisplay)

		clus.Nodes[0].Time = time.Now()

		clus.Print(*showTimePtr)

		time.Sleep(time.Duration(cfg.Tick) * time.Second)
	}

}
