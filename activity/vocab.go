package activity

import (
	"net/url"
)

const context = "https://www.w3.org/ns/activitystreams"

// A Group actor
// Ref: https://www.w3.org/TR/activitystreams-vocabulary/#dfn-organization
type Group struct {
	Context string   `json:"@context"`
	ID      *url.URL `json:"id"`
	Name    string   `json:"name"`
	Type    string   `json:"type"`
}

// NewOrganization creates an ActivityStreams Group
func NewGroup(name string, id *url.URL) Group {
	return Group{
		Context: context,
		Name:    name,
		ID:      id,
		Type:    "Group",
	}
}

// An Organization actor
// Ref: https://www.w3.org/TR/activitystreams-vocabulary/#dfn-organization
type Organization struct {
	Context string   `json:"@context"`
	ID      *url.URL `json:"id"`
	Name    string   `json:"name"`
	Type    string   `json:"type"`
}

// NewOrganization creates an ActivityStreams Organization
func NewOrganization(name string, id *url.URL) Organization {
	return Organization{
		Context: context,
		Name:    name,
		ID:      id,
		Type:    "Organization",
	}
}
