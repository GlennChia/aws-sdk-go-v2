// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An entitlement represents capacity in a product owned by the customer. For
// example, a customer might own some number of users or seats in an SaaS
// application or some amount of data capacity in a multi-tenant database.
type Entitlement struct {

	// The customer identifier is a handle to each unique customer in an application.
	// Customer identifiers are obtained through the ResolveCustomer operation in AWS
	// Marketplace Metering Service.
	CustomerIdentifier *string

	// The dimension for which the given entitlement applies. Dimensions represent
	// categories of capacity in a product and are specified when the product is listed
	// in AWS Marketplace.
	Dimension *string

	// The expiration date represents the minimum date through which this entitlement
	// is expected to remain valid. For contractual products listed on AWS Marketplace,
	// the expiration date is the date at which the customer will renew or cancel their
	// contract. Customers who are opting to renew their contract will still have
	// entitlements with an expiration date.
	ExpirationDate *time.Time

	// The product code for which the given entitlement applies. Product codes are
	// provided by AWS Marketplace when the product listing is created.
	ProductCode *string

	// The EntitlementValue represents the amount of capacity that the customer is
	// entitled to for the product.
	Value EntitlementValue

	noSmithyDocumentSerde
}

// The EntitlementValue represents the amount of capacity that the customer is
// entitled to for the product.
//
// The following types satisfy this interface:
//  EntitlementValueMemberBooleanValue
//  EntitlementValueMemberDoubleValue
//  EntitlementValueMemberIntegerValue
//  EntitlementValueMemberStringValue
type EntitlementValue interface {
	isEntitlementValue()
}

// The BooleanValue field will be populated with a boolean value when the
// entitlement is a boolean type. Otherwise, the field will not be set.
type EntitlementValueMemberBooleanValue struct {
	Value bool

	noSmithyDocumentSerde
}

func (*EntitlementValueMemberBooleanValue) isEntitlementValue() {}

// The DoubleValue field will be populated with a double value when the entitlement
// is a double type. Otherwise, the field will not be set.
type EntitlementValueMemberDoubleValue struct {
	Value float64

	noSmithyDocumentSerde
}

func (*EntitlementValueMemberDoubleValue) isEntitlementValue() {}

// The IntegerValue field will be populated with an integer value when the
// entitlement is an integer type. Otherwise, the field will not be set.
type EntitlementValueMemberIntegerValue struct {
	Value int32

	noSmithyDocumentSerde
}

func (*EntitlementValueMemberIntegerValue) isEntitlementValue() {}

// The StringValue field will be populated with a string value when the entitlement
// is a string type. Otherwise, the field will not be set.
type EntitlementValueMemberStringValue struct {
	Value string

	noSmithyDocumentSerde
}

func (*EntitlementValueMemberStringValue) isEntitlementValue() {}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isEntitlementValue() {}
