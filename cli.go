package main

import (
	"os"
	"path/filepath"

	"github.com/creativeprojects/diskfree/disk"
	"github.com/pterm/pterm"
)

func runCli() {
	disks, err := os.ReadDir("/Volumes")
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	tableData := pterm.TableData{
		{"Disk", "Size", "Used", "Free", "Available", "Usage"},
	}
	for _, diskEntry := range disks {
		diskPath := filepath.Join("/Volumes", diskEntry.Name())
		usage, err := disk.New(diskPath)
		if err != nil {
			pterm.Error.Println(err)
			return
		}
		tableData = append(tableData, []string{
			diskPath,
			displayWithUnit(usage.Size()),
			displayWithUnit(usage.Used()),
			displayWithUnit(usage.Free()),
			displayWithUnit(usage.Available()),
			pterm.Sprintf("%.2f %s", usage.Usage()*100, "%"),
		})
		pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
	}
}

func displayWithUnit(intValue uint64) string {
	value, unit := minimizeDisplay(intValue)
	return pterm.Sprintf("%.2f %s", value, unit)
}
