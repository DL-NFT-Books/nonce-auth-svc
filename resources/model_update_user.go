/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateUser struct {
	Key
	Attributes UpdateUserAttributes `json:"attributes"`
}
type UpdateUserRequest struct {
	Data     UpdateUser `json:"data"`
	Included Included   `json:"included"`
}

type UpdateUserListRequest struct {
	Data     []UpdateUser `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustUpdateUser - returns UpdateUser from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateUser(key Key) *UpdateUser {
	var updateUser UpdateUser
	if c.tryFindEntry(key, &updateUser) {
		return &updateUser
	}
	return nil
}
