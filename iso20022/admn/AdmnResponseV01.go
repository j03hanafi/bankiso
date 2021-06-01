package admn

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:admn.002.001.01 Document" json:"-"`
	Message *AdmnResponseV01 `xml:"AdmnResp" json:"AdmnResp"`
}

type AdmnResponseV01 struct {
	GroupHeader  *iso20022.GroupHeader1    `xml:"GrpHdr" json:"GrpHdr"`
	AdmnResponse []*iso20022.AdminResponse `xml:"AdmnResponse" json:"AdmnResponse"`
}
