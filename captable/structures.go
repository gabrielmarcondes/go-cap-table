package captable

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	ID        uuid.UUID `json:"id"`
	Owner     string    `json:"owner"`
	Quantity  int       `json:"quantity"`
	IssueDate time.Time `json:"issueDate"`
}

type NewCertificate struct {
	Owner     string    `json:"owner"`
	Quantity  int       `json:"quantity"`
	IssueDate time.Time `json:"issueDate"`
}

type CapTableOwner struct {
	Owner                  string        `json:"owner"`
	Certificates           []Certificate `json:"certificates"`
	FullyDiluted           int           `json:"fullyDiluted"`
	FullyDilutedPercentage float32       `json:"fullyDilutedPercentage"`
}

type CapTable struct {
	Owners []CapTableOwner `json:"owners"`
	Total  int             `json:"total"`
	AsOf   time.Time       `json:"asOf"`
}
