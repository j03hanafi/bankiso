package prxy

import (
	"encoding/xml"
	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.001.001.01 Document" json:"Document"`
	Message *ProxyRegistrationV01 `xml:"PrxyRegn" json:"PrxyRegn"`
}

type ProxyRegistrationV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader59 `xml:"GrpHdr" json:"GrpHdr"`

	Regn *iso20022.ProxyRegistration1 `xml:"Regn" json:"Regn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.BI_SupplementaryData1 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}
