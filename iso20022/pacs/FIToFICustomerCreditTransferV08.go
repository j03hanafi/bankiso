package pacs

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

// res := new(pacs.Document00800108)
type Document00800108 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Document" json:"-"`
	Message *FIToFICustomerCreditTransferV08 `xml:"FIToFICstmrCdtTrf" json:"FIToFICstmrCdtTrf"`
}

// res.AddMessage()
func (d *Document00800108) AddMessage() *FIToFICustomerCreditTransferV08 {
	d.Message = new(FIToFICustomerCreditTransferV08)
	return d.Message
}

type FIToFICustomerCreditTransferV08 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader93 `xml:"GrpHdr" json:"GrpHdr"`

	// Set of elements providing information specific to the individual credit transfer(s).
	CreditTransferTransactionInformation []*iso20022.CreditTransferTransaction39 `xml:"CdtTrfTxInf" json:"CdtTrfTxInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

// res.AddMessage().AddGroupHeader()
func (f *FIToFICustomerCreditTransferV08) AddGroupHeader() *iso20022.GroupHeader93 {
	f.GroupHeader = new(iso20022.GroupHeader93)
	return f.GroupHeader
}

// res.AddMessage(). AddCreditTransferTransactionInformation()
func (f *FIToFICustomerCreditTransferV08) AddCreditTransferTransactionInformation() *iso20022.CreditTransferTransaction39 {
	newValue := new(iso20022.CreditTransferTransaction39)
	f.CreditTransferTransactionInformation = append(f.CreditTransferTransactionInformation, newValue)
	return newValue
}

// res.AddMessage(). AddCreditTransferTransactionInformation()
func (f *FIToFICustomerCreditTransferV08) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
