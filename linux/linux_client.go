package linux

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"os/exec"
	"time"
)

var DB *gorm.DB

func LinuxClient() {
	fmt.Println("Starting linux client")

	// Get the database path from the env var
	// or default to a spot in /tmp
	dbPath := os.Getenv("DMCLIENT_DB")
	if dbPath == "" {
		dbPath = "/tmp/dmclient.db"
	}

	// Try to open the database, fail out if an error happens
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		fmt.Println("Error opening the database, closing")
		os.Exit(1)
	}

	// Attempt to migrate the tables if our data structs change
	// Crash if there is an error
	err = db.AutoMigrate(&UserData{}, &DiskInfo{})
	if err != nil {
		fmt.Println("Something went wrong migrating the database, stopping program")
		fmt.Println(err)
	}
	DB = db
	// startScheduler()
	// CheckLogins()
	GetDisks()
}

func RunCommand(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)
	// fmt.Printf("Command: %s\n", command)
	// fmt.Println(cmd)
	output, err := cmd.Output()
	if err != nil {
		return ""
	} else {
		return string(output)
	}
}

func startScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(20).Second().Do(simple_task)
}

func simple_task() {
	fmt.Println("simple_task executed")
}
