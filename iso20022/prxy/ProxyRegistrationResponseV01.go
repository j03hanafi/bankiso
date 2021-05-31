package prxy

import (
	"encoding/xml"
	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.002.001.01 Document" json:"-"`
	Message *ProxyRegistrationResponseV01 `xml:"PrxyRegnRspn" json:"PrxyRegnRspn"`
}

type ProxyRegistrationResponseV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader60 `xml:"GrpHdr" json:"GrpHdr"`

	OrgnlGrpInf *iso20022.OriginalGroupInformation3 `xml:"OrgnlGrpInf" json:"OrgnlGrpInf"`

	RegnRspn *iso20022.ProxyRegistrationResponse1 `xml:"RegnRspn" json:"RegnRspn"`
}
