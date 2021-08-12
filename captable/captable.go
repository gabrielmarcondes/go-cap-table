package captable

import (
	"sort"
	"time"

	"github.com/google/uuid"
)

func GetCertificates() []Certificate {
	certs := []Certificate{}
	for _, cert := range certificatesDatabase {
		certs = append(certs, cert)
	}
	return certs
}

func GetCertificateByID(id uuid.UUID) Certificate {
	return certificatesDatabase[id]
}

func AddCertificate(newCert NewCertificate) Certificate {
	newId := uuid.New()

	certificatesDatabase[newId] = Certificate{
		ID:        newId,
		Owner:     newCert.Owner,
		Quantity:  newCert.Quantity,
		IssueDate: newCert.IssueDate,
	}

	return certificatesDatabase[newId]
}

func GetCapTable(asOf time.Time) CapTable {
	certificatesByOwner := make(map[string][]Certificate)
	totalByOwner := make(map[string]int)
	totalFullyDiluted := 0

	for _, certificate := range certificatesDatabase {
		if certificate.IssueDate.After(asOf) {
			continue
		}

		certificatesByOwner[certificate.Owner] = append(certificatesByOwner[certificate.Owner], certificate)
		totalByOwner[certificate.Owner] += certificate.Quantity
		totalFullyDiluted += certificate.Quantity
	}

	owners := []CapTableOwner{}
	for owner, certificates := range certificatesByOwner {
		newOwner := CapTableOwner{
			Owner: owner, Certificates: certificates, FullyDiluted: totalByOwner[owner], FullyDilutedPercentage: float32(totalByOwner[owner]) / float32(totalFullyDiluted),
		}
		owners = append(owners, newOwner)
	}

	sort.Slice(owners, func(i, j int) bool { return owners[i].FullyDiluted > owners[j].FullyDiluted })

	capTable := CapTable{Total: totalFullyDiluted, Owners: owners, AsOf: asOf}

	return capTable
}
