package admn

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 Document" json:"-"`
	Message *AdmnRequestV01 `xml:"AdmnReq" json:"AdmnReq"`
}

func (d *Document00100101) AddMessage() *AdmnRequestV01 {
	d.Message = new(AdmnRequestV01)
	return d.Message
}

type AdmnRequestV01 struct {
	GroupHeader                 *iso20022.GroupHeader1                  `xml:"GrpHdr" json:"GrpHdr"`
	AdminTransactionInformation []*iso20022.AdminTransactionInformation `xml:"AdmnTxInf" json:"AdmnTxInf"`
}

func (a *AdmnRequestV01) AddGroupHeader() *iso20022.GroupHeader1 {
	a.GroupHeader = new(iso20022.GroupHeader1)
	return a.GroupHeader
}

func (a *AdmnRequestV01) AddCreditTransferTransactionInformation() *iso20022.AdminTransactionInformation {
	newValue := new(iso20022.AdminTransactionInformation)
	a.AdminTransactionInformation = append(a.AdminTransactionInformation, newValue)
	return newValue
}

func (d *Document00100101) String() (result string, ok bool) { return }
