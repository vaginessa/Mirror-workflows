// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MailboxSettings undocumented
type MailboxSettings struct {
	// Object is the base model of MailboxSettings
	Object
	// AutomaticRepliesSetting undocumented
	AutomaticRepliesSetting *AutomaticRepliesSetting `json:"automaticRepliesSetting,omitempty"`
	// ArchiveFolder undocumented
	ArchiveFolder *string `json:"archiveFolder,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
	// Language undocumented
	Language *LocaleInfo `json:"language,omitempty"`
	// DelegateMeetingMessageDeliveryOptions undocumented
	DelegateMeetingMessageDeliveryOptions *DelegateMeetingMessageDeliveryOptions `json:"delegateMeetingMessageDeliveryOptions,omitempty"`
	// WorkingHours undocumented
	WorkingHours *WorkingHours `json:"workingHours,omitempty"`
	// DateFormat undocumented
	DateFormat *string `json:"dateFormat,omitempty"`
	// TimeFormat undocumented
	TimeFormat *string `json:"timeFormat,omitempty"`
}

// MailboxUsageDetail undocumented
type MailboxUsageDetail struct {
	// Entity is the base model of MailboxUsageDetail
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// DeletedDate undocumented
	DeletedDate *Date `json:"deletedDate,omitempty"`
	// CreatedDate undocumented
	CreatedDate *Date `json:"createdDate,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// ItemCount undocumented
	ItemCount *int `json:"itemCount,omitempty"`
	// StorageUsedInBytes undocumented
	StorageUsedInBytes *int `json:"storageUsedInBytes,omitempty"`
	// DeletedItemCount undocumented
	DeletedItemCount *int `json:"deletedItemCount,omitempty"`
	// DeletedItemSizeInBytes undocumented
	DeletedItemSizeInBytes *int `json:"deletedItemSizeInBytes,omitempty"`
	// IssueWarningQuotaInBytes undocumented
	IssueWarningQuotaInBytes *int `json:"issueWarningQuotaInBytes,omitempty"`
	// ProhibitSendQuotaInBytes undocumented
	ProhibitSendQuotaInBytes *int `json:"prohibitSendQuotaInBytes,omitempty"`
	// ProhibitSendReceiveQuotaInBytes undocumented
	ProhibitSendReceiveQuotaInBytes *int `json:"prohibitSendReceiveQuotaInBytes,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// MailboxUsageMailboxCounts undocumented
type MailboxUsageMailboxCounts struct {
	// Entity is the base model of MailboxUsageMailboxCounts
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
	// Active undocumented
	Active *int `json:"active,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// MailboxUsageQuotaStatusMailboxCounts undocumented
type MailboxUsageQuotaStatusMailboxCounts struct {
	// Entity is the base model of MailboxUsageQuotaStatusMailboxCounts
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// UnderLimit undocumented
	UnderLimit *int `json:"underLimit,omitempty"`
	// WarningIssued undocumented
	WarningIssued *int `json:"warningIssued,omitempty"`
	// SendProhibited undocumented
	SendProhibited *int `json:"sendProhibited,omitempty"`
	// SendReceiveProhibited undocumented
	SendReceiveProhibited *int `json:"sendReceiveProhibited,omitempty"`
	// Indeterminate undocumented
	Indeterminate *int `json:"indeterminate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

// MailboxUsageStorage undocumented
type MailboxUsageStorage struct {
	// Entity is the base model of MailboxUsageStorage
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// StorageUsedInBytes undocumented
	StorageUsedInBytes *int `json:"storageUsedInBytes,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}