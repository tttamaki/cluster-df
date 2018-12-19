package main

import (
	"github.com/tttamaki/cluster-df/dev"
	"github.com/pebbe/zmq4"
	"github.com/vmihailenco/msgpack"
	"log"
	"time"
)

// hash map of all devices
var work_devices map[string]*dev.Device

var clus Cluster

func main() {

	work_devices = make(map[string]*dev.Device)

	// load ports and ip-address
	cfg := LoadConfig()
	cfg.Print()

	// sending messages (PUSH-PULL)
	SocketAddr := "tcp://" + cfg.RouterIp + ":" + cfg.Ports.Nodes
	log.Println("Now pushing to", SocketAddr)
	socket, err := zmq4.NewSocket(zmq4.PUSH)
	if err != nil {
		panic(err)
	}
	defer socket.Close()
	socket.Connect(SocketAddr)

	node := Node{}
	InitNode(&node)
	clus.Nodes = append(clus.Nodes, node)

	for {
		// reset most devices
		dev.MarkDirtyDeviceList(work_devices)

		dev.UpdateDeviceList(work_devices)

		clus.Nodes[0].Devices = GetDevices(work_devices, cfg.MaxDisplay)

		clus.Nodes[0].Time = time.Now()

		// encode data
		msg, err := msgpack.Marshal(&clus.Nodes[0])
		if err != nil {
			log.Fatal("encode error:", err)
			panic(err)
		}

		// send data
		socket.SendBytes(msg, 0)

		time.Sleep(time.Duration(cfg.Tick) * time.Second)

	}

}
