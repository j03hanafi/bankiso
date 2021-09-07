package prxy

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00400101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:prxy.004.001.01 Document" json:"Document"`
	Message *ProxyLookUpResponseV01 `xml:"PrxyLookUpRspn" json:"PrxyLookUpRspn"`
}

func (d *Document00400101) AddMessage() *ProxyLookUpResponseV01 {
	d.Message = new(ProxyLookUpResponseV01)
	return d.Message
}

type ProxyLookUpResponseV01 struct {
	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader60             `xml:"GrpHdr" json:"GrpHdr"`
	OrgnlGrpInf *iso20022.OriginalGroupInformation3 `xml:"OrgnlGrpInf" json:"OrgnlGrpInf"`
	LkUpRspn    *iso20022.ProxyLookUpResponse1      `xml:"LkUpRspn" json:"LkUpRspn"`
	SplmtryData []*iso20022.BI_SupplementaryData1   `xml:"SplmtryData" json:"SplmtryData"`
}

func (p *ProxyLookUpResponseV01) AddGroupHeader() *iso20022.GroupHeader60 {
	p.GroupHeader = new(iso20022.GroupHeader60)
	return p.GroupHeader
}

func (p *ProxyLookUpResponseV01) AddOrgnlGrpInf() *iso20022.OriginalGroupInformation3 {
	p.OrgnlGrpInf = new(iso20022.OriginalGroupInformation3)
	return p.OrgnlGrpInf
}

func (p *ProxyLookUpResponseV01) AddLkUpRspn() *iso20022.ProxyLookUpResponse1 {
	p.LkUpRspn = new(iso20022.ProxyLookUpResponse1)
	return p.LkUpRspn
}

func (p *ProxyLookUpResponseV01) AddSplmtryData() *iso20022.BI_SupplementaryData1 {
	newValue := new(iso20022.BI_SupplementaryData1)
	p.SplmtryData = append(p.SplmtryData, newValue)
	return p.SplmtryData
}

func (d *Document00400101) String() (result string, ok bool) { return }
