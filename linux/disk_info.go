package linux

import (
	"time"

	"os"
	"strings"
)

func GetDisks() {
	diskNames := getDiskNames()
	disks := getDiskInfo(diskNames)
	for _, v := range disks {
		DB.Create(&v)
	}

	// Now parse the disk information
}

func getDiskNames() []string {
	// IMPORTANT -- call lsblk with the -l switch
	// Golang fails to cast stdout to a string with
	// the lsblk partition lines, super annoying...
	stdout := RunCommand("lsblk -l -o NAME,TYPE | grep disk | grep -iv ram | awk '{print $1}'")
	stdout = strings.Trim(stdout, "\n")
	disks := strings.Split(stdout, "\n")
	return disks
}

func getDiskInfo(disks []string) []DiskInfo {
	/* -- Get disk information using gopsutil
	 -- "github.com/shirou/gopsutil/v3/disk"
	data, err := disk.SerialNumber("/dev/" + v)
	if err != nil {
		fmt.Println(err)
	} else {
		diskInfo = append(diskInfo, data)
	}
	*/
	diskInfo := make([]DiskInfo, 0)
	for _, v := range disks {
		serial := RunCommand("hdparm -I /dev/" + v + " | grep 'Serial Number' | xargs | awk '{print $3}'")
		serial = strings.Trim(serial, "\n")
		model := RunCommand("hdparm -I /dev/" + v + "| grep 'Model Number' | xargs")
		model = strings.Trim(model, "\n")
		model = strings.ReplaceAll(model, "Model Number: ", "")
		diskItem := DiskInfo{
			Name:      v,
			Serial:    serial,
			Model:     model,
			Timestamp: time.Now().Unix(),
		}
		diskInfo = append(diskInfo, diskItem)
	}
	return diskInfo
}

func checkForMegaRAID() bool {
	_, err := os.Stat("/opt/MegaRAID/MegaCLI")
	if err != nil {
		return false
	} else {
		return true
	}
}
