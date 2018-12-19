package main

import (
	"fmt"
	"github.com/apcera/termtables"
	"github.com/tttamaki/cluster-df/dev"
	"os"
	"sort"
	"time"
)

func minimum(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type Cluster struct {
	Nodes []Node `json:"nodes"`
}


type Device struct {
     FileSystem string
     TotalKb    int64
     UsedKb     int64
     AvailKb    int64
     MountPoint string
}


type Node struct {
	Name      string
	Devices   []Device
	Time      time.Time `json:"time"` // current timestamp from message
}

func InitNode(n *Node) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	n.Name = name
}

type ByUsage []Device

func (a ByUsage) Len() int      { return len(a) }
func (a ByUsage) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByUsage) Less(i, j int) bool {
	return a[i].FileSystem > a[j].FileSystem
}

type ByName []Node

func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (c *Cluster) Sort() {
	sort.Sort(ByName(c.Nodes))
}


func K2MG(kb int64) string {
	if int(kb / 1024 / 1024) > 0 {
		return fmt.Sprintf("%d GB", int(kb / 1024 / 1024))
	} else if int(kb / 1024) > 0 {
		return fmt.Sprintf("%d MB", int(kb / 1024))
	} else {
		return fmt.Sprintf("%d kB", kb)
	}
}



func (c *Cluster) Print(show_time bool) {

	table := termtables.CreateTable()

	tableHeader := []interface{}{"Node", "Filesystem", "Total", "Used", "Avail", "Use%", "mount"}
	if show_time {
		tableHeader = append(tableHeader, "Last Seen")
	}
	table.AddHeaders(tableHeader...)

	for n_id, n := range c.Nodes {

		node_lastseen := n.Time.Format("Mon Jan 2 15:04:05 2006")

		if len(n.Devices) == 0 {

			tableRow := []interface{}{
				n.Name,
				"",
				"",
				"",
				"",
				"",
				"",
			}

			if show_time {
				tableRow = append(tableRow, node_lastseen)
			}
			table.AddRow(tableRow...)

		} else {
			name := ""
			named := false
			for p_id, p := range n.Devices {

				if ! named {
					name = n.Name
					named = true
				} else {
					name = ""
				}

				tableRow := []interface{}{
					name,
					p.FileSystem,
					K2MG(p.TotalKb),
					K2MG(p.UsedKb),
					K2MG(p.AvailKb),
					fmt.Sprintf("%.1f", float64(p.UsedKb) / float64(p.TotalKb) * 100),
					p.MountPoint,
				}

				if show_time {
					if p_id == 0 {
						tableRow = append(tableRow, node_lastseen)

					} else {
						tableRow = append(tableRow, "")

					}
				}

				table.AddRow(tableRow...)
				table.SetAlign(termtables.AlignRight, 3)
				table.SetAlign(termtables.AlignRight, 4)
				table.SetAlign(termtables.AlignRight, 5)
				table.SetAlign(termtables.AlignRight, 6)
			}
		}
		if n_id < len(c.Nodes)-1 {
			table.AddSeparator()
		}

	}
	fmt.Printf("\033[2J")
	fmt.Println(time.Now().Format("Mon Jan 2 15:04:05 2006") + " (http://github.com/tttamaki/cluster-df)")
	fmt.Println(table.Render())
}

func GetDevices(devs map[string]*dev.Device, max int) []Device {

	var m_display []Device

	for _, v := range devs {
		if v.TotalKb > 0 {
			copy_dev := Device{
				v.FileSystem,
				v.TotalKb,
				v.UsedKb,
				v.AvailKb,
				v.MountPoint,
			}
			m_display = append(m_display, copy_dev)
		}
	}

	sort.Sort(ByUsage(m_display))
	m_display = m_display[:minimum(max, len(m_display))]

	return m_display

}
