package prxy

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document90100101 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.901.001.01 Document" json:"Document"`
	Message *ProxyEnquiryV01 `xml:"PrxyNqryRspn" json:"PrxyNqryRspn"`
}

func (d *Document90100101) AddMessage() *ProxyEnquiryV01 {
	d.Message = new(ProxyEnquiryV01)
	return d.Message
}

type ProxyNtfctnV01struct struct {
	GroupHeader *iso20022.GroupHeader60 `xml:"GrpHdr" json:"GrpHdr"`
	Ntfctn      *iso20022.ProxyNtfctn1  `xml:"Ntfctn" json:"Ntfctn"`
}

func (p *ProxyNtfctnV01struct) AddGroupHeader() *iso20022.GroupHeader60 {
	p.GroupHeader = new(iso20022.GroupHeader60)
	return p.GroupHeader
}

func (p *ProxyNtfctnV01struct) AddNtfctn() *iso20022.ProxyNtfctn1 {
	p.Ntfctn = new(iso20022.ProxyNtfctn1)
	return p.Ntfctn
}

func (d *Document90100101) String() (result string, ok bool) { return }
