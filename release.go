package main

import (
	"net/url"
	"strings"
	"time"
)

// Release of a repository tagged via GitHub.
type Release struct {
	ID           string
	Name         string
	Description  string
	URL          url.URL
	PublishedAt  time.Time
	IsDraft      bool
	IsPrerelease bool
}

// IsReleaseCandidate returns true if the release name hints at an RC release.
func (r Release) IsReleaseCandidate() bool {
	return strings.Contains(strings.ToLower(r.Name), "-rc")
}

// IsBeta returns true if the release name hints at a beta version release.
func (r Release) IsBeta() bool {
	return strings.Contains(strings.ToLower(r.Name), "beta")
}

// IsDevelopment returns true if the release name hints at a development version release.
func (r Release) IsDevelopment() bool {
	return strings.Contains(strings.ToLower(r.Name), "dev")
}

// IsNonstable returns true if one of the non-stable release-checking functions return true.
func (r Release) IsNonstable() bool {
	return r.IsReleaseCandidate() || r.IsBeta() || r.IsDevelopment() || r.IsDraft || r.IsPrerelease
}
