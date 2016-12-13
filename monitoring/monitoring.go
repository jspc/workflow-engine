package monitoring

import (
    "os"
)

// Mon ...
// Monitoring/ Metadata container
type Mon struct {
    // Job node's hostname
    Hostname            string
    EUID, GID, PID, UID int
    UTF8                string
}

// NewMon ...
// Initialise and return a `Mon`,
// containing monitoring and metadata
func NewMon() (m Mon) {
    h, err := os.Hostname()
    if err != nil {
        m.Hostname = err.Error()
    } else {
        m.Hostname = h
    }

    m.UTF8 = "✔"

    m.EUID = os.Geteuid()
    m.GID = os.Getgid()
    m.PID = os.Getpid()
    m.UID = os.Getuid()

    return
}
