package admn

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:admn.002.001.01 Document" json:"-"`
	Message *AdmnResponseV01 `xml:"AdmnResp" json:"AdmnResp"`
}

func (d *Document00200101) AddMessage() *AdmnResponseV01 {
	d.Message = new(AdmnResponseV01)
	return d.Message
}

type AdmnResponseV01 struct {
	GroupHeader  *iso20022.GroupHeader1    `xml:"GrpHdr" json:"GrpHdr"`
	AdmnResponse []*iso20022.AdminResponse `xml:"AdmnResponse" json:"AdmnResponse"`
}

func (a *AdmnResponseV01) AddGroupHeader() *iso20022.GroupHeader1 {
	a.GroupHeader = new(iso20022.GroupHeader1)
	return a.GroupHeader
}

func (a *AdmnResponseV01) AddAdmnResponse() *iso20022.AdminResponse {
	newValue := new(iso20022.AdminResponse)
	a.AdmnResponse = append(a.AdmnResponse, newValue)
	return newValue
}

func (d *Document00200101) String() (result string, ok bool) { return }
