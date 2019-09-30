// Package models contains application specific entities.
package models

// Document holds teh Document definition
type Document struct {
	Name       string `json:"name,omitempty"`
	DocumentId string `json:"documentId,omitempty"`
	FolderId   string `json:"folderId,omitempty"`
	CreatedAt  int64  `json:"created_at,omitempty"`
	ModifiedAt int64  `json:"modified_at,omitempty"`
	Owner      string `json:"owner,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	UploadedBy string `json:"uploaded_by,omitempty"`
}
