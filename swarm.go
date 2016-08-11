// Copyright 2016 go-dockerclient authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package docker

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

// InitSwarmOptions specify parameters to the InitSwarm function.
// See https://goo.gl/hzkgWu for more details.
type InitSwarmOptions struct {
	swarm.InitRequest
	Context context.Context
}

// InitSwarm initializes a new Swarm and returns the node ID.
// See https://goo.gl/hzkgWu for more details.
func (c *Client) InitSwarm(opts InitSwarmOptions) (string, error) {
	path := "/swarm/init"
	resp, err := c.do("POST", path, doOptions{
		data:      opts.InitRequest,
		forceJSON: true,
		context:   opts.Context,
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var response string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}
	return response, nil
}

// JoinSwarmOptions specify parameters to the JoinSwarm function.
// See https://goo.gl/TdhJWU for more details.
type JoinSwarmOptions struct {
	swarm.JoinRequest
	Context context.Context
}

// JoinSwarm joins an existing Swarm.
// See https://goo.gl/TdhJWU for more details.
func (c *Client) JoinSwarm(opts JoinSwarmOptions) error {
	path := "/swarm/join"
	_, err := c.do("POST", path, doOptions{
		data:      opts.JoinRequest,
		forceJSON: true,
		context:   opts.Context,
	})
	return err
}

// LeaveSwarmOptions specify parameters to the LeaveSwarm function.
// See https://goo.gl/UWDlLg for more details.
type LeaveSwarmOptions struct {
	Force   bool
	Context context.Context
}

// LeaveSwarm leaves a Swarm.
// See https://goo.gl/UWDlLg for more details.
func (c *Client) LeaveSwarm(opts LeaveSwarmOptions) error {
	params := make(url.Values)
	if opts.Force {
		params.Set("force", "1")
	}
	path := "/swarm/leave?" + params.Encode()
	_, err := c.do("POST", path, doOptions{
		context: opts.Context,
	})
	return err
}

// UpdateSwarmOptions specify parameters to the UpdateSwarm function.
// See https://goo.gl/vFbq36 for more details.
type UpdateSwarmOptions struct {
	Version            int
	RotateWorkerToken  bool
	RotateManagerToken bool
	Swarm              swarm.Spec
	Context            context.Context
}

// UpdateSwarm updates a Swarm.
// See https://goo.gl/vFbq36 for more details.
func (c *Client) UpdateSwarm(opts UpdateSwarmOptions) error {
	params := make(url.Values)
	params.Set("version", strconv.Itoa(opts.Version))
	params.Set("rotateWorkerToken", strconv.FormatBool(opts.RotateWorkerToken))
	params.Set("rotateManagerToken", strconv.FormatBool(opts.RotateManagerToken))
	path := "/swarm/update?" + params.Encode()
	_, err := c.do("POST", path, doOptions{
		data:      opts.Swarm,
		forceJSON: true,
		context:   opts.Context,
	})
	return err
}
