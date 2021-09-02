package pacs

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document02800104 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04 Document" json:"-"`
	Message *FIToFIPaymentStatusRequestV04 `xml:"FItoFIPmtStsReq" json:"FItoFIPmtStsReq"`
}

// res.AddMessage()
func (d *Document02800104) AddMessage() *FIToFIPaymentStatusRequestV04 {
	d.Message = new(FIToFIPaymentStatusRequestV04)
	return d.Message
}

type FIToFIPaymentStatusRequestV04 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader91 `xml:"GrpHdr" json:"GrpHdr"`

	OrgnlGrpInf []*iso20022.OriginalGroupInformation27 `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`

	TxInf []*iso20022.PaymentTransaction121 `xml:"TxInf,omitempty" json:"TxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

// res.AddMessage().AddGroupHeader()
func (f *FIToFIPaymentStatusRequestV04) AddGroupHeader() *iso20022.GroupHeader91 {
	f.GroupHeader = new(iso20022.GroupHeader91)
	return f.GroupHeader
}

func (f *FIToFIPaymentStatusRequestV04) AddOrgnlGrpInf() *iso20022.OriginalGroupInformation27 {
	newValue := new(iso20022.OriginalGroupInformation27)
	f.OrgnlGrpInf = append(f.OrgnlGrpInf, newValue)
	return newValue
}

// res.AddMessage(). AddTxInf()
func (f *FIToFIPaymentStatusRequestV04) AddTxInf() *iso20022.PaymentTransaction121 {
	newValue := new(iso20022.PaymentTransaction121)
	f.TxInf = append(f.TxInf, newValue)
	return newValue
}
