// Package gitea deals with using gitea as a store for pushed applications and their deployment info
package gitea

import (
	giteaSDK "code.gitea.io/sdk/gitea"
	"github.com/epinio/epinio/deployments"
	"github.com/epinio/epinio/helpers/kubernetes"
	"github.com/pkg/errors"
)

// Client provides functionality for talking to a
// Gitea installation on Kubernetes
type Client struct {
	Client *giteaSDK.Client
	Auth   deployments.GiteaAuth
}

const (
	GiteaCredentialsSecret = "gitea-creds"
)

var clientMemo *Client

// New loads the config and returns a new gitea client
func New() (*Client, error) {
	if clientMemo != nil {
		return clientMemo, nil
	}

	cluster, err := kubernetes.GetCluster()
	if err != nil {
		return nil, err
	}

	// See also deployments/gitea.go (service, func `apply`).
	// See also deployments/tekton.go, func `createGiteaCredsSecret`
	auth, err := getGiteaCredentials(cluster)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve gitea credentials")
	}

	client, err := giteaSDK.NewClient(deployments.GiteaURL)
	if err != nil {
		return nil, errors.Wrap(err, "gitea client creation failed")
	}

	client.SetBasicAuth(auth.Username, auth.Password)

	c := &Client{
		Client: client,
		Auth:   *auth,
	}

	clientMemo = c

	return c, nil
}

// getGiteaCredentials resolves Gitea's credentials
func getGiteaCredentials(cluster *kubernetes.Cluster) (*deployments.GiteaAuth, error) {
	// See deployments/tekton.go, func `createGiteaCredsSecret`
	// for where `install` configures tekton for the credentials
	// retrieved here.
	//
	// See deployments/gitea.go func `apply` where `install`
	// configures gitea for the same credentials.
	s, err := cluster.GetSecret(deployments.TektonStagingNamespace, GiteaCredentialsSecret)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read gitea credentials")
	}

	username, ok := s.Data["username"]
	if !ok {
		return nil, errors.Wrap(err, "username key not found in gitea credentials secret")
	}

	password, ok := s.Data["password"]
	if !ok {
		return nil, errors.Wrap(err, "password key not found in gitea credentials secret")
	}

	return &deployments.GiteaAuth{
		Username: string(username),
		Password: string(password),
	}, nil
}

func (c *Client) DeleteRepo(org, repo string) error {
	_, err := c.Client.DeleteRepo(org, repo)

	return err
}

func (c *Client) CreateOrg(org string) error {
	_, _, err := c.Client.CreateOrg(giteaSDK.CreateOrgOption{
		Name: org,
	})

	return err
}

func (c *Client) DeleteOrg(org string) error {
	_, err := c.Client.DeleteOrg(org)

	return err
}
