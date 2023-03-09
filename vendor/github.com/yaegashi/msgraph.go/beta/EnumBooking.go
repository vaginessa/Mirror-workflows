// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BookingInvoiceStatus undocumented
type BookingInvoiceStatus string

const (
	// BookingInvoiceStatusVDraft undocumented
	BookingInvoiceStatusVDraft BookingInvoiceStatus = "draft"
	// BookingInvoiceStatusVReviewing undocumented
	BookingInvoiceStatusVReviewing BookingInvoiceStatus = "reviewing"
	// BookingInvoiceStatusVOpen undocumented
	BookingInvoiceStatusVOpen BookingInvoiceStatus = "open"
	// BookingInvoiceStatusVCanceled undocumented
	BookingInvoiceStatusVCanceled BookingInvoiceStatus = "canceled"
	// BookingInvoiceStatusVPaid undocumented
	BookingInvoiceStatusVPaid BookingInvoiceStatus = "paid"
	// BookingInvoiceStatusVCorrective undocumented
	BookingInvoiceStatusVCorrective BookingInvoiceStatus = "corrective"
)

var (
	// BookingInvoiceStatusPDraft is a pointer to BookingInvoiceStatusVDraft
	BookingInvoiceStatusPDraft = &_BookingInvoiceStatusPDraft
	// BookingInvoiceStatusPReviewing is a pointer to BookingInvoiceStatusVReviewing
	BookingInvoiceStatusPReviewing = &_BookingInvoiceStatusPReviewing
	// BookingInvoiceStatusPOpen is a pointer to BookingInvoiceStatusVOpen
	BookingInvoiceStatusPOpen = &_BookingInvoiceStatusPOpen
	// BookingInvoiceStatusPCanceled is a pointer to BookingInvoiceStatusVCanceled
	BookingInvoiceStatusPCanceled = &_BookingInvoiceStatusPCanceled
	// BookingInvoiceStatusPPaid is a pointer to BookingInvoiceStatusVPaid
	BookingInvoiceStatusPPaid = &_BookingInvoiceStatusPPaid
	// BookingInvoiceStatusPCorrective is a pointer to BookingInvoiceStatusVCorrective
	BookingInvoiceStatusPCorrective = &_BookingInvoiceStatusPCorrective
)

var (
	_BookingInvoiceStatusPDraft      = BookingInvoiceStatusVDraft
	_BookingInvoiceStatusPReviewing  = BookingInvoiceStatusVReviewing
	_BookingInvoiceStatusPOpen       = BookingInvoiceStatusVOpen
	_BookingInvoiceStatusPCanceled   = BookingInvoiceStatusVCanceled
	_BookingInvoiceStatusPPaid       = BookingInvoiceStatusVPaid
	_BookingInvoiceStatusPCorrective = BookingInvoiceStatusVCorrective
)

// BookingPriceType undocumented
type BookingPriceType string

const (
	// BookingPriceTypeVUndefined undocumented
	BookingPriceTypeVUndefined BookingPriceType = "undefined"
	// BookingPriceTypeVFixedPrice undocumented
	BookingPriceTypeVFixedPrice BookingPriceType = "fixedPrice"
	// BookingPriceTypeVStartingAt undocumented
	BookingPriceTypeVStartingAt BookingPriceType = "startingAt"
	// BookingPriceTypeVHourly undocumented
	BookingPriceTypeVHourly BookingPriceType = "hourly"
	// BookingPriceTypeVFree undocumented
	BookingPriceTypeVFree BookingPriceType = "free"
	// BookingPriceTypeVPriceVaries undocumented
	BookingPriceTypeVPriceVaries BookingPriceType = "priceVaries"
	// BookingPriceTypeVCallUs undocumented
	BookingPriceTypeVCallUs BookingPriceType = "callUs"
	// BookingPriceTypeVNotSet undocumented
	BookingPriceTypeVNotSet BookingPriceType = "notSet"
)

var (
	// BookingPriceTypePUndefined is a pointer to BookingPriceTypeVUndefined
	BookingPriceTypePUndefined = &_BookingPriceTypePUndefined
	// BookingPriceTypePFixedPrice is a pointer to BookingPriceTypeVFixedPrice
	BookingPriceTypePFixedPrice = &_BookingPriceTypePFixedPrice
	// BookingPriceTypePStartingAt is a pointer to BookingPriceTypeVStartingAt
	BookingPriceTypePStartingAt = &_BookingPriceTypePStartingAt
	// BookingPriceTypePHourly is a pointer to BookingPriceTypeVHourly
	BookingPriceTypePHourly = &_BookingPriceTypePHourly
	// BookingPriceTypePFree is a pointer to BookingPriceTypeVFree
	BookingPriceTypePFree = &_BookingPriceTypePFree
	// BookingPriceTypePPriceVaries is a pointer to BookingPriceTypeVPriceVaries
	BookingPriceTypePPriceVaries = &_BookingPriceTypePPriceVaries
	// BookingPriceTypePCallUs is a pointer to BookingPriceTypeVCallUs
	BookingPriceTypePCallUs = &_BookingPriceTypePCallUs
	// BookingPriceTypePNotSet is a pointer to BookingPriceTypeVNotSet
	BookingPriceTypePNotSet = &_BookingPriceTypePNotSet
)

var (
	_BookingPriceTypePUndefined   = BookingPriceTypeVUndefined
	_BookingPriceTypePFixedPrice  = BookingPriceTypeVFixedPrice
	_BookingPriceTypePStartingAt  = BookingPriceTypeVStartingAt
	_BookingPriceTypePHourly      = BookingPriceTypeVHourly
	_BookingPriceTypePFree        = BookingPriceTypeVFree
	_BookingPriceTypePPriceVaries = BookingPriceTypeVPriceVaries
	_BookingPriceTypePCallUs      = BookingPriceTypeVCallUs
	_BookingPriceTypePNotSet      = BookingPriceTypeVNotSet
)

// BookingReminderRecipients undocumented
type BookingReminderRecipients string

const (
	// BookingReminderRecipientsVAllAttendees undocumented
	BookingReminderRecipientsVAllAttendees BookingReminderRecipients = "allAttendees"
	// BookingReminderRecipientsVStaff undocumented
	BookingReminderRecipientsVStaff BookingReminderRecipients = "staff"
	// BookingReminderRecipientsVCustomer undocumented
	BookingReminderRecipientsVCustomer BookingReminderRecipients = "customer"
)

var (
	// BookingReminderRecipientsPAllAttendees is a pointer to BookingReminderRecipientsVAllAttendees
	BookingReminderRecipientsPAllAttendees = &_BookingReminderRecipientsPAllAttendees
	// BookingReminderRecipientsPStaff is a pointer to BookingReminderRecipientsVStaff
	BookingReminderRecipientsPStaff = &_BookingReminderRecipientsPStaff
	// BookingReminderRecipientsPCustomer is a pointer to BookingReminderRecipientsVCustomer
	BookingReminderRecipientsPCustomer = &_BookingReminderRecipientsPCustomer
)

var (
	_BookingReminderRecipientsPAllAttendees = BookingReminderRecipientsVAllAttendees
	_BookingReminderRecipientsPStaff        = BookingReminderRecipientsVStaff
	_BookingReminderRecipientsPCustomer     = BookingReminderRecipientsVCustomer
)

// BookingStaffRole undocumented
type BookingStaffRole string

const (
	// BookingStaffRoleVGuest undocumented
	BookingStaffRoleVGuest BookingStaffRole = "guest"
	// BookingStaffRoleVAdministrator undocumented
	BookingStaffRoleVAdministrator BookingStaffRole = "administrator"
	// BookingStaffRoleVViewer undocumented
	BookingStaffRoleVViewer BookingStaffRole = "viewer"
	// BookingStaffRoleVExternalGuest undocumented
	BookingStaffRoleVExternalGuest BookingStaffRole = "externalGuest"
)

var (
	// BookingStaffRolePGuest is a pointer to BookingStaffRoleVGuest
	BookingStaffRolePGuest = &_BookingStaffRolePGuest
	// BookingStaffRolePAdministrator is a pointer to BookingStaffRoleVAdministrator
	BookingStaffRolePAdministrator = &_BookingStaffRolePAdministrator
	// BookingStaffRolePViewer is a pointer to BookingStaffRoleVViewer
	BookingStaffRolePViewer = &_BookingStaffRolePViewer
	// BookingStaffRolePExternalGuest is a pointer to BookingStaffRoleVExternalGuest
	BookingStaffRolePExternalGuest = &_BookingStaffRolePExternalGuest
)

var (
	_BookingStaffRolePGuest         = BookingStaffRoleVGuest
	_BookingStaffRolePAdministrator = BookingStaffRoleVAdministrator
	_BookingStaffRolePViewer        = BookingStaffRoleVViewer
	_BookingStaffRolePExternalGuest = BookingStaffRoleVExternalGuest
)

// BookingType undocumented
type BookingType string

const (
	// BookingTypeVUnknown undocumented
	BookingTypeVUnknown BookingType = "unknown"
	// BookingTypeVStandard undocumented
	BookingTypeVStandard BookingType = "standard"
	// BookingTypeVReserved undocumented
	BookingTypeVReserved BookingType = "reserved"
)

var (
	// BookingTypePUnknown is a pointer to BookingTypeVUnknown
	BookingTypePUnknown = &_BookingTypePUnknown
	// BookingTypePStandard is a pointer to BookingTypeVStandard
	BookingTypePStandard = &_BookingTypePStandard
	// BookingTypePReserved is a pointer to BookingTypeVReserved
	BookingTypePReserved = &_BookingTypePReserved
)

var (
	_BookingTypePUnknown  = BookingTypeVUnknown
	_BookingTypePStandard = BookingTypeVStandard
	_BookingTypePReserved = BookingTypeVReserved
)