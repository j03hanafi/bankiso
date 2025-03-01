package sese

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00500104 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.04 Document"`
	Message *TransferInInstructionV04 `xml:"TrfInInstr"`
}

func (d *Document00500104) AddMessage() *TransferInInstructionV04 {
	d.Message = new(TransferInInstructionV04)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the TransferInInstruction message to the executing party, for example, a transfer agent, to instruct the receipt of a financial instrument, free of payment, on a given date from a specified party.
// This message may also be used to instruct the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// Usage
// The TransferInInstruction message is used to instruct the receipt of a financial instrument from another account, either owned by the instructing party or by a third party.
type TransferInInstructionV04 struct {

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
	TransferDetails []*iso20022.Transfer21 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *iso20022.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *iso20022.DeliverInformation9 `xml:"SttlmDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInInstructionV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInInstructionV04) AddPoolReference() *iso20022.AdditionalReference2 {
	t.PoolReference = new(iso20022.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInInstructionV04) AddPreviousReference() *iso20022.AdditionalReference2 {
	t.PreviousReference = new(iso20022.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInInstructionV04) AddRelatedReference() *iso20022.AdditionalReference2 {
	t.RelatedReference = new(iso20022.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInInstructionV04) SetMasterReference(value string) {
	t.MasterReference = (*iso20022.Max35Text)(&value)
}

func (t *TransferInInstructionV04) AddTransferDetails() *iso20022.Transfer21 {
	newValue := new(iso20022.Transfer21)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferInInstructionV04) AddAccountDetails() *iso20022.InvestmentAccount22 {
	t.AccountDetails = new(iso20022.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferInInstructionV04) AddSettlementDetails() *iso20022.DeliverInformation9 {
	t.SettlementDetails = new(iso20022.DeliverInformation9)
	return t.SettlementDetails
}

func (t *TransferInInstructionV04) AddCopyDetails() *iso20022.CopyInformation2 {
	t.CopyDetails = new(iso20022.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferInInstructionV04) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
func (d *Document00500104) String() (result string, ok bool) { return }
