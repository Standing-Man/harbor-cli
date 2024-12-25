package api

import (
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/securityhub"
	"github.com/goharbor/harbor-cli/pkg/utils"
)

func ListVulnerability(opts ...ListFlags) (*securityhub.ListVulnerabilitiesOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return &securityhub.ListVulnerabilitiesOK{}, err
	}

	var listFlags ListFlags
	if len(opts) > 0 {
		listFlags = opts[0]
	}

	response, err := client.Securityhub.ListVulnerabilities(ctx,
		&securityhub.ListVulnerabilitiesParams{
			Page:     &listFlags.Page,
			PageSize: &listFlags.PageSize,
		})
	if err != nil {
		return &securityhub.ListVulnerabilitiesOK{}, err
	}
	return response, nil
}