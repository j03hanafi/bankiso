package prxy

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00300101 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.003.001.01 Document" json:"Document"`
	Message *ProxyLookUpV01 `xml:"PrxyLookUp" json:"PrxyLookUp"`
}

func (d *Document00300101) AddMessage() *ProxyLookUpV01 {
	d.Message = new(ProxyLookUpV01)
	return d.Message
}

type ProxyLookUpV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader69      `xml:"GrpHdr" json:"GrpHdr"`
	LookUp      *iso20022.ProxyLookUpChoice1 `xml:"LookUp" json:"LookUp"`
}

func (p *ProxyLookUpV01) AddGroupHeader() *iso20022.GroupHeader69 {
	p.GroupHeader = new(iso20022.GroupHeader69)
	return p.GroupHeader
}

func (p *ProxyLookUpV01) AddLookUp() *iso20022.ProxyLookUpChoice1 {
	p.LookUp = new(iso20022.ProxyLookUpChoice1)
	return p.LookUp
}

func (d *Document00300101) String() (result string, ok bool) { return }
