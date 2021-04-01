package date

import (
	"os"
	"strings"
	"time"
)

// Date date man
type Date struct {
}

// FormatFolderMigration format to folder string
func (d *Date) FormatFolderMigration(time time.Time) string {
	return time.Format("01-02-2006T15M04M05")
}

// GetNowDate get now date
func (d *Date) GetNowDate(dates []time.Time) time.Time {
	now := dates[0]

	for index, date := range dates {
		if date.After(now) && index != 1 {
			now = date
		}
	}

	return now
}

// ConvertFileInfoTime convert fileinfo
func (d *Date) ConvertFileInfoTime(fileInfos []os.FileInfo) []time.Time {
	var times []time.Time
	layout := "01-02-2006 15:04:05"

	for _, f := range fileInfos {
		stringTime := strings.ReplaceAll(strings.ReplaceAll(f.Name(), "T", " "), "M", ":")
		t, err := time.Parse(layout, stringTime)

		if err == nil {
			times = append(times, t)
		}
	}
	return times
}

// NewDate Ioc
func NewDate() *Date {
	return &Date{}
}
