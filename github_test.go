package resource_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	resource "github.com/telia-oss/github-pr-resource"
)

func TestTeamMembers(t *testing.T) {
	source := resource.Source{
		Repository:  "git@github.com:cloudfoundry/korifi",
		AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN"),
	}

	client, err := resource.NewGithubClient(&source)
	require.NoError(t, err)

	val, err := client.ListTeamMembers("wg-cf-on-k8s-bots")
	require.NoError(t, err)

	require.Contains(t, val, "korifi-bot")
}
