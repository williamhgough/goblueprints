package backup

import (
	"fmt"
	"path/filepath"
	"time"
)

// Monitor holds the paths and associated hashes,
// the current achiver in use and where to hold
// the archives
type Monitor struct {
	Paths       map[string]string
	Archiver    Archiver
	Destination string
}

// Now iterates over every path in a Monitors paths and
// generates the latest hash of that.
func (m *Monitor) Now() (int, error) {
	var counter int
	for path, lastHash := range m.Paths {
		newHash, err := DirHash(path)
		if err != nil {
			return 0, err
		}

		if newHash != lastHash {
			err := m.act(path)
			if err != nil {
				return 0, err
			}
			m.Paths[path] = newHash // update the hash
			counter++
		}
	}
	return counter, nil
}

func (m *Monitor) act(path string) error {
	dirname := filepath.Base(path)
	filename := fmt.Sprintf(m.Archiver.DestFmt(), time.Now().UnixNano())
	return m.Archiver.Archive(path, filepath.Join(m.Destination, dirname, filename))
}
