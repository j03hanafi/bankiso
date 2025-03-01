package sese

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.003.001.03 Document"`
	Message *TransferOutConfirmationV03 `xml:"TrfOutConf"`
}

func (d *Document00300103) AddMessage() *TransferOutConfirmationV03 {
	d.Message = new(TransferOutConfirmationV03)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferOutConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to confirm the delivery of a financial instrument, free of payment, on a given date, to a specified party.
// This message may also be used to confirm the delivery of a financial instrument, free of payment, to another of the instructing parties own accounts or to a third party.
// Usage
// The TransferOutConfirmation message is used to confirm the withdrawal of a financial instrument from the owner's account and its delivery to another own account, or to a third party, has taken place.
// The reference of the transfer confirmation is identified in TransferConfirmationReference. The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferOutInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferOutConfirmationV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *iso20022.Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*iso20022.Transfer13 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *iso20022.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *iso20022.ReceiveInformation6 `xml:"SttlmDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOutConfirmationV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferOutConfirmationV03) AddPoolReference() *iso20022.AdditionalReference2 {
	t.PoolReference = new(iso20022.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferOutConfirmationV03) AddPreviousReference() *iso20022.AdditionalReference2 {
	t.PreviousReference = new(iso20022.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferOutConfirmationV03) AddRelatedReference() *iso20022.AdditionalReference2 {
	t.RelatedReference = new(iso20022.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferOutConfirmationV03) SetMasterReference(value string) {
	t.MasterReference = (*iso20022.Max35Text)(&value)
}

func (t *TransferOutConfirmationV03) AddTransferDetails() *iso20022.Transfer13 {
	newValue := new(iso20022.Transfer13)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOutConfirmationV03) AddAccountDetails() *iso20022.InvestmentAccount22 {
	t.AccountDetails = new(iso20022.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOutConfirmationV03) AddSettlementDetails() *iso20022.ReceiveInformation6 {
	t.SettlementDetails = new(iso20022.ReceiveInformation6)
	return t.SettlementDetails
}

func (t *TransferOutConfirmationV03) AddCopyDetails() *iso20022.CopyInformation2 {
	t.CopyDetails = new(iso20022.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferOutConfirmationV03) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
func (d *Document00300103) String() (result string, ok bool) { return }
