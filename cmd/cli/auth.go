/*
Copyright © 2022 Johnson Shi <Johnson.Shi@microsoft.com>
*/
package main

import (
	"context"
	"fmt"

	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
)

type registryRepositoryAuthOpts struct {
	username   string
	password   string
	registry   string
	repository string
}

func (opts *registryRepositoryAuthOpts) getAuthenticatedRemoteRepositoryClient() (*remote.Repository, error) {
	// Create a client to the remote repository identified by a reference.
	repo, err := remote.NewRepository(fmt.Sprintf("%s/%s", opts.registry, opts.repository))
	if err != nil {
		return nil, err
	}

	// Set the repository auth credential client.
	repo.Client = &auth.Client{
		Credential: func(ctx context.Context, reg string) (auth.Credential, error) {
			return auth.Credential{
				Username: opts.username,
				Password: opts.password,
			}, nil
		},
	}

	return repo, nil
}
