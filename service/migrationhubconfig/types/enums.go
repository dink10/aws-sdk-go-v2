// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type TargetType string

// Enum values for TargetType
const (
	TargetTypeAccount TargetType = "ACCOUNT"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"ACCOUNT",
	}
}
