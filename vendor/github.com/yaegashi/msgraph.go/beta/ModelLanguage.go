// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// LanguageProficiency undocumented
type LanguageProficiency struct {
	// ItemFacet is the base model of LanguageProficiency
	ItemFacet
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Tag undocumented
	Tag *string `json:"tag,omitempty"`
	// Proficiency undocumented
	Proficiency *LanguageProficiencyLevel `json:"proficiency,omitempty"`
}