package prxy

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00600101 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.006.001.01 Document" json:"Document"`
	Message *ProxyEnquiryV01 `xml:"PrxyNqryRspn" json:"PrxyNqryRspn"`
}

type ProxyEnquiryResponseV01 struct {
	GroupHeader *iso20022.GroupHeader60 `xml:"GrpHdr" json:"GrpHdr"`

	OrgnlGrpInf *iso20022.OriginalGroupInformation3 `xml:"OrgnlGrpInf" json:"OrgnlGrpInf"`

	NqryRspn *iso20022.ProxyEnquiryResponse1 `xml:"NqryRspn" json:"NqryRspn"`
}

func (d *Document00600101) AddMessage() *ProxyEnquiryV01 {
	d.Message = new(ProxyEnquiryV01)
	return d.Message
}

func (d *Document00600101) String() (result string, ok bool) { return }
