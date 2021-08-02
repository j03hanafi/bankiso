package pacs

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00200110 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Document" json:"-"`
	Message *FIToFIPaymentStatusReportV10 `xml:"FIToFIPmtStsRpt" json:"FIToFIPmtStsRpt"`
}

func (d *Document00200110) AddMessage() *FIToFIPaymentStatusReportV10 {
	d.Message = new(FIToFIPaymentStatusReportV10)
	return d.Message
}

type FIToFIPaymentStatusReportV10 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *iso20022.GroupHeader91 `xml:"GrpHdr" json:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus []*iso20022.OriginalGroupHeader17 `xml:"OrgnlGrpInfAndSts,omitempty" json:"OrgnlGrpInfAndSts,omitempty"`

	// Information concerning the original transactions, to which the status report message refers.
	TransactionInformationAndStatus []*iso20022.PaymentTransaction110 `xml:"TxInfAndSts,omitempty" json:"TxInfAndSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

func (f *FIToFIPaymentStatusReportV10) AddGroupHeader() *iso20022.GroupHeader91 {
	f.GroupHeader = new(iso20022.GroupHeader91)
	return f.GroupHeader
}

func (f *FIToFIPaymentStatusReportV10) AddOriginalGroupInformationAndStatus() *iso20022.OriginalGroupHeader17 {
	newValue := new(iso20022.OriginalGroupHeader17)
	f.OriginalGroupInformationAndStatus = append(f.OriginalGroupInformationAndStatus, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV10) AddTransactionInformationAndStatus() *iso20022.PaymentTransaction110 {
	newValue := new(iso20022.PaymentTransaction110)
	f.TransactionInformationAndStatus = append(f.TransactionInformationAndStatus, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV10) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}

func (d *Document00200110) String() (result string, ok bool) { return }
