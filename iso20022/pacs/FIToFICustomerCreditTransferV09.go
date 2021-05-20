package pacs

//
//import (
//	"encoding/xml"
//	"github.com/j03hanafi/bankiso/iso20022"
//)
//
//// TODO: change to pacs.008.001.08
//type Document00800109 struct {
//	XMLName	xml.Name							`xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Document"`
//	Message *FIToFICustomerCreditTransferV09	`xml:"FIToFICstmrCdtTrf"`
//}
//
//func (d *Document00800109) AddMessage() *FIToFICustomerCreditTransferV09 {
//	d.Message = new(FIToFICustomerCreditTransferV09)
//	return d.Message
//}
//
//type FIToFICustomerCreditTransferV09 struct {
//	// TODO: Add common model for GroupHeader93 and CreditTransferTransaction43
//
//	// Set of characteristics shared by all individual transactions included in the message.
//	GroupHeader *iso20022.GroupHeader93 `xml:"GrpHdr"`
//	// TODO: (GroupHeader93) common model for SettlementInstruction7, PaymentTypeInformation28, BranchAndFinancialInstitutionIdentification6
//	// TODO: (SettlementInstruction7) common model for CashAccount38, BranchAndFinancialInstitutionIdentification6
//
//	// Set of elements providing information specific to the individual credit transfer(s).
//	CreditTransferTransactionInformation []*iso20022.CreditTransferTransaction43 `xml:"CdtTrfTxInf"`
//
//	// Additional information that cannot be captured in the structured elements and/or any other specific block.
//	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
//}
//
//func (f *FIToFICustomerCreditTransferV09) AddGroupHeader() *iso20022.GroupHeader93 {
//	f.GroupHeader = new(iso20022.GroupHeader93)
//	return f.GroupHeader
//}
//
//func (f *FIToFICustomerCreditTransferV09) AddCreditTransferTransactionInformation() *iso20022.CreditTransferTransaction43 {
//	newValue := new(iso20022.CreditTransferTransaction43)
//	f.CreditTransferTransactionInformation = append(f.CreditTransferTransactionInformation, newValue)
//	return newValue
//}
//
//func (f *FIToFICustomerCreditTransferV09) AddSupplementaryData() *iso20022.SupplementaryData1 {
//	newValue := new(iso20022.SupplementaryData1)
//	f.SupplementaryData = append(f.SupplementaryData, newValue)
//	return newValue
//}
//func ( d *Document00800109 ) String() (result string, ok bool) { return }
