package sese

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document01300105 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.05 Document"`
	Message *PortfolioTransferConfirmationV05 `xml:"PrtflTrfConf"`
}

func (d *Document01300105) AddMessage() *PortfolioTransferConfirmationV05 {
	d.Message = new(PortfolioTransferConfirmationV05)
	return d.Message
}

// Scope
// An executing party, for example, a (old) plan manager (Transferor), sends the PortfolioTransferConfirmation message to the instructing party, for example, a (new) plan manager (Transferee), to confirm the transfer of one or more ISA or portfolio products from the client's account at the old plan manager (Transferor) to the client's account at the new plan manager (Transferee) through a nominee account.
// Usage
// The PortfolioTransferConfirmation message is used to confirm the transfer of one or more ISA or portfolio products.
// The reference of each product transfer confirmation is identified in TransferConfirmationIdentification. The reference of the original product transfer is specified in TransferInstructionReference. The message identification of the PortfolioTransferInstruction message in which the product transfers were conveyed may also be quoted in RelatedReference.
type PortfolioTransferConfirmationV05 struct {

	// Identifies the message.
	MessageReference *iso20022.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information identifying the primary individual investor, eg, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *iso20022.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, eg, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *iso20022.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, eg, name, address, social security number and date of birth.
	OtherIndividualInvestor []*iso20022.IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, eg, name and address.
	PrimaryCorporateInvestor *iso20022.Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, eg, name and address.
	SecondaryCorporateInvestor *iso20022.Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, eg, name and address.
	OtherCorporateInvestor []*iso20022.Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *iso20022.Account15 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *iso20022.Account16 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *iso20022.PartyIdentification2Choice `xml:"Trfee"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *iso20022.CashAccount29 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*iso20022.ISATransfer13 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PortfolioTransferConfirmationV05) AddMessageReference() *iso20022.MessageIdentification1 {
	p.MessageReference = new(iso20022.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferConfirmationV05) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PortfolioTransferConfirmationV05) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PortfolioTransferConfirmationV05) AddRelatedReference() *iso20022.AdditionalReference3 {
	p.RelatedReference = new(iso20022.AdditionalReference3)
	return p.RelatedReference
}

func (p *PortfolioTransferConfirmationV05) AddPrimaryIndividualInvestor() *iso20022.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PortfolioTransferConfirmationV05) AddSecondaryIndividualInvestor() *iso20022.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PortfolioTransferConfirmationV05) AddOtherIndividualInvestor() *iso20022.IndividualPerson8 {
	newValue := new(iso20022.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferConfirmationV05) AddPrimaryCorporateInvestor() *iso20022.Organisation4 {
	p.PrimaryCorporateInvestor = new(iso20022.Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PortfolioTransferConfirmationV05) AddSecondaryCorporateInvestor() *iso20022.Organisation4 {
	p.SecondaryCorporateInvestor = new(iso20022.Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PortfolioTransferConfirmationV05) AddOtherCorporateInvestor() *iso20022.Organisation4 {
	newValue := new(iso20022.Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferConfirmationV05) AddTransferorAccount() *iso20022.Account15 {
	p.TransferorAccount = new(iso20022.Account15)
	return p.TransferorAccount
}

func (p *PortfolioTransferConfirmationV05) AddNomineeAccount() *iso20022.Account16 {
	p.NomineeAccount = new(iso20022.Account16)
	return p.NomineeAccount
}

func (p *PortfolioTransferConfirmationV05) AddTransferee() *iso20022.PartyIdentification2Choice {
	p.Transferee = new(iso20022.PartyIdentification2Choice)
	return p.Transferee
}

func (p *PortfolioTransferConfirmationV05) AddCashAccount() *iso20022.CashAccount29 {
	p.CashAccount = new(iso20022.CashAccount29)
	return p.CashAccount
}

func (p *PortfolioTransferConfirmationV05) AddProductTransfer() *iso20022.ISATransfer13 {
	newValue := new(iso20022.ISATransfer13)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PortfolioTransferConfirmationV05) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	p.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return p.MarketPracticeVersion
}

func (p *PortfolioTransferConfirmationV05) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
func (d *Document01300105) String() (result string, ok bool) { return }
