package sese

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document03300102 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.033.001.02 Document"`
	Message *SecuritiesFinancingInstructionV02 `xml:"SctiesFincgInstr"`
}

func (d *Document03300102) AddMessage() *SecuritiesFinancingInstructionV02 {
	d.Message = new(SecuritiesFinancingInstructionV02)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesFinancingInstruction to a securities financing transaction account servicer to notify the securities financing transaction account servicer of the details of a repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing transaction to allow the account servicer to manage the settlement and follow-up of the opening and closing leg of the transaction.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct the settlement of securities financing transactions to a central securities depository or another settlement market infrastructure.
// using the relevant elements in the Business Application Header.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingInstructionV02 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionTypeAndAdditionalParameters *iso20022.TransactionTypeAndAdditionalParameters1 `xml:"TxTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages9 `xml:"Lnkgs,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails3 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes20 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount17 `xml:"QtyAndAcctDtls"`

	// Details for the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails1 `xml:"SctiesFincgDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails31 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction3 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties7 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *iso20022.AmountAndDirection2 `xml:"OpngSttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts3 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties9 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingInstructionV02) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SecuritiesFinancingInstructionV02) AddTransactionTypeAndAdditionalParameters() *iso20022.TransactionTypeAndAdditionalParameters1 {
	s.TransactionTypeAndAdditionalParameters = new(iso20022.TransactionTypeAndAdditionalParameters1)
	return s.TransactionTypeAndAdditionalParameters
}

func (s *SecuritiesFinancingInstructionV02) AddNumberCounts() *iso20022.NumberCount1Choice {
	s.NumberCounts = new(iso20022.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesFinancingInstructionV02) AddLinkages() *iso20022.Linkages9 {
	newValue := new(iso20022.Linkages9)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesFinancingInstructionV02) AddTradeDetails() *iso20022.SecuritiesTradeDetails3 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails3)
	return s.TradeDetails
}

func (s *SecuritiesFinancingInstructionV02) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingInstructionV02) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes20 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes20)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingInstructionV02) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount17 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount17)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingInstructionV02) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails1 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails1)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingInstructionV02) AddSettlementParameters() *iso20022.SettlementDetails31 {
	s.SettlementParameters = new(iso20022.SettlementDetails31)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingInstructionV02) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction3 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction3)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingInstructionV02) AddDeliveringSettlementParties() *iso20022.SettlementParties10 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingInstructionV02) AddReceivingSettlementParties() *iso20022.SettlementParties10 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingInstructionV02) AddCashParties() *iso20022.CashParties7 {
	s.CashParties = new(iso20022.CashParties7)
	return s.CashParties
}

func (s *SecuritiesFinancingInstructionV02) AddOpeningSettlementAmount() *iso20022.AmountAndDirection2 {
	s.OpeningSettlementAmount = new(iso20022.AmountAndDirection2)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingInstructionV02) AddOtherAmounts() *iso20022.OtherAmounts3 {
	s.OtherAmounts = new(iso20022.OtherAmounts3)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingInstructionV02) AddOtherBusinessParties() *iso20022.OtherParties9 {
	s.OtherBusinessParties = new(iso20022.OtherParties9)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingInstructionV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
func (d *Document03300102) String() (result string, ok bool) { return }
