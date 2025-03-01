package semt

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document01800207 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.018.002.07 Document"`
	Message *SecuritiesTransactionPendingReport002V07 `xml:"SctiesTxPdgRpt"`
}

func (d *Document01800207) AddMessage() *SecuritiesTransactionPendingReport002V07 {
	d.Message = new(SecuritiesTransactionPendingReport002V07)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesTransactionPendingReport to an account owner to provide, as at a specified time, the details of pending increases and decreases of holdings, for all or selected securities in a specified safekeeping account, for all or selected reasons why the transaction is pending.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The statement may also include future settlement or forward transactions which have become binding on the account owner.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesTransactionPendingReport002V07 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *iso20022.Statement55 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Status information.
	Status []*iso20022.StatusAndReason30 `xml:"Sts,omitempty"`

	// Details of the transactions reported.
	Transactions []*iso20022.Transaction49 `xml:"Txs,omitempty"`
}

func (s *SecuritiesTransactionPendingReport002V07) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPendingReport002V07) AddStatementGeneralDetails() *iso20022.Statement55 {
	s.StatementGeneralDetails = new(iso20022.Statement55)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPendingReport002V07) AddAccountOwner() *iso20022.PartyIdentification109 {
	s.AccountOwner = new(iso20022.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPendingReport002V07) AddSafekeepingAccount() *iso20022.SecuritiesAccount27 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount27)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPendingReport002V07) AddStatus() *iso20022.StatusAndReason30 {
	newValue := new(iso20022.StatusAndReason30)
	s.Status = append(s.Status, newValue)
	return newValue
}

func (s *SecuritiesTransactionPendingReport002V07) AddTransactions() *iso20022.Transaction49 {
	newValue := new(iso20022.Transaction49)
	s.Transactions = append(s.Transactions, newValue)
	return newValue
}
func (d *Document01800207) String() (result string, ok bool) { return }
