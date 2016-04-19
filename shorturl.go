package shorturl

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Service configuration
const (
	domain = "yx.fi"
	idBase = 36
)

// Errors
var (
	ErrNotFound = errors.New("Shorturl not found")
)

// Shorturl database structure
type Shorturl struct {
	ID    int64
	URL   string
	Host  string
	Added time.Time
}

// UID is the base-36 string representation of ID
func (s *Shorturl) UID() string {
	return strconv.FormatInt(s.ID, idBase)
}

// URLString is the shortened URL as string
func (s *Shorturl) URLString() string {
	return "http://" + domain + "/" + s.UID()
}

// PreviewURL is the view that shows where URL directs
func (s *Shorturl) PreviewURL() string {
	return "/p/" + s.UID()
}

// Represent Short URL in pretty format
func (s Shorturl) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s\n", s.URLString())
	fmt.Fprintf(&buf, "  Target: %s\n", truncate(s.URL, 64))
	fmt.Fprintf(&buf, "   Added: %s\n", s.Added.Format("2006-01-02 15:04:05 MST"))
	fmt.Fprintf(&buf, "      IP: %s", s.Host)
	return buf.String()
}