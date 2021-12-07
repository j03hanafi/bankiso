package prxy

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.002.001.01 Document" json:"-"`
	Message *ProxyRegistrationResponseV01 `xml:"PrxyRegnRspn" json:"PrxyRegnRspn"`
}

func (d *Document00200101) AddMessage() *ProxyRegistrationResponseV01 {
	d.Message = new(ProxyRegistrationResponseV01)
	return d.Message
}

type ProxyRegistrationResponseV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader60 `xml:"GrpHdr" json:"GrpHdr"`

	OrgnlGrpInf *iso20022.OriginalGroupInformation3 `xml:"OrgnlGrpInf" json:"OrgnlGrpInf"`

	RegnRspn *iso20022.ProxyRegistrationResponse1 `xml:"RegnRspn" json:"RegnRspn"`

	SplmtryData []*iso20022.BI_SupplementaryData2 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

func (p *ProxyRegistrationResponseV01) AddGroupHeader() *iso20022.GroupHeader60 {
	p.GroupHeader = new(iso20022.GroupHeader60)
	return p.GroupHeader
}

func (p *ProxyRegistrationResponseV01) AddOrgnlGrpInf() *iso20022.OriginalGroupInformation3 {
	p.OrgnlGrpInf = new(iso20022.OriginalGroupInformation3)
	return p.OrgnlGrpInf
}

func (p *ProxyRegistrationResponseV01) AddRegnRspn() *iso20022.ProxyRegistrationResponse1 {
	p.RegnRspn = new(iso20022.ProxyRegistrationResponse1)
	return p.RegnRspn
}

func (d *Document00200101) String() (result string, ok bool) { return }
