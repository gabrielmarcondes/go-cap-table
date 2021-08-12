package captable

import (
	"time"

	"github.com/google/uuid"
)

var certificates = []Certificate{
	{ID: uuid.New(), Owner: "Gabriel", Quantity: 100000, IssueDate: time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Victor", Quantity: 200000, IssueDate: time.Date(2011, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Rajat", Quantity: 300000, IssueDate: time.Date(2012, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Robert", Quantity: 400000, IssueDate: time.Date(2013, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Marco", Quantity: 500000, IssueDate: time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Gabriel", Quantity: 900000, IssueDate: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Victor", Quantity: 800000, IssueDate: time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Rajat", Quantity: 700000, IssueDate: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Robert", Quantity: 600000, IssueDate: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.New(), Owner: "Marco", Quantity: 500000, IssueDate: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)},
}

var certificatesDatabase = map[uuid.UUID]Certificate{}

func init() {
	for _, certificate := range certificates {
		certificatesDatabase[certificate.ID] = certificate
	}
}
