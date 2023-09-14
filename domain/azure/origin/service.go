package origin

import (
	"fmt"

	"github.com/cgxarrie-go/prq/domain/errors"
	"github.com/cgxarrie-go/prq/domain/ports"
	"github.com/cgxarrie-go/prq/utils"
	"github.com/muesli/termenv"
)

type service struct{}

// CreatePRsURL implements ports.OriginSvc.
func (s service) CreatePRsURL(o utils.Origin) (url string, err error) {
	if !o.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	azOrigin := NewAzureOrigin(o)
	base := s.baseUrl(azOrigin)
	url = fmt.Sprintf("%s/repositories/%s/pullrequests?api-version=7.0"+
		"&supportsIterations=true", base, azOrigin.Repository())

	return
}

// GetPRsURL implements ports.OriginSvc.
func (s service) GetPRsURL(o utils.Origin, status ports.PRStatus) (
	url string, err error) {

	if !o.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	azOrigin := NewAzureOrigin(o)
	base := s.baseUrl(azOrigin)
	url = fmt.Sprintf("%s/repositories/%s/pullrequests?api-version=7.0"+
		"&searchCriteria.status=%d", base, azOrigin.Repository(), status)

	return
}

// PRLink implements ports.OriginSvc.
func (s service) PRLink(o utils.Origin, id, text string) (
	url string, err error) {

	if !o.IsAzure() {
		return url, errors.NewErrInvalidRepositoryType(o)
	}

	azOrigin := NewAzureOrigin(o)
	base := s.baseUrl(azOrigin)
	url = fmt.Sprintf("%s/%s/pullrequest/%s", base, azOrigin.Repository(), id)
	return termenv.Hyperlink(url, text), nil
}

func NewService() ports.OriginSvc {
	return service{}
}

func (s service) baseUrl(o AzureOrigin) string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git", 
	o.Organizaion(), 
	o.Project())
}
