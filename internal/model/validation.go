package model

import (
	"path/filepath"
	"strings"
)

func NormalizePath(p string) string { return filepath.Clean(strings.TrimSpace(p)) }
func IsSafePath(p string) bool      { return p != "" && !strings.HasPrefix(NormalizePath(p), "..") }
func (a Album) Valid() bool         { return a.ID != "" && a.Title != "" && len(a.AccessCode) >= 4 }
func (a Album) ActiveCount() int {
	n := 0
	for _, r := range a.Photos {
		if !r.Archived {
			n++
		}
	}
	return n
}
