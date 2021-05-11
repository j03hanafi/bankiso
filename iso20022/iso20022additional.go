package iso20022

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader91 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId" json:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm" json:"CreDtTm"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"json:"InstdAgt,omitempty"`
}

func (g *GroupHeader53) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader53) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader53) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification6 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstructingAgent
}

func (g *GroupHeader53) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification6 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstructedAgent
}

type BranchAndFinancialInstitutionIdentification6 struct {

	// Unique and unambiguous identification of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification18 `xml:"FinInstnId" json:"FinInstnId,omitempty"`

	// Identifies a specific branch of a financial institution.
	//
	// Usage: This component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BranchIdentification *BranchData3 `xml:"BrnchId,omitempty" json:"BrnchId,omitempty"`
}

// Set of elements used to identify a financial institution.
type FinancialInstitutionIdentification18 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI,omitempty" json:"BICFI,omitempty"`

	// Information used to identify a member within a clearing system.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:"ClrSysMmbId,omitempty"`

	LEI *LEIIdentifier `xml:"LEI,omitempty" json:"LEI,omitempty"`

	// Name by which an agent is known and which is usually used to identify that agent.
	Name *Max140Text `xml:"Nm,omitempty" json:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress24 `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`

	// Unique identification of an agent, as assigned by an institution, using an identification scheme.
	Other *GenericFinancialIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress24 struct {

	// Identifies the nature of the postal address.
	AddressType *AddressType3Choice `xml:"AdrTp,omitempty" json:"AdrTp,omitempty"`

	// Identification of a division of a large organisation or building.
	Department *Max70Text `xml:"Dept,omitempty" json:"Dept,omitempty"`

	// Identification of a sub-division of a large organisation or building.
	SubDepartment *Max70Text `xml:"SubDept,omitempty" json:"SubDept,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty" json:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty" json:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty" json:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty" json:"TwnNm,omitempty"`

	// Identifies a subdivision of a country such as state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty" json:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry,omitempty" json:"Ctry,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty" json:"AdrLine,omitempty"`
}

func (p *PostalAddress24) SetAddressType(value AddressType3Choice) {
	p.AddressType = (*AddressType3Choice)(&value)
}

func (p *PostalAddress24) SetDepartment(value string) {
	p.Department = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetSubDepartment(value string) {
	p.SubDepartment = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress24) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress24) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PostalAddress24) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

type AddressType3Choice struct {
	Cd    *AddressType2Code        `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry *GenericIdentification30 `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// Information that locates and identifies a specific branch of a financial institution.
type BranchData3 struct {

	// Unique and unambiguous identification of a branch of a financial institution.
	Identification *Max35Text `xml:"Id,omitempty" json:"Id,omitempty"`

	LEI *LEIIdentifier `xml:"LEI,omitempty" json:"LEI,omitempty"`

	// Name by which an agent is known and which is usually used to identify that agent.
	Name *Max140Text `xml:"Nm,omitempty" json:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress24 `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

func (b *BranchData2) SetIdentification(value string) {
	b.Identification = (*Max35Text)(&value)
}

func (b *BranchData2) SetName(value string) {
	b.Name = (*Max140Text)(&value)
}

func (b *BranchData2) AddPostalAddress() *PostalAddress6 {
	b.PostalAddress = new(PostalAddress6)
	return b.PostalAddress
}

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader17 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group of transactions.
	GroupStatus *ExternalPaymentGroupStatus1Code `xml:"GrpSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation12 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupHeader17) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader17) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader17) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader17) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader17) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader17) SetGroupStatus(value string) {
	o.GroupStatus = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalGroupHeader17) AddStatusReasonInformation() *StatusReasonInformation12 {
	newValue := new(StatusReasonInformation12)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupHeader17) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

type ExternalPaymentGroupStatus1Code string

// Set of elements used to provide information on the status reason of the transaction.
type StatusReasonInformation12 struct {

	// Party that issues the status.
	Originator *PartyIdentification135 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the status report.
	Reason *StatusReason6Choice `xml:"Rsn,omitempty"`

	// Further details on the status reason.
	//
	// Usage: Additional information can be used for several purposes such as the reporting of repaired information.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (s *StatusReasonInformation9) AddOriginator() *PartyIdentification135 {
	s.Originator = new(PartyIdentification135)
	return s.Originator
}

func (s *StatusReasonInformation9) AddReason() *StatusReason6Choice {
	s.Reason = new(StatusReason6Choice)
	return s.Reason
}

func (s *StatusReasonInformation9) AddAdditionalInformation(value string) {
	s.AdditionalInformation = append(s.AdditionalInformation, (*Max105Text)(&value))
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification135 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress24 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party38Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *Contact4 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification135) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification135) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification135) AddIdentification() *Party38Choice {
	p.Identification = new(Party38Choice)
	return p.Identification
}

func (p *PartyIdentification135) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification135) AddContactDetails() *Contact4 {
	p.ContactDetails = new(Contact4)
	return p.ContactDetails
}

// Nature or use of the account.
type Party38Choice struct {

	// Unique and unambiguous way to identify an organisation.
	OrganisationIdentification *OrganisationIdentification29 `xml:"OrgId"`

	// Unique and unambiguous identification of a person, eg, passport.
	PrivateIdentification *PersonIdentification13 `xml:"PrvtId"`
}

func (p *Party38Choice) AddOrganisationIdentification() *OrganisationIdentification29 {
	p.OrganisationIdentification = new(OrganisationIdentification29)
	return p.OrganisationIdentification
}

func (p *Party38Choice) AddPrivateIdentification() *PersonIdentification13 {
	p.PrivateIdentification = new(PersonIdentification13)
	return p.PrivateIdentification
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification29 struct {

	// Code allocated to a financial institution or non financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty"`

	LEI *LEIIdentifier `xml:"LEI,omitempty" json:"LEI,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification29) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification29) AddOther() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Other = append(o.Other, newValue)
	return newValue
}

type PersonIdentification13 struct {

	// Date and place of birth of a person.
	DateAndPlaceOfBirth *DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty"`

	// Unique identification of a person, as assigned by an institution, using an identification scheme.
	Other []*GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

func (p *PersonIdentification13) AddDateAndPlaceOfBirth() *DateAndPlaceOfBirth {
	p.DateAndPlaceOfBirth = new(DateAndPlaceOfBirth)
	return p.DateAndPlaceOfBirth
}

func (p *PersonIdentification13) AddOther() *GenericPersonIdentification1 {
	newValue := new(GenericPersonIdentification1)
	p.Other = append(p.Other, newValue)
	return newValue
}

// Communication device number or electronic address used for communication.
type Contact4 struct {

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix2Code `xml:"NmPrfx,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb,omitempty"`

	// Collection of information that identifies a mobile phone number, as defined by telecom services.
	MobileNumber *PhoneNumber `xml:"MobNb,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max2048Text `xml:"EmailAdr,omitempty"`

	EmailPurp *Max35Text `xml:"EmailPurp,omitempty"`

	JobTitl *Max35Text `xml:"JobTitl,omitempty"`

	Rspnsblty *Max35Text `xml:"Rspnsblty,omitempty"`

	// Contact details in an other form.
	Other []*OtherContact1 `xml:"Othr,omitempty"`

	// Preferred method used to reach the technical contact.
	PreferredMethod *PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty"`
}

func (c *Contact4) SetNamePrefix(value string) {
	c.NamePrefix = (*NamePrefix1Code)(&value)
}

func (c *Contact4) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *Contact4) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *Contact4) SetMobileNumber(value string) {
	c.MobileNumber = (*PhoneNumber)(&value)
}

func (c *Contact4) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *Contact4) SetEmailAddress(value string) {
	c.EmailAddress = (*Max2048Text)(&value)
}

func (c *Contact4) AddOther() *OtherContact1 {
	newValue := new(OtherContact1)
	c.Other = append(c.Other, newValue)
	return newValue
}

func (c *Contact4) SetPreferredMethod(value string) {
	c.PreferredMethod = (*PreferredContactMethod1Code)(&value)
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

// Set of elements used to provide detailed information on the number of transactions that are reported with a specific transaction status.
type NumberOfTransactionsPerStatus5 struct {

	// Number of individual transactions contained in the message, detailed per status.
	DetailedNumberOfTransactions *Max15NumericText `xml:"DtldNbOfTxs"`

	// Common transaction status for all individual transactions reported.
	DetailedStatus *ExternalPaymentGroupStatus1Code `xml:"DtldSts"`

	// Total of all individual amounts included in the message, irrespective of currencies, detailed per status.
	DetailedControlSum *DecimalNumber `xml:"DtldCtrlSum,omitempty"`
}

func (n *NumberOfTransactionsPerStatus5) SetDetailedNumberOfTransactions(value string) {
	n.DetailedNumberOfTransactions = (*Max15NumericText)(&value)
}

func (n *NumberOfTransactionsPerStatus5) SetDetailedStatus(value string) {
	n.DetailedStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (n *NumberOfTransactionsPerStatus5) SetDetailedControlSum(value string) {
	n.DetailedControlSum = (*DecimalNumber)(&value)
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction110 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalGroupInformation *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *ExternalPaymentTransactionStatus1Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation12 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges7 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference28 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction110) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) AddOriginalGroupInformation() *OriginalGroupInformation29 {
	p.OriginalGroupInformation = new(OriginalGroupInformation29)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction110) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction110) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction110) AddChargesInformation() *Charges7 {
	newValue := new(Charges7)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction110) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction110) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification6 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification6)
	return p.InstructingAgent
}

func (p *PaymentTransaction110) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification6 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification6)
	return p.InstructedAgent
}

func (p *PaymentTransaction110) AddOriginalTransactionReference() *OriginalTransactionReference28 {
	p.OriginalTransactionReference = new(OriginalTransactionReference28)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction110) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation29 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Original date and time at which the message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`
}

func (o *OriginalGroupInformation29) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation29) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation29) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

// May be no more than 4 items long
type ExternalPaymentTransactionStatus1Code string

// Set of elements used to provide information on the charges related to the payment transaction.
type Charges7 struct {

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Agent that takes the transaction charges or to which the transaction charges are due.
	Agent *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

func (c *Charges7) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges7) AddAgent() *BranchAndFinancialInstitutionIdentification6 {
	c.Agent = new(BranchAndFinancialInstitutionIdentification6)
	return c.Agent
}

// Key elements used to refer the original transaction.
type OriginalTransactionReference28 struct {

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType4Choice `xml:"Amt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *DateAndDateTime2Choice `xml:"ReqdExctnDt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification135 `xml:"CdtrSchmeId,omitempty"`

	// Specifies the details on how the settlement of the original transaction(s) between the instructing agent and the instructed agent was completed.
	SettlementInformation *SettlementInstruction7 `xml:"SttlmInf,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod4Code `xml:"PmtMtd,omitempty"`

	// Provides further details of the mandate signed between the creditor and the debtor.
	MandateRelatedInformation *MandateRelatedInformation10 `xml:"MndtRltdInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`
}

func (o *OriginalTransactionReference28) SetInterbankSettlementAmount(value, currency string) {
	o.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalTransactionReference28) AddAmount() *AmountType4Choice {
	o.Amount = new(AmountType4Choice)
	return o.Amount
}

func (o *OriginalTransactionReference28) SetInterbankSettlementDate(value string) {
	o.InterbankSettlementDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference28) SetRequestedCollectionDate(value string) {
	o.RequestedCollectionDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference28) SetRequestedExecutionDate(value string) {
	o.RequestedExecutionDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference28) AddCreditorSchemeIdentification() *PartyIdentification43 {
	o.CreditorSchemeIdentification = new(PartyIdentification43)
	return o.CreditorSchemeIdentification
}

func (o *OriginalTransactionReference28) AddSettlementInformation() *SettlementInstruction7 {
	o.SettlementInformation = new(SettlementInstruction7)
	return o.SettlementInformation
}

func (o *OriginalTransactionReference28) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	o.PaymentTypeInformation = new(PaymentTypeInformation25)
	return o.PaymentTypeInformation
}

func (o *OriginalTransactionReference28) SetPaymentMethod(value string) {
	o.PaymentMethod = (*PaymentMethod4Code)(&value)
}

func (o *OriginalTransactionReference28) AddMandateRelatedInformation() *MandateRelatedInformation10 {
	o.MandateRelatedInformation = new(MandateRelatedInformation10)
	return o.MandateRelatedInformation
}

func (o *OriginalTransactionReference28) AddRemittanceInformation() *RemittanceInformation11 {
	o.RemittanceInformation = new(RemittanceInformation11)
	return o.RemittanceInformation
}

func (o *OriginalTransactionReference28) AddUltimateDebtor() *PartyIdentification43 {
	o.UltimateDebtor = new(PartyIdentification43)
	return o.UltimateDebtor
}

func (o *OriginalTransactionReference28) AddDebtor() *PartyIdentification43 {
	o.Debtor = new(PartyIdentification43)
	return o.Debtor
}

func (o *OriginalTransactionReference28) AddDebtorAccount() *CashAccount24 {
	o.DebtorAccount = new(CashAccount24)
	return o.DebtorAccount
}

func (o *OriginalTransactionReference28) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification6 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification6)
	return o.DebtorAgent
}

func (o *OriginalTransactionReference28) AddDebtorAgentAccount() *CashAccount24 {
	o.DebtorAgentAccount = new(CashAccount24)
	return o.DebtorAgentAccount
}

func (o *OriginalTransactionReference28) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification6 {
	o.CreditorAgent = new(BranchAndFinancialInstitutionIdentification6)
	return o.CreditorAgent
}

func (o *OriginalTransactionReference28) AddCreditorAgentAccount() *CashAccount24 {
	o.CreditorAgentAccount = new(CashAccount24)
	return o.CreditorAgentAccount
}

func (o *OriginalTransactionReference28) AddCreditor() *PartyIdentification43 {
	o.Creditor = new(PartyIdentification43)
	return o.Creditor
}

func (o *OriginalTransactionReference28) AddCreditorAccount() *CashAccount24 {
	o.CreditorAccount = new(CashAccount24)
	return o.CreditorAccount
}

func (o *OriginalTransactionReference28) AddUltimateCreditor() *PartyIdentification43 {
	o.UltimateCreditor = new(PartyIdentification43)
	return o.UltimateCreditor
}

// Choice between a date or a date and time format.
type DateAndDateTime2Choice struct {

	// Numeric representation of the day of the month and year.
	Date *ISODate `xml:"Dt,omitempty"`

	// Numeric representation of time of the day and the day of the month and year.
	DateTime *ISODateTime `xml:"DtTm,omitempty"`
}

func (d *DateAndDateTime2Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateAndDateTime2Choice) SetDateTime(value string) {
	d.DateTime = (*ISODateTime)(&value)
}

// Provides further details on the settlement of the instruction.
type SettlementInstruction7 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount24 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount24 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount24 `xml:"InstdRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If ThirdReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ThrdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThirdReimbursementAgentAccount *CashAccount24 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInstruction7) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInstruction7) AddSettlementAccount() *CashAccount24 {
	s.SettlementAccount = new(CashAccount24)
	return s.SettlementAccount
}

func (s *SettlementInstruction7) AddClearingSystem() *ClearingSystemIdentification3Choice {
	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
	return s.ClearingSystem
}

func (s *SettlementInstruction7) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInstruction7) AddInstructingReimbursementAgentAccount() *CashAccount24 {
	s.InstructingReimbursementAgentAccount = new(CashAccount24)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInstruction7) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInstruction7) AddInstructedReimbursementAgentAccount() *CashAccount24 {
	s.InstructedReimbursementAgentAccount = new(CashAccount24)
	return s.InstructedReimbursementAgentAccount
}

func (s *SettlementInstruction7) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.ThirdReimbursementAgent
}

func (s *SettlementInstruction7) AddThirdReimbursementAgentAccount() *CashAccount24 {
	s.ThirdReimbursementAgentAccount = new(CashAccount24)
	return s.ThirdReimbursementAgentAccount
}

//
//// Set of characteristics shared by all individual transactions included in the message.
//type GroupHeader93 struct {
//
//	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
//	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
//	MessageIdentification *Max35Text `xml:"MsgId"`
//
//	// Date and time at which the message was created.
//	CreationDateTime *ISODateTime `xml:"CreDtTm"`
//
//	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
//	// Usage: Batch booking is used to request and not order a possible batch booking.
//	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`
//
//	// Number of individual transactions contained in the message.
//	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`
//
//	// Total of all individual amounts included in the message, irrespective of currencies.
//	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`
//
//	// Total amount of money moved between the instructing agent and the instructed agent.
//	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`
//
//	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
//	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`
//
//	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
//	SettlementInformation *SettlementInstruction7 `xml:"SttlmInf"`
//
//	// Set of elements used to further specify the type of transaction.
//	PaymentTypeInformation *PaymentTypeInformation28 `xml:"PmtTpInf,omitempty"`
//
//	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
//	InstructingAgent *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
//
//	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
//	InstructedAgent *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
//}
//
//func (g *GroupHeader93) SetMessageIdentification(value string) {
//	g.MessageIdentification = (*Max35Text)(&value)
//}
//
//func (g *GroupHeader93) SetCreationDateTime(value string) {
//	g.CreationDateTime = (*ISODateTime)(&value)
//}
//
//func (g *GroupHeader93) SetBatchBooking(value string) {
//	g.BatchBooking = (*BatchBookingIndicator)(&value)
//}
//
//func (g *GroupHeader93) SetNumberOfTransactions(value string) {
//	g.NumberOfTransactions = (*Max15NumericText)(&value)
//}
//
//func (g *GroupHeader93) SetControlSum(value string) {
//	g.ControlSum = (*DecimalNumber)(&value)
//}
//
//func (g *GroupHeader93) SetTotalInterbankSettlementAmount(value, currency string) {
//	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
//}
//
//func (g *GroupHeader93) SetInterbankSettlementDate(value string) {
//	g.InterbankSettlementDate = (*ISODate)(&value)
//}
//
//func (g *GroupHeader93) AddSettlementInformation() *SettlementInstruction7 {
//	g.SettlementInformation = new(SettlementInstruction7)
//	return g.SettlementInformation
//}
//
//func (g *GroupHeader93) AddPaymentTypeInformation() *PaymentTypeInformation28 {
//	g.PaymentTypeInformation = new(PaymentTypeInformation28)
//	return g.PaymentTypeInformation
//}
//
//func (g *GroupHeader93) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification6 {
//	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification6)
//	return g.InstructingAgent
//}
//
//func (g *GroupHeader93) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification6 {
//	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification6)
//	return g.InstructedAgent
//}
//
//type SettlementInstruction7 struct {
//
//	// Method used to settle the (batch of) payment instructions.
//	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`
//
//	// A specific purpose account used to post debit and credit entries as a result of the transaction.
//	SettlementAccount *CashAccount38 `xml:"SttlmAcct,omitempty"`
//
//	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
//	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`
//
//	// Agent through which the instructing agent will reimburse the instructed agent.
//	//
//	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
//	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty"`
//
//	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
//	InstructingReimbursementAgentAccount *CashAccount38 `xml:"InstgRmbrsmntAgtAcct,omitempty"`
//
//	// Agent at which the instructed agent will be reimbursed.
//	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
//	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
//	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty"`
//
//	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
//	InstructedReimbursementAgentAccount *CashAccount38 `xml:"InstdRmbrsmntAgtAcct,omitempty"`
//
//	// Agent at which the instructed agent will be reimbursed.
//	// Usage: If ThirdReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
//	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty"`
//
//	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
//	ThirdReimbursementAgentAccount *CashAccount38 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
//}
//
//func (s *SettlementInstruction7) SetSettlementMethod(value string) {
//	s.SettlementMethod = (*SettlementMethod1Code)(&value)
//}
//
//func (s *SettlementInstruction7) AddSettlementAccount() *CashAccount38 {
//	s.SettlementAccount = new(CashAccount38)
//	return s.SettlementAccount
//}
//
//func (s *SettlementInstruction7) AddClearingSystem() *ClearingSystemIdentification3Choice {
//	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
//	return s.ClearingSystem
//}
//
//func (s *SettlementInstruction7) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification6 {
//	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification6)
//	return s.InstructingReimbursementAgent
//}
//
//func (s *SettlementInstruction7) AddInstructingReimbursementAgentAccount() *CashAccount38 {
//	s.InstructingReimbursementAgentAccount = new(CashAccount38)
//	return s.InstructingReimbursementAgentAccount
//}
//
//func (s *SettlementInstruction7) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification6 {
//	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification6)
//	return s.InstructedReimbursementAgent
//}
//
//func (s *SettlementInstruction7) AddInstructedReimbursementAgentAccount() *CashAccount38 {
//	s.InstructedReimbursementAgentAccount = new(CashAccount38)
//	return s.InstructedReimbursementAgentAccount
//}
//
//func (s *SettlementInstruction7) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification6 {
//	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification6)
//	return s.ThirdReimbursementAgent
//}
//
//func (s *SettlementInstruction7) AddThirdReimbursementAgentAccount() *CashAccount38 {
//	s.ThirdReimbursementAgentAccount = new(CashAccount38)
//	return s.ThirdReimbursementAgentAccount
//}
