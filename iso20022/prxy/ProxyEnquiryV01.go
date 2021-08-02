package prxy

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00500101 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.005.001.01 Document" json:"Document"`
	Message *ProxyEnquiryV01 `xml:"PrxyNqryReq" json:"PrxyNqryReq"`
}

func (d *Document00500101) AddMessage() *ProxyEnquiryV01 {
	d.Message = new(ProxyEnquiryV01)
	return d.Message
}

type ProxyEnquiryV01 struct {
	GroupHeader *iso20022.GroupHeader59       `xml:"GrpHdr" json:"GrpHdr"`
	Nqry        *iso20022.ProxyEnquiryChoice1 `xml:"Nqry" json:"Nqry"`
}

func (p *ProxyEnquiryV01) AddGroupHeader() *iso20022.GroupHeader59 {
	p.GroupHeader = new(iso20022.GroupHeader59)
	return p.GroupHeader
}

func (p *ProxyEnquiryV01) AddNqry() *iso20022.ProxyEnquiryChoice1 {
	p.Nqry = new(iso20022.ProxyEnquiryChoice1)
	return p.Nqry
}

func (d *Document00500101) String() (result string, ok bool) { return }
