package prxy

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.001.001.01 Document" json:"-"`
	Message *ProxyRegistrationV01 `xml:"PrxyRegn" json:"PrxyRegn"`
}

func (d *Document00100101) AddMessage() *ProxyRegistrationV01 {
	d.Message = new(ProxyRegistrationV01)
	return d.Message
}

type ProxyRegistrationV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader59 `xml:"GrpHdr" json:"GrpHdr"`

	Regn *iso20022.ProxyRegistration1 `xml:"Regn" json:"Regn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.BI_SupplementaryData2 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

func (p *ProxyRegistrationV01) AddGroupHeader() *iso20022.GroupHeader59 {
	p.GroupHeader = new(iso20022.GroupHeader59)
	return p.GroupHeader
}

func (p *ProxyRegistrationV01) AddRegn() *iso20022.ProxyRegistration1 {
	p.Regn = new(iso20022.ProxyRegistration1)
	return p.Regn
}

func (p *ProxyRegistrationV01) AddSupplementaryData() *iso20022.BI_SupplementaryData2 {
	newValue := new(iso20022.BI_SupplementaryData2)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

func (d *Document00100101) String() (result string, ok bool) { return }
