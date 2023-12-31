package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/creativeprojects/diskfree/disk"
	"github.com/creativeprojects/diskfree/icon"
	"github.com/getlantern/systray"
)

var diskItems []*systray.MenuItem = make([]*systray.MenuItem, maxDiskItems)
var maxDiskItems = 10

func runSystray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("")

	for i := 0; i < maxDiskItems; i++ {
		diskItems[i] = systray.AddMenuItem("_", "")
		diskItems[i].Hide()
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	go func() {
		for {
			disks, err := os.ReadDir("/Volumes")
			if err != nil {
				diskItems[0].SetTitle(err.Error())
				diskItems[0].Show()
				diskItems[0].Disable()
			}
			totalAvailable := uint64(0)
			for diskIndex, diskEntry := range disks {
				diskPath := filepath.Join("/Volumes", diskEntry.Name())
				usage, err := disk.New(diskPath)
				if err != nil {
					diskItems[diskIndex].SetTitle(err.Error())
					diskItems[diskIndex].Show()
					diskItems[diskIndex].Disable()
					continue
				}
				avail, availUnit := minimizeDisplay(usage.Available())
				capacity, capacityUnit := minimizeDisplay(usage.Size())
				diskItems[diskIndex].SetTitle(fmt.Sprintf("%s: %.2f%s free (%.2f%s)",
					diskEntry.Name(),
					avail, availUnit,
					capacity, capacityUnit,
				))
				diskItems[diskIndex].Show()

				totalAvailable += usage.Available()
			}
			// hide remaining slots
			for i := len(disks); i < maxDiskItems; i++ {
				diskItems[i].Hide()
			}
			// change title with total available
			total, totalUnit := minimizeDisplay(totalAvailable)
			systray.SetTitle(fmt.Sprintf("%.2f%s", total, totalUnit))
			time.Sleep(15 * time.Second)
		}
	}()
}

func onExit() {
}
