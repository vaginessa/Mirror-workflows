// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AuthenticationDetail undocumented
type AuthenticationDetail struct {
	// Object is the base model of AuthenticationDetail
	Object
	// AuthenticationStepDateTime undocumented
	AuthenticationStepDateTime *time.Time `json:"authenticationStepDateTime,omitempty"`
	// AuthenticationMethod undocumented
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
	// AuthenticationMethodDetail undocumented
	AuthenticationMethodDetail *string `json:"authenticationMethodDetail,omitempty"`
	// Succeeded undocumented
	Succeeded *bool `json:"succeeded,omitempty"`
	// AuthenticationStepResultDetail undocumented
	AuthenticationStepResultDetail *string `json:"authenticationStepResultDetail,omitempty"`
	// AuthenticationStepRequirement undocumented
	AuthenticationStepRequirement *string `json:"authenticationStepRequirement,omitempty"`
}

// AuthenticationRequirementPolicy undocumented
type AuthenticationRequirementPolicy struct {
	// Object is the base model of AuthenticationRequirementPolicy
	Object
	// RequirementProvider undocumented
	RequirementProvider *RequirementProvider `json:"requirementProvider,omitempty"`
	// Detail undocumented
	Detail *string `json:"detail,omitempty"`
}