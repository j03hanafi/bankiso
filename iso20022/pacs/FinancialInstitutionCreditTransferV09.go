package pacs

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00900109 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.09 Document" json:"-"`
	Message *FinancialInstitutionCreditTransferV09 `xml:"FICdtTrf" json:"FICdtTrf"`
}

// res.AddMessage()
func (d *Document00900109) AddMessage() *FinancialInstitutionCreditTransferV09 {
	d.Message = new(FinancialInstitutionCreditTransferV09)
	return d.Message
}

type FinancialInstitutionCreditTransferV09 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader93 `xml:"GrpHdr" json:"GrpHdr"`

	// Set of elements providing information specific to the individual credit transfer(s).
	CreditTransferTransactionInformation []*iso20022.CreditTransferTransaction44 `xml:"CdtTrfTxInf" json:"CdtTrfTxInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

// res.AddMessage().AddGroupHeader()
func (f *FinancialInstitutionCreditTransferV09) AddGroupHeader() *iso20022.GroupHeader93 {
	f.GroupHeader = new(iso20022.GroupHeader93)
	return f.GroupHeader
}

// res.AddMessage(). AddCreditTransferTransactionInformation()
func (f *FinancialInstitutionCreditTransferV09) AddCreditTransferTransactionInformation() *iso20022.CreditTransferTransaction44 {
	newValue := new(iso20022.CreditTransferTransaction44)
	f.CreditTransferTransactionInformation = append(f.CreditTransferTransactionInformation, newValue)
	return newValue
}

// res.AddMessage(). AddSupplementaryData()
func (f *FinancialInstitutionCreditTransferV09) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
