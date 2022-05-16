package linux

type UserData struct {
	ID        uint
	Username  string
	State     string
	Hostname  string
	Timestamp int64
}

type DiskInfo struct {
	ID        uint
	Name      string
	Serial    string
	Make      string
	Model     string
	Timestamp int64
}
