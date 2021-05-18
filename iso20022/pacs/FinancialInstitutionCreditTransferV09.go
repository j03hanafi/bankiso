package pacs

import (
	"encoding/xml"

	"github.com/figassis/bankiso/iso20022"
)

type Document00900109 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.09 Document" json:"Document"`
	Message *FinancialInstitutionCreditTransferV09 `xml:"FICdtTrf" json:"FICdtTrf"`
}

type FinancialInstitutionCreditTransferV09 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader93 `xml:"GrpHdr" json:"GrpHdr"`

	// Set of elements providing information specific to the individual credit transfer(s).
	CreditTransferTransactionInformation []*iso20022.CreditTransferTransaction44 `xml:"CdtTrfTxInf" json:"CdtTrfTxInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}
