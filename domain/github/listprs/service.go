package listprs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/cgxarrie-go/prq/domain/github/origin"
	"github.com/cgxarrie-go/prq/domain/models"
	"github.com/cgxarrie-go/prq/domain/ports"
)

type service struct {
	pat string
	originSvc ports.OriginSvc
}

// NewService return new instnce of github service
func NewService(pat string, originSvc ports.OriginSvc) ports.PRReader {
	return service{
		pat: fmt.Sprintf("`:%s", pat),
		originSvc: originSvc,
	}
}

// GetPRs implements ports.ProviderService
func (svc service) GetPRs(req ports.ListPRRequest) (
	prs []models.PullRequest, err error) {

	g := errgroup.Group{}

	for _, o := range req.Origins() {
		ghOrigin := origin.NewGithubOrigin(o)
		url, err := svc.originSvc.GetPRsURL(o, req.Status())
		if err != nil {
			return prs, fmt.Errorf("gettig url from origin %s: %w", 
			o, err)
		}		
		g.Go(func() error {
			ghPRs, err := svc.getData(url)
			if err == nil {
				for _, ghPR := range ghPRs {
					pr := ghPR.ToPullRequest(ghOrigin.User())
					pr.Link, err = svc.originSvc.PRLink(ghOrigin.Origin, pr.ID, 
						"open")
					prs = append(prs, pr)
				}				
			}
			return err
		})

	}

	return prs, g.Wait()
}

func (svc service) getData(url string) (
	prs []ResponsePullRequest, err error) {

	err = svc.doGet(url, &prs)
	return
}

func (svc service) doGet(url string, resp interface{}) (err error) {
	b64PAT := base64.RawStdEncoding.EncodeToString([]byte(svc.pat))
	bearer := fmt.Sprintf("Basic %s", b64PAT)

	ghReq, err := http.NewRequest("GET", url, nil)
	ghReq.Header.Add("Authorization", bearer)
	ghReq.Header.Add("Accept", "application/vnd.github+json")
	ghReq.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	azResp, err := client.Do(ghReq)
	if err != nil {
		return err
	}

	if azResp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d - %s", azResp.StatusCode, azResp.Status)
	}

	defer azResp.Body.Close()
	return json.NewDecoder(azResp.Body).Decode(resp)

}