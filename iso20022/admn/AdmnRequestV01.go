package admn

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 Document" json:"-"`
	Message *AdmnRequestV01 `xml:"AdmnReq" json:"AdmnReq"`
}

type AdmnRequestV01 struct {
	GroupHeader                 *iso20022.GroupHeader1                  `xml:"GrpHdr" json:"GrpHdr"`
	AdminTransactionInformation []*iso20022.AdminTransactionInformation `xml:"AdmnTxInf" json:"AdmnTxInf"`
}
