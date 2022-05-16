package linux

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CheckLogins() {
	// Get user sessions and trim the output
	stdout := RunCommand("loginctl --no-legend | grep -v root | awk '{print $1}'")
	stdout = strings.Trim(stdout, "\n")

	// Make an array of the sessions
	userSessions := strings.Split(stdout, "\n")

	userSlice := make([]UserData, 0)

	for _, v := range userSessions {
		userSlice = append(userSlice, getSession(v))
	}

	for _, v := range userSlice {
		fmt.Printf("Saving %+v\n", v)
		DB.Create(&v)
	}

}

func getSession(id string) UserData {
	// use RunCommand to get the individual user session
	stdout := RunCommand("loginctl show-session " + id)
	stdout = strings.Trim(stdout, "\n")

	// Make an array of the output
	output := strings.Split(stdout, "\n")

	// Pass the array to parse_session_into_data()
	return parseSessionIntoData(output)
}

func parseSessionIntoData(session []string) UserData {
	// Create a new map<string, string>
	userMap := make(map[string]string)

	// Loop over the session array, split it, then add to the map
	for _, v := range session {
		tmp := strings.Split(v, "=")
		userMap[tmp[0]] = tmp[1]
	}

	// Get the hostname
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "hostnameError"
	}

	// Parse the user state
	state := getState(userMap)

	// Create a new UserData struct and return it
	user := UserData{
		Username:  userMap["Name"],
		State:     state,
		Hostname:  hostname,
		Timestamp: time.Now().Unix(),
	}
	return user
}

func getState(user map[string]string) string {
	if strings.Compare(user["Active"], "yes") == 0 {
		if strings.Compare(user["IdleHint"], "no") == 0 {
			if strings.Compare(user["LockedHint"], "no") == 0 {
				return "active"
			}
		}
	}
	return "locked"
}
