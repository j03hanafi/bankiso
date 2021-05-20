package prxy

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00400101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.004.001.01 Document" json:"Document"`
	Message *ProxyLookUpResponseV01 `xml:"PrxyLookUpRspn" json:"PrxyLookUpRspn"`
}

type ProxyLookUpResponseV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader60 `xml:"GrpHdr" json:"GrpHdr"`

	OrgnlGrpInf *iso20022.OriginalGroupInformation3 `xml:"OrgnlGrpInf" json:"OrgnlGrpInf"`

	LkUpRspn *iso20022.ProxyLookUpResponse1 `xml:"LkUpRspn" json:"LkUpRspn"`

	SplmtryData *iso20022.BI_SupplementaryData1 `xml:"SplmtryData" json:"SplmtryData"`
}

func (d *Document00400101) AddMessage() *ProxyLookUpResponseV01 {
	d.Message = new(ProxyLookUpResponseV01)
	return d.Message
}

func (d *Document00400101) String() (result string, ok bool) { return }
