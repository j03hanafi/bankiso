package pacs

import (
	"encoding/xml"
	"github.com/figassis/bankiso/iso20022"
)

type Document00200110 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Document" json:"Document"'`
	Message *FIToFIPaymentStatusReportV10 `xml:"FIToFIPmtStsRpt" json:"FIToFIPmtStsRpt"`
}

func (d *Document00200110) AddMessage() *FIToFIPaymentStatusReportV10 {
	d.Message = new(FIToFIPaymentStatusReportV10)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The FIToFIPaymentStatusReport message is exchanged between agents to provide status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.
// The FIToFIPaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of a credit transfer instruction, a direct debit instruction, as well as other intra-agent instructions (for example FIToFIPaymentCancellationRequest).
// The FIToFIPaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentStatusReport message can be used in domestic and cross-border scenarios.
type FIToFIPaymentStatusReportV10 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *iso20022.GroupHeader91 `xml:"GrpHdr" json:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus []*iso20022.OriginalGroupHeader17 `xml:"OrgnlGrpInfAndSts,omitempty" json:"OrgnlGrpInfAndSts,omitempty"`

	// Information concerning the original transactions, to which the status report message refers.
	TransactionInformationAndStatus []*iso20022.PaymentTransaction110 `xml:"TxInfAndSts,omitempty" json:"TxInfAndSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

func (f *FIToFIPaymentStatusReportV10) AddGroupHeader() *iso20022.GroupHeader91 {
	f.GroupHeader = new(iso20022.GroupHeader91)
	return f.GroupHeader
}

func (f *FIToFIPaymentStatusReportV10) AddOriginalGroupInformationAndStatus() *iso20022.OriginalGroupHeader17 {
	newValue := new(iso20022.OriginalGroupHeader17)
	f.OriginalGroupInformationAndStatus = append(f.OriginalGroupInformationAndStatus, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV10) AddTransactionInformationAndStatus() *iso20022.PaymentTransaction110 {
	newValue := new(iso20022.PaymentTransaction110)
	f.TransactionInformationAndStatus = append(f.TransactionInformationAndStatus, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV10) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
func (d *Document00200110) String() (result string, ok bool) { return }
