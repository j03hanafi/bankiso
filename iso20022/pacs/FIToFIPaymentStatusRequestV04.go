package pacs

import (
	"encoding/xml"
	"github.com/figassis/bankiso/iso20022"
)

type Document02800104 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04 Document" json:"Document"`
	Message *FIToFIPaymentStatusRequestV04 `xml:"FIToFIPmtStsReq" json:"FIToFIPmtStsReq"`
}

type FIToFIPaymentStatusRequestV04 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader91 `xml:"GrpHdr" json:"GrpHdr"`

	OrgnlGrpInf []*iso20022.OriginalGroupInformation27 `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`

	TxInf []*iso20022.PaymentTransaction121 `xml:"TxInf,omitempty" json:"TxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}
