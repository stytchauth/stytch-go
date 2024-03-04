package mfa

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

// MemberOptions:
type MemberOptions struct {
	// MFAPhoneNumber: The Member's MFA phone number.
	MFAPhoneNumber string `json:"mfa_phone_number,omitempty"`
	// TOTPRegistrationID: The Member's MFA TOTP registration ID.
	TOTPRegistrationID string `json:"totp_registration_id,omitempty"`
}

// MfaRequired:
type MfaRequired struct {
	// MemberOptions: Information about the Member's options for completing MFA.
	MemberOptions *MemberOptions `json:"member_options,omitempty"`
	// SecondaryAuthInitiated: If null, indicates that no secondary authentication has been initiated. If equal
	// to "sms_otp", indicates that the Member has a phone number, and a one time passcode has been sent to the
	// Member's phone number. No secondary authentication will be initiated during calls to the discovery
	// authenticate or list organizations endpoints, even if the Member has a phone number.
	SecondaryAuthInitiated string `json:"secondary_auth_initiated,omitempty"`
}
