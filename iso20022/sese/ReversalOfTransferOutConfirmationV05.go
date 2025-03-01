package sese

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00400105 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.004.001.05 Document"`
	Message *ReversalOfTransferOutConfirmationV05 `xml:"RvslOfTrfOutConf"`
}

func (d *Document00400105) AddMessage() *ReversalOfTransferOutConfirmationV05 {
	d.Message = new(ReversalOfTransferOutConfirmationV05)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the ReversalOfTransferOutConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to cancel a previously sent TransferOutConfirmation message.
// Usage
// The ReversalOfTransferOutConfirmation message is used to reverse a previously sent TransferOutConfirmation.
// There are two ways to specify the reversal of the transfer out confirmation. Either:
// - the business references, for example, TransferReference, TransferConfirmationIdentification, of the transfer confirmation are quoted, or,
// - all the details of the transfer confirmation (this includes TransferReference and TransferConfirmationIdentification) are quoted but this is not recommended.
// The message identification of the TransferOutConfirmation message in which the transfer out confirmation was conveyed may also be quoted in PreviousReference. The message identification of the TransferOutInstruction message in which the transfer out instruction was conveyed may also be quoted in RelatedReference.
type ReversalOfTransferOutConfirmationV05 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*iso20022.References15 `xml:"Refs,omitempty"`

	// Choice between reversal by reference or by reversal details.
	Reversal *iso20022.Reversal3Choice `xml:"Rvsl"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *ReversalOfTransferOutConfirmationV05) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *ReversalOfTransferOutConfirmationV05) AddReferences() *iso20022.References15 {
	newValue := new(iso20022.References15)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *ReversalOfTransferOutConfirmationV05) AddReversal() *iso20022.Reversal3Choice {
	r.Reversal = new(iso20022.Reversal3Choice)
	return r.Reversal
}

func (r *ReversalOfTransferOutConfirmationV05) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	r.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return r.MarketPracticeVersion
}

func (r *ReversalOfTransferOutConfirmationV05) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}
func (d *Document00400105) String() (result string, ok bool) { return }
