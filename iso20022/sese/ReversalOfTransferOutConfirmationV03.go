package sese

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00400103 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.004.001.03 Document"`
	Message *ReversalOfTransferOutConfirmationV03 `xml:"RvslOfTrfOutConf"`
}

func (d *Document00400103) AddMessage() *ReversalOfTransferOutConfirmationV03 {
	d.Message = new(ReversalOfTransferOutConfirmationV03)
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
type ReversalOfTransferOutConfirmationV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*iso20022.References11 `xml:"Refs"`

	// Reference of the transfer out confirmation to be reversed.
	ReversalByReference *iso20022.TransferReference2 `xml:"RvslByRef,omitempty"`

	// Copy of the transfer out confirmation to reverse.
	ReversalByTransferOutConfirmationDetails *iso20022.TransferOut8 `xml:"RvslByTrfOutConfDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *ReversalOfTransferOutConfirmationV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *ReversalOfTransferOutConfirmationV03) AddReferences() *iso20022.References11 {
	newValue := new(iso20022.References11)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *ReversalOfTransferOutConfirmationV03) AddReversalByReference() *iso20022.TransferReference2 {
	r.ReversalByReference = new(iso20022.TransferReference2)
	return r.ReversalByReference
}

func (r *ReversalOfTransferOutConfirmationV03) AddReversalByTransferOutConfirmationDetails() *iso20022.TransferOut8 {
	r.ReversalByTransferOutConfirmationDetails = new(iso20022.TransferOut8)
	return r.ReversalByTransferOutConfirmationDetails
}

func (r *ReversalOfTransferOutConfirmationV03) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}
func (d *Document00400103) String() (result string, ok bool) { return }
