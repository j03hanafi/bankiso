package pacs

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00200103 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.03 Document"`
	Message *FIToFIPaymentStatusReportV03 `xml:"FIToFIPmtStsRpt"`
}

func (d *Document00200103) AddMessage() *FIToFIPaymentStatusReportV03 {
	d.Message = new(FIToFIPaymentStatusReportV03)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The FIToFIPaymentStatusReport message is exchanged between agents to provide status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.
// The FIToFIPaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of a credit transfer instruction, a direct debit instruction, as well as other intra-agent instructions (for example FIToFIPaymentCancellationRequest).
// The FIToFIPaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentStatusReport message can be used in domestic and cross-border scenarios.
type FIToFIPaymentStatusReportV03 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *iso20022.GroupHeader37 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *iso20022.OriginalGroupInformation20 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original transactions, to which the status report message refers.
	TransactionInformationAndStatus []*iso20022.PaymentTransactionInformation26 `xml:"TxInfAndSts,omitempty"`
}

func (f *FIToFIPaymentStatusReportV03) AddGroupHeader() *iso20022.GroupHeader37 {
	f.GroupHeader = new(iso20022.GroupHeader37)
	return f.GroupHeader
}

func (f *FIToFIPaymentStatusReportV03) AddOriginalGroupInformationAndStatus() *iso20022.OriginalGroupInformation20 {
	f.OriginalGroupInformationAndStatus = new(iso20022.OriginalGroupInformation20)
	return f.OriginalGroupInformationAndStatus
}

func (f *FIToFIPaymentStatusReportV03) AddTransactionInformationAndStatus() *iso20022.PaymentTransactionInformation26 {
	newValue := new(iso20022.PaymentTransactionInformation26)
	f.TransactionInformationAndStatus = append(f.TransactionInformationAndStatus, newValue)
	return newValue
}
func (d *Document00200103) String() (result string, ok bool) { return }
