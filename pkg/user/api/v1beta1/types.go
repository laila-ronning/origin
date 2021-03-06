package v1beta1

import kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta3"

// Auth system gets identity name and provider
// POST to UserIdentityMapping, get back error or a filled out UserIdentityMapping object

type User struct {
	kapi.TypeMeta   `json:",inline"`
	kapi.ObjectMeta `json:"metadata,omitempty"`

	FullName string `json:"fullName,omitempty"`

	Identities []string `json:"identities"`
}

type UserList struct {
	kapi.TypeMeta `json:",inline"`
	kapi.ListMeta `json:"metadata,omitempty"`
	Items         []User `json:"items"`
}

type Identity struct {
	kapi.TypeMeta   `json:",inline"`
	kapi.ObjectMeta `json:"metadata,omitempty"`

	// ProviderName is the source of identity information
	ProviderName string `json:"providerName"`

	// ProviderUserName uniquely represents this identity in the scope of the provider
	ProviderUserName string `json:"providerUserName"`

	// User is a reference to the user this identity is associated with
	// Both Name and UID must be set
	User kapi.ObjectReference `json:"user"`

	Extra map[string]string `json:"extra,omitempty"`
}

type IdentityList struct {
	kapi.TypeMeta `json:",inline"`
	kapi.ListMeta `json:"metadata,omitempty"`
	Items         []Identity `json:"items"`
}

type UserIdentityMapping struct {
	kapi.TypeMeta   `json:",inline"`
	kapi.ObjectMeta `json:"metadata,omitempty"`

	Identity kapi.ObjectReference `json:"identity,omitempty"`
	User     kapi.ObjectReference `json:"user,omitempty"`
}

func (*User) IsAnAPIObject()                {}
func (*UserList) IsAnAPIObject()            {}
func (*Identity) IsAnAPIObject()            {}
func (*IdentityList) IsAnAPIObject()        {}
func (*UserIdentityMapping) IsAnAPIObject() {}
