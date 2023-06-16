// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/mainflux/mainflux/pkg/messaging"
	sdk "github.com/mainflux/mainflux/pkg/sdk/go"
	"github.com/ultravioletrs/mainflux-ui/ui"
)

const maxNameSize = 1024

type indexReq struct {
}

type loginReq struct {
}

type PasswordResetReq struct {
}

type tokenReq struct {
	Identity string `json:"identity"`
	Secret   string `json:"secret"`
}

type createUserReq struct {
	token string
	user  sdk.User
}

func (req createUserReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.user.Credentials.Secret == "" || req.user.Credentials.Identity == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type createUsersReq struct {
	token     string
	Names     []string
	Emails    []string
	Passwords []string
}

func (req createUsersReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type listUsersReq struct {
	token string
}

func (req listUsersReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type viewResourceReq struct {
	token string
	id    string
}

func (req viewResourceReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	if req.id == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type updateUserReq struct {
	token    string
	id       string
	Name     string                 `json:"name,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (req updateUserReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type updateUserTagsReq struct {
	token string
	id    string
	Tags  []string `json:"tags,omitempty"`
}

func (req updateUserTagsReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type updateUserIdentityReq struct {
	token    string
	id       string
	Identity string `json:"identity,omitempty"`
}

func (req updateUserIdentityReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.Identity == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type updateUserStatusReq struct {
	token  string
	UserID string `json:"userId,omitempty"`
	Status string `json:"status,omitempty"`
}

func (req updateUserStatusReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.Status == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type updateUserPasswordReq struct {
	token   string
	id      string
	OldPass string `json:"oldpass,omitempty"`
	NewPass string `json:"newpass,omitempty"`
}

func (req updateUserPasswordReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.OldPass == "" || req.NewPass == "" {
		return ui.ErrMalformedEntity
	}
	return nil
}

type createThingReq struct {
	token string
	thing sdk.Thing
}

func (req createThingReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type listThingsReq struct {
	token string
}

func (req listThingsReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type updateThingReq struct {
	token    string
	id       string
	Name     string                 `json:"name,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (req updateThingReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.id == "" {
		return ui.ErrMalformedEntity
	}

	if len(req.Name) > maxNameSize {
		return ui.ErrMalformedEntity
	}

	return nil
}

type updateThingTagsReq struct {
	token string
	id    string
	Tags  []string `json:"tags,omitempty"`
}

func (req updateThingTagsReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type updateThingSecretReq struct {
	token  string
	id     string
	Secret string `json:"secret,omitempty"`
}

func (req updateThingSecretReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.Secret == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type updateThingStatusReq struct {
	token   string
	ThingID string `json:"thingID,omitempty"`
	Status  string `json:"status,omitempty"`
}

func (req updateThingStatusReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.Status == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type updateThingOwnerReq struct {
	token string
	id    string
	Owner string `json:"owner,omitempty"`
}

func (req updateThingOwnerReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type createThingsReq struct {
	token     string
	Names     []string                 `json:"names,omitempty"`
	Metadatas []map[string]interface{} `json:"metadata,omitempty"`
}

func (req createThingsReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type createChannelReq struct {
	token   string
	Channel sdk.Channel
}

func (req createChannelReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type createChannelsReq struct {
	token     string
	Names     []string                 `json:"name,omitempty"`
	IDs       []string                 `json:"id,omitempty"`
	Metadatas []map[string]interface{} `json:"metadata,omitempty"`
}

func (req createChannelsReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type updateChannelReq struct {
	token       string
	id          string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

func (req updateChannelReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.id == "" {
		return ui.ErrMalformedEntity
	}

	if len(req.Name) > maxNameSize {
		return ui.ErrMalformedEntity
	}

	return nil
}

type listChannelsReq struct {
	token string
}

func (req listChannelsReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type connectThingReq struct {
	token   string
	ConnIDs sdk.ConnectionIDs
}

func (req connectThingReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type connectChannelReq struct {
	token   string
	ConnIDs sdk.ConnectionIDs
}

func (req connectChannelReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type connectReq struct {
	token   string
	ChanID  []string `json:"chan_id,omitempty"`
	ThingID []string `json:"thing_id,omitempty"`
}

func (req connectReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type disconnectThingReq struct {
	token   string
	ChanID  string `json:"chan_id,omitempty"`
	ThingID string `json:"thing_id,omitempty"`
}

func (req disconnectThingReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	if req.ChanID == "" || req.ThingID == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type disconnectChannelReq struct {
	token   string
	ChanID  string `json:"chan_id,omitempty"`
	ThingID string `json:"thing_id,omitempty"`
}

func (req disconnectChannelReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	if req.ChanID == "" || req.ThingID == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type disconnectReq struct {
	token   string
	ChanID  []string `json:"chan_id,omitempty"`
	ThingID []string `json:"thing_id,omitempty"`
}

func (req disconnectReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type updateChannelStatusReq struct {
	token     string
	ChannelID string `json:"channelID,omitempty"`
	Status    string `json:"status,omitempty"`
}

func (req updateChannelStatusReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.Status == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type createGroupReq struct {
	token string
	Group sdk.Group
}

func (req createGroupReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	return nil
}

type listGroupsReq struct {
	token string
}

func (req listGroupsReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	return nil
}

type updateGroupReq struct {
	token       string
	id          string
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	ParentID    string                 `json:"parent_id,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

func (req updateGroupReq) validate() error {
	if req.token == "" {
		return ui.ErrMalformedEntity
	}

	if len(req.Name) > maxNameSize {
		return ui.ErrMalformedEntity
	}

	return nil
}

type assignReq struct {
	token    string
	groupID  string
	Type     []string `json:"Type,omitempty"`
	MemberID string   `json:"member_id,omitempty"`
}

func (req assignReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	if req.groupID == "" || req.MemberID == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type unassignReq struct {
	token    string
	groupID  string
	Type     string `json:"type,omitempty"`
	MemberID string `json:"member_id"`
}

func (req unassignReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}

	if req.groupID == "" || req.MemberID == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type updateGroupStatusReq struct {
	token   string
	GroupID string `json:"groupId,omitempty"`
	Status  string `json:"status,omitempty"`
}

func (req updateGroupStatusReq) validate() error {
	if req.token == "" {
		return ui.ErrUnauthorizedAccess
	}
	if req.Status == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type listPoliciesReq struct {
	token string
}

func (req listPoliciesReq) validate() error {
	if req.token == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type addPolicyReq struct {
	token  string
	Policy sdk.Policy `json:"policy,omitempty"`
}

func (req addPolicyReq) validate() error {
	if req.token == "" {
		return ui.ErrMalformedEntity
	}
	return nil
}

type updatePolicyReq struct {
	token  string
	Policy sdk.Policy `json:"policy,omitempty"`
}

func (req updatePolicyReq) validate() error {
	if req.token == "" {
		return ui.ErrMalformedEntity
	}
	return nil
}

type deletePolicyReq struct {
	token  string
	Policy sdk.Policy `json:"policy,omitempty"`
}

func (req deletePolicyReq) validate() error {
	if req.token == "" {
		return ui.ErrMalformedEntity
	}
	return nil
}

type publishReq struct {
	msg      *messaging.Message
	thingKey string
	token    string
}

func (req publishReq) validate() error {
	if req.token == "" {
		return ui.ErrMalformedEntity
	}

	if req.thingKey == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}

type readMessageReq struct {
}

type wsConnectionReq struct {
	ChanID   string `json:"chan_id,omitempty"`
	ThingKey string `json:"thing_key,omitempty"`
}

func (req wsConnectionReq) validate() error {
	if req.ChanID == "" {
		return ui.ErrMalformedEntity
	}
	if req.ThingKey == "" {
		return ui.ErrMalformedEntity
	}

	return nil
}
