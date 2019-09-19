// Package models contains application specific entities.
package models

// Links holds the Links definition
type Links struct {
	LinkURL string `json:"link,omitempty"`
}

//LinkToken holds the infromation in a link token
type LinkToken struct {
	FolderID  string
	UserID    string
	AccountID string
}
