package models

type ListObjects struct {
	Name         string `json:"name"`
	Size         int    `json:"size"`
	LastModified string `json:"last_modified"`
}
