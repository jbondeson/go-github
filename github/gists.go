// Copyright 2013 Google. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

package github

import (
	"fmt"
	"time"
)

// GistsService handles communication with the gist related
// methods of the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/gists/
type GistsService struct {
	client *Client
}

// Gist represents a GitHub gist
type Gist struct {
	ID          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Public      bool   `json:"public,omitempty"`
	Owner       *User  `json:"owner,omitempty"`
	//TODO: Files
	CreatedAt *time.Time `json:"created_at,omitempty"`
	//TODO: Forks
	//TODO: History
}

// Get fetches a gist.
//
// GitHub API docs: http://developer.github.com/v3/gists/#get-a-single-gist
func (s *GistsService) Get(id int) (*Gist, error) {
	g := fmt.Sprintf("gists/%n", id)

	req, err := s.client.NewRequest("GET", g, nil)
	if err != nil {
		return nil, err
	}

	gResp := new(Gist)
	_, err = s.client.Do(req, gResp)
	return gResp, err
}
