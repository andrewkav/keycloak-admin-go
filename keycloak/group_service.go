package keycloak

import (
	"context"
	"net/url"
	"strings"
)

// GroupService manages realm groups
type GroupService service

// NewGroupService creates a new group service
func NewGroupService(c *Client) *GroupService {
	return &GroupService{
		client: c,
	}
}

// Find returns groups based on query params
// Params:
// - first
// - max
// - search
func (gs *GroupService) GetGroups(ctx context.Context, realm string, params map[string]string) (groups []GroupRepresentation, err error) {
	path := "/realms/{realm}/groups"

	_, err = gs.client.newRequest(ctx).
		SetQueryParams(params).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&groups).
		Get(path)

	return
}

// Create will update the group and set the parent if it exists. Create it and set the parent if the group doesn’t exist.
func (gs *GroupService) Create(ctx context.Context, realm string, g GroupRepresentation) (ID string, err error) {
	path := "/realms/{realm}/groups"

	response, err := gs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(g).
		Post(path)

	if err != nil {
		return
	}

	location, err := url.Parse(response.Header().Get("Location"))
	if err != nil {
		return
	}

	components := strings.Split(location.Path, "/")
	ID = components[len(components)-1]

	return
}

// Count gets the number of groups
func (gs *GroupService) Count(ctx context.Context, realm string) (count uint32, err error) {
	path := "/realms/{realm}/groups/count"

	_, err = gs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&count).
		Get(path)

	return
}

// Get gets the group by group ID
func (gs *GroupService) Get(ctx context.Context, realm, ID string) (*GroupRepresentation, error) {
	path := "/realms/{realm}/groups/{id}"

	var g GroupRepresentation
	_, err := gs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    ID,
		}).
		SetResult(&g).
		Get(path)

	if err != nil {
		return nil, err
	}

	return &g, nil
}

// Update updates the group
func (gs *GroupService) Update(ctx context.Context, realm string, g GroupRepresentation) error {
	// nolint: goconst
	path := "/realms/{realm}/groups/{id}"

	_, err := gs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    g.ID,
		}).
		SetBody(g).
		Put(path)

	return err
}

// Delete deletes the group
func (gs *GroupService) Delete(ctx context.Context, realm, ID string) error {
	// nolint: goconst
	path := "/realms/{realm}/groups/{id}"

	_, err := gs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    ID,
		}).
		Delete(path)

	return err
}

// AddChild adds a child group to the group
func (gs *GroupService) AddChild(ctx context.Context, realm, ID string, g GroupRepresentation) (childID string, err error) {
	path := "/realms/{realm}/groups/{id}/children"

	response, err := gs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    ID,
		}).
		SetBody(g).
		Post(path)

	if err != nil {
		return
	}

	location, err := url.Parse(response.Header().Get("Location"))
	if err != nil {
		return
	}

	components := strings.Split(location.Path, "/")
	childID = components[len(components)-1]

	return
}
