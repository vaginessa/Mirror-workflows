// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ACL undocumented
type ACL struct {
	// Object is the base model of ACL
	Object
	// Type undocumented
	Type *ACLType `json:"type,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
	// AccessType undocumented
	AccessType *AccessType `json:"accessType,omitempty"`
	// IdentitySource undocumented
	IdentitySource *string `json:"identitySource,omitempty"`
}