package setr

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document05000101 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:setr.050.001.01 Document"`
	Message *SubscriptionBulkOrderConfirmationAmendmentV01 `xml:"SbcptBlkOrdrConfAmdmntV01"`
}

func (d *Document05000101) AddMessage() *SubscriptionBulkOrderConfirmationAmendmentV01 {
	d.Message = new(SubscriptionBulkOrderConfirmationAmendmentV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent sends the SubscriptionBulkOrderConfirmationAmendment message to the instructing party, for example, an investment manager or its authorised representative to amend a previously sent SubscriptionBulkOrderConfirmation message.
// Usage
// The SubscriptionBulkOrderConfirmationAmendment message is used to amend one or more previously sent subscription bulk order confirmations.
// Each bulk order confirmation amendment specified is identified in DealReference. The reference of the original individual order is specified in OrderReference.
// The message identification of the SubscriptionBulkOrder message in which the orders were conveyed may also be quoted in RelatedReference. The message identification of the SubscriptionBulkOrderConfirmation message in which the original order confirmations were conveyed may also be quoted in PreviousReference.
type SubscriptionBulkOrderConfirmationAmendmentV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// General information related to the execution of investment orders.
	BulkExecutionDetails *iso20022.SubscriptionBulkExecution3 `xml:"BlkExctnDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*iso20022.Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderConfirmationAmendmentV01) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderConfirmationAmendmentV01) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderConfirmationAmendmentV01) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationAmendmentV01) AddRelatedReference() *iso20022.AdditionalReference3 {
	s.RelatedReference = new(iso20022.AdditionalReference3)
	return s.RelatedReference
}

func (s *SubscriptionBulkOrderConfirmationAmendmentV01) AddBulkExecutionDetails() *iso20022.SubscriptionBulkExecution3 {
	s.BulkExecutionDetails = new(iso20022.SubscriptionBulkExecution3)
	return s.BulkExecutionDetails
}

func (s *SubscriptionBulkOrderConfirmationAmendmentV01) AddRelatedPartyDetails() *iso20022.Intermediary9 {
	newValue := new(iso20022.Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationAmendmentV01) AddCopyDetails() *iso20022.CopyInformation2 {
	s.CopyDetails = new(iso20022.CopyInformation2)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderConfirmationAmendmentV01) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
func (d *Document05000101) String() (result string, ok bool) { return }
