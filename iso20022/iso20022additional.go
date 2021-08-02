package iso20022

// TODO: pacs.002.001.10, json tag

type GroupHeader91 struct {
	MsgId    *Max35Text                                    `xml:"MsgId" json:"MsgId"`
	CreDtTm  *ISODateTime                                  `xml:"CreDtTm" json:"CreDtTm"`
	InstgAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
}

func (g *GroupHeader91) SetMsgId(value string) {
	g.MsgId = (*Max35Text)(&value)
}

func (g *GroupHeader91) SetCreDtTm(value string) {
	g.CreDtTm = (*ISODateTime)(&value)
}

func (g *GroupHeader91) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	g.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstgAgt
}

func (g *GroupHeader91) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	g.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstdAgt
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId *FinancialInstitutionIdentification18 `xml:"FinInstnId" json:"FinInstnId"`
	BrnchId    *BranchData3                          `xml:"BrnchId,omitempty" json:"BrnchId,omitempty"`
}

func (b *BranchAndFinancialInstitutionIdentification6) AddFinInstnId() *FinancialInstitutionIdentification18 {
	b.FinInstnId = new(FinancialInstitutionIdentification18)
	return b.FinInstnId
}

func (b *BranchAndFinancialInstitutionIdentification6) AddBrnchId() *BranchData3 {
	b.BrnchId = new(BranchData3)
	return b.BrnchId
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *BICFIIdentifier                     `xml:"BICFI,omitempty" json:"BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:"ClrSysMmbId,omitempty"`
	LEI         *LEIIdentifier                       `xml:"LEI,omitempty" json:"LEI,omitempty"`
	Nm          *Max140Text                          `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

func (f *FinancialInstitutionIdentification18) SetBICFI(value string) {
	f.BICFI = (*BICFIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification18) AddClrSysMmbId() *ClearingSystemMemberIdentification2 {
	f.ClrSysMmbId = new(ClearingSystemMemberIdentification2)
	return f.ClrSysMmbId
}

func (f *FinancialInstitutionIdentification18) SetLEI(value string) {
	f.LEI = (*LEIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification18) SetNm(value string) {
	f.Nm = (*Max140Text)(&value)
}

func (f *FinancialInstitutionIdentification18) AddOthr(value string) *GenericFinancialIdentification1 {
	f.Othr = new(GenericFinancialIdentification1)
	return f.Othr
}

type BranchData3 struct {
	Id      *Max35Text       `xml:"Id,omitempty" json:"Id,omitempty"`
	LEI     *LEIIdentifier   `xml:"LEI,omitempty" json:"LEI,omitempty"`
	Nm      *Max140Text      `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr *PostalAddress24 `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

func (b *BranchData3) SetId(value string) {
	b.Id = (*Max35Text)(&value)
}

func (b *BranchData3) SetLEI(value string) {
	b.LEI = (*LEIIdentifier)(&value)
}

func (b *BranchData3) SetNm(value string) {
	b.Nm = (*Max140Text)(&value)
}

func (b *BranchData3) AddPstlAdr(value string) *PostalAddress24 {
	b.PstlAdr = new(PostalAddress24)
	return b.PstlAdr
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp,omitempty" json:"AdrTp,omitempty"`
	Dept        *Max70Text          `xml:"Dept,omitempty" json:"Dept,omitempty"`
	SubDept     *Max70Text          `xml:"SubDept,omitempty" json:"SubDept,omitempty"`
	StrtNm      *Max70Text          `xml:"StrtNm,omitempty" json:"StrtNm,omitempty"`
	BldgNb      *Max16Text          `xml:"BldgNb,omitempty" json:"BldgNb,omitempty"`
	BldgNm      *Max35Text          `xml:"BldgNm,omitempty" json:"BldgNm,omitempty"`
	Flr         *Max70Text          `xml:"Flr,omitempty" json:"Flr,omitempty"`
	PstBx       *Max16Text          `xml:"PstBx,omitempty" json:"PstBx,omitempty"`
	Room        *Max70Text          `xml:"Room,omitempty" json:"Room,omitempty"`
	PstCd       *Max16Text          `xml:"PstCd,omitempty" json:"PstCd,omitempty"`
	TwnNm       *Max35Text          `xml:"TwnNm,omitempty" json:"TwnNm,omitempty"`
	TwnLctnNm   *Max35Text          `xml:"TwnLctnNm,omitempty" json:"TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text          `xml:"DstrctNm,omitempty" json:"DstrctNm,omitempty"`
	CtrySubDvsn *Max35Text          `xml:"CtrySubDvsn,omitempty" json:"CtrySubDvsn,omitempty"`
	Ctry        *CountryCode        `xml:"Ctry,omitempty" json:"Ctry,omitempty"`
	AdrLine     []*Max70Text        `xml:"AdrLine,omitempty" json:"AdrLine,omitempty"`
}

func (p *PostalAddress24) AddAdrTp() *AddressType3Choice {
	p.AdrTp = new(AddressType3Choice)
	return p.AdrTp
}

func (p *PostalAddress24) SetDept(value string) {
	p.Dept = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetSubDept(value string) {
	p.SubDept = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetStrtNm(value string) {
	p.StrtNm = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetBldgNb(value string) {
	p.BldgNb = (*Max16Text)(&value)
}

func (p *PostalAddress24) SetBldgNm(value string) {
	p.BldgNm = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetFlr(value string) {
	p.StrtNm = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetPstBx(value string) {
	p.PstBx = (*Max16Text)(&value)
}

func (p *PostalAddress24) SetRoom(value string) {
	p.Room = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetPstCd(value string) {
	p.PstCd = (*Max16Text)(&value)
}

func (p *PostalAddress24) SetTwnNm(value string) {
	p.TwnNm = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetTwnLctnNm(value string) {
	p.TwnLctnNm = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetDstrctNm(value string) {
	p.DstrctNm = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetCtrySubDvsn(value string) {
	p.CtrySubDvsn = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetCtry(value string) {
	p.Ctry = (*CountryCode)(&value)
}

func (p *PostalAddress24) AddAdrLine(value string) {
	p.AdrLine = append(p.AdrLine, (*Max70Text)(&value))
}

type AddressType3Choice struct {
	Cd    *AddressType2Code        `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry *GenericIdentification30 `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

func (a *AddressType3Choice) AddCd() *AddressType2Code {
	a.Cd = new(AddressType2Code)
	return a.Cd
}

func (a *AddressType3Choice) AddPrtry() *GenericIdentification30 {
	a.Prtry = new(GenericIdentification30)
	return a.Prtry
}

type OriginalGroupHeader17 struct {
	OrgnlMsgId    *Max35Text                        `xml:"OrgnlMsgId" json:"OrgnlMsgId"`
	OrgnlMsgNmId  *Max35Text                        `xml:"OrgnlMsgNmId" json:"OrgnlMsgNmId"`
	OrgnlCreDtTm  *ISODateTime                      `xml:"OrgnlCreDtTm,omitempty" json:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  *Max15NumericText                 `xml:"OrgnlNbOfTxs,omitempty" json:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  *DecimalNumber                    `xml:"OrgnlCtrlSum,omitempty" json:"OrgnlCtrlSum,omitempty"`
	GrpSts        *ExternalPaymentGroupStatus1Code  `xml:"GrpSts,omitempty" json:"GrpSts,omitempty"`
	StsRsnInf     []*StatusReasonInformation12      `xml:"StsRsnInf,omitempty" json:"StsRsnInf,omitempty"`
	NbOfTxsPerSts []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty" json:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupHeader17) SetOrgnlMsgId(value string) {
	o.OrgnlMsgId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader17) SetOrgnlMsgNmId(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader17) SetOrgnlCreDtTm(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader17) SetOrgnlNbOfTxs(value string) {
	o.OrgnlNbOfTxs = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader17) SetOrgnlCtrlSum(value string) {
	o.OrgnlCtrlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader17) SetGrpSts(value string) {
	o.GrpSts = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalGroupHeader17) AddStsRsnInf() *StatusReasonInformation12 {
	newValue := new(StatusReasonInformation12)
	o.StsRsnInf = append(o.StsRsnInf, newValue)
	return newValue
}

func (o *OriginalGroupHeader17) AddNbOfTxsPerSts() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NbOfTxsPerSts = append(o.NbOfTxsPerSts, newValue)
	return newValue
}

// May be no more than 4 items long
type ExternalPaymentGroupStatus1Code string

type StatusReasonInformation12 struct {
	Orgtr    PartyIdentification135 `xml:"Orgtr,omitempty" json:"Orgtr,omitempty"`
	Rsn      StatusReason6Choice    `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
	AddtlInf []Max105Text           `xml:"AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"Id,omitempty" json:"Id,omitempty"`
	CtryOfRes CountryCode     `xml:"CtryOfRes,omitempty" json:"CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"CtctDtls,omitempty" json:"CtctDtls,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId,omitempty" json:"OrgId,omitempty"`
	PrvtId PersonIdentification13       `xml:"PrvtId,omitempty" json:"PrvtId,omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC AnyBICIdentifier                     `xml:"AnyBIC,omitempty" json:"AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"LEI,omitempty" json:"LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"DtAndPlcOfBirth,omitempty" json:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"BirthDt" json:"BirthDt"`
	PrvcOfBirth Max35Text   `xml:"PrvcOfBirth,omitempty" json:"PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"CityOfBirth" json:"CityOfBirth"`
	CtryOfBirth CountryCode `xml:"CtryOfBirth" json:"CtryOfBirth"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"NmPrfx,omitempty" json:"NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"PhneNb,omitempty" json:"PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"MobNb,omitempty" json:"MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"FaxNb,omitempty" json:"FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"EmailAdr,omitempty" json:"EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"EmailPurp,omitempty" json:"EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"JobTitl,omitempty" json:"JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"Rspnsblty,omitempty" json:"Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"Dept,omitempty" json:"Dept,omitempty"`
	Othr      []OtherContact1             `xml:"Othr,omitempty" json:"Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:"PrefrdMtd,omitempty"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs Max15NumericText                      `xml:"DtldNbOfTxs" json:"DtldNbOfTxs"`
	DtldSts     ExternalPaymentTransactionStatus1Code `xml:"DtldSts" json:"DtldSts"`
	DtldCtrlSum DecimalNumber                         `xml:"DtldCtrlSum,omitempty" json:"DtldCtrlSum,omitempty"`
}

// May be no more than 4 items long
type ExternalPaymentTransactionStatus1Code string

type PaymentTransaction110 struct {
	StsId             Max35Text                                    `xml:"StsId,omitempty" json:"StsId,omitempty"`
	OrgnlGrpInf       OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId      Max35Text                                    `xml:"OrgnlInstrId,omitempty" json:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId   Max35Text                                    `xml:"OrgnlEndToEndId,omitempty" json:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId         Max35Text                                    `xml:"OrgnlTxId,omitempty" json:"OrgnlTxId,omitempty"`
	OrgnlUETR         UUIDv4Identifier                             `xml:"OrgnlUETR,omitempty" json:"OrgnlUETR,omitempty"`
	TxSts             ExternalPaymentTransactionStatus1Code        `xml:"TxSts,omitempty" json:"TxSts,omitempty"`
	StsRsnInf         []StatusReasonInformation12                  `xml:"StsRsnInf,omitempty" json:"StsRsnInf,omitempty"`
	ChrgsInf          []Charges7                                   `xml:"ChrgsInf,omitempty" json:"ChrgsInf,omitempty"`
	AccptncDtTm       ISODateTime                                  `xml:"AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	FctvIntrBkSttlmDt DateAndDateTime2Choice                       `xml:"FctvIntrBkSttlmDt,omitempty" json:"FctvIntrBkSttlmDt,omitempty"`
	AcctSvcrRef       Max35Text                                    `xml:"AcctSvcrRef,omitempty" json:"AcctSvcrRef,omitempty"`
	ClrSysRef         Max35Text                                    `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
	InstgAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	OrgnlTxRef        OriginalTransactionReference28               `xml:"OrgnlTxRef,omitempty" json:"OrgnlTxRef,omitempty"`
	SplmtryData       []BI_SupplementaryData1                      `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type OriginalGroupInformation29 struct {
	OrgnlMsgId   Max35Text   `xml:"OrgnlMsgId" json:"OrgnlMsgId"`
	OrgnlMsgNmId Max35Text   `xml:"OrgnlMsgNmId" json:"OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"OrgnlCreDtTm,omitempty"`
}

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt" json:"Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt" json:"Agt"`
}

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"Dt,omitempty" json:"Dt,omitempty"`
	DtTm ISODateTime `xml:"DtTm,omitempty" json:"DtTm,omitempty"`
}

type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:"IntrBkSttlmAmt,omitempty"`
	Amt            AmountType4Choice                            `xml:"Amt,omitempty" json:"Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                      `xml:"ReqdColltnDt,omitempty" json:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:"CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:"SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:"PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedInformation14                  `xml:"MndtRltdInf,omitempty" json:"MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UltmtDbtr      Party40Choice                                `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr           Party40Choice                                `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	DbtrAcct       CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr           Party40Choice                                `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	CdtrAcct       CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr      Party40Choice                                `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	Purp           Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
}

type SettlementInstruction7 struct {
	SttlmMtd             *SettlementMethod1Code                        `xml:"SttlmMtd" json:"SttlmMtd"`
	SttlmAcct            *CashAccount38                                `xml:"SttlmAcct,omitempty" json:"SttlmAcct,omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty" json:"ClrSys,omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:"InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:"InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:"InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:"InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty" json:"ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInstruction7) SetSttlmMtd(value string) {
	s.SttlmMtd = (*SettlementMethod1Code)(&value)
}

// SttlmAcct            *CashAccount38
// ClrSys               *ClearingSystemIdentification3Choice
// InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6
// InstgRmbrsmntAgtAcct *CashAccount38
// InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6
// InstdRmbrsmntAgtAcct *CashAccount38
// ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6
// ThrdRmbrsmntAgtAcct  *CashAccount38

type CashAccount38 struct {
	Id   *AccountIdentification4Choice `xml:"Id" json:"Id"`
	Tp   *CashAccountType2Choice       `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Ccy  *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:"Ccy,omitempty"`
	Nm   *Max70Text                    `xml:"Nm,omitempty" json:"Nm,omitempty"`
	Prxy *ProxyAccountIdentification1  `xml:"Prxy,omitempty" json:"Prxy,omitempty"`
}

// Id   *AccountIdentification4Choice
// Tp   *CashAccountType2Choice
// Ccy  *ActiveOrHistoricCurrencyCode
// Nm   *Max70Text
// Prxy *ProxyAccountIdentification1
type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Id *Max2048Text             `xml:"Id" json:"Id"`
}

func (p *ProxyAccountIdentification1) AddTp() *ProxyAccountType1Choice {
	p.Tp = new(ProxyAccountType1Choice)
	return p.Tp
}

func (p *ProxyAccountIdentification1) SetId(value string) {
	p.Id = (*Max2048Text)(&value)
}

type ProxyAccountType1Choice struct {
	Cd    *ExternalProxyAccountType1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                     `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

func (p *ProxyAccountType1Choice) SetCd(value string) {
	p.Cd = (*ExternalProxyAccountType1Code)(&value)
}

func (p *ProxyAccountType1Choice) SetPrtry(value string) {
	p.Prtry = (*Max35Text)(&value)
}

// May be no more than 4 items long
type ExternalProxyAccountType1Code string

type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code          `xml:"InstrPrty,omitempty" json:"InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:"ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	SeqTp     SequenceType3Code      `xml:"SeqTp,omitempty" json:"SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
}

type MandateRelatedInformation14 struct {
	MndtId        Max35Text                     `xml:"MndtId,omitempty" json:"MndtId,omitempty"`
	DtOfSgntr     ISODate                       `xml:"DtOfSgntr,omitempty" json:"DtOfSgntr,omitempty"`
	AmdmntInd     TrueFalseIndicator            `xml:"AmdmntInd,omitempty" json:"AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty" json:"AmdmntInfDtls,omitempty"`
	ElctrncSgntr  Max1025Text                   `xml:"ElctrncSgntr,omitempty" json:"ElctrncSgntr,omitempty"`
	FrstColltnDt  ISODate                       `xml:"FrstColltnDt,omitempty" json:"FrstColltnDt,omitempty"`
	FnlColltnDt   ISODate                       `xml:"FnlColltnDt,omitempty" json:"FnlColltnDt,omitempty"`
	Frqcy         Frequency36Choice             `xml:"Frqcy,omitempty" json:"Frqcy,omitempty"`
	Rsn           MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
	TrckgDays     Exact2NumericText             `xml:"TrckgDays,omitempty" json:"TrckgDays,omitempty"`
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      Max35Text                                    `xml:"OrgnlMndtId,omitempty" json:"OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty" json:"OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty" json:"OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:"OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification135                       `xml:"OrgnlDbtr,omitempty" json:"OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty" json:"OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty" json:"OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:"OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt ISODate                                      `xml:"OrgnlFnlColltnDt,omitempty" json:"OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty" json:"OrgnlFrqcy,omitempty"`
	OrgnlRsn         MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:"OrgnlRsn,omitempty"`
	OrgnlTrckgDays   Exact2NumericText                            `xml:"OrgnlTrckgDays,omitempty" json:"OrgnlTrckgDays,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Prd    FrequencyPeriod1    `xml:"Prd,omitempty" json:"Prd,omitempty"`
	PtInTm FrequencyAndMoment1 `xml:"PtInTm,omitempty" json:"PtInTm,omitempty"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"Tp" json:"Tp"`
	PtInTm Exact2NumericText `xml:"PtInTm" json:"PtInTm"`
}

// Must match the pattern [0-9]{2}
type Exact2NumericText string

type RemittanceInformation16 struct {
	Ustrd []Max140Text                        `xml:"Ustrd,omitempty" json:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:"Strd,omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:"RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"RfrdDocAmt,omitempty" json:"RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"CdtrRefInf,omitempty" json:"CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"Invcr,omitempty" json:"Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"Invcee,omitempty" json:"Invcee,omitempty"`
	TaxRmt      TaxInformation7                `xml:"TaxRmt,omitempty" json:"TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"GrnshmtRmt,omitempty" json:"GrnshmtRmt,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"AddtlRmtInf,omitempty" json:"AddtlRmtInf,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	AdmstnZone      Max35Text                         `xml:"AdmstnZone,omitempty" json:"AdmstnZone,omitempty"`
	RefNb           Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"Mtd,omitempty" json:"Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	SeqNb           DecimalNumber                     `xml:"SeqNb,omitempty" json:"SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"Rcrd,omitempty" json:"Rcrd,omitempty"`
}

type TaxRecord2 struct {
	Tp       Max35Text  `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Ctgy     Max35Text  `xml:"Ctgy,omitempty" json:"Ctgy,omitempty"`
	CtgyDtls Max35Text  `xml:"CtgyDtls,omitempty" json:"CtgyDtls,omitempty"`
	DbtrSts  Max35Text  `xml:"DbtrSts,omitempty" json:"DbtrSts,omitempty"`
	CertId   Max35Text  `xml:"CertId,omitempty" json:"CertId,omitempty"`
	FrmsCd   Max35Text  `xml:"FrmsCd,omitempty" json:"FrmsCd,omitempty"`
	Prd      TaxPeriod2 `xml:"Prd,omitempty" json:"Prd,omitempty"`
	TaxAmt   TaxAmount2 `xml:"TaxAmt,omitempty" json:"TaxAmt,omitempty"`
	AddtlInf Max140Text `xml:"AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

type TaxPeriod2 struct {
	Yr     ISODate              `xml:"Yr,omitempty" json:"Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"Tp,omitempty" json:"Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"FrToDt,omitempty" json:"FrToDt,omitempty"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"FrDt" json:"FrDt"`
	ToDt ISODate `xml:"ToDt" json:"ToDt"`
}

type TaxAmount2 struct {
	Rate         DecimalNumber                     `xml:"Rate,omitempty" json:"Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:"TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:"TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"Dtls,omitempty" json:"Dtls,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"Prd,omitempty" json:"Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"Amt"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"Tp" json:"Tp"`
	Grnshee           PartyIdentification135            `xml:"Grnshee,omitempty" json:"Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:"GrnshmtAdmstr,omitempty"`
	RefNb             Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Dt                ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"RmtdAmt,omitempty"`
	FmlyMdclInsrncInd TrueFalseIndicator                `xml:"FmlyMdclInsrncInd,omitempty" json:"FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  TrueFalseIndicator                `xml:"MplyeeTermntnInd,omitempty" json:"MplyeeTermntnInd,omitempty"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty,omitempty" json:"Pty,omitempty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty" json:"Agt,omitempty"`
}

type BI_SupplementaryData1 struct {
	PlcAndNm Max350Text                    `xml:"PlcAndNm,omitempty" json:"PlcAndNm,omitempty"`
	Envlp    BI_SupplementaryDataEnvelope1 `xml:"Envlp" json:"Envlp"`
}

type BI_SupplementaryDataEnvelope1 struct {
	Dbtr        BI_AddtlCstmrInf `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	Cdtr        BI_AddtlCstmrInf `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	DbtrAgtAcct CashAccount38    `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgtAcct CashAccount38    `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cstmr       BI_AddtlCstmrInf `xml:"Cstmr,omitempty" json:"Cstmr,omitempty"`
}

type BI_AddtlCstmrInf struct {
	Tp       Max35Text `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Id       Max35Text `xml:"Id,omitempty" json:"Id,omitempty"`
	RsdntSts Max35Text `xml:"RsdntSts,omitempty" json:"RsdntSts,omitempty"`
	TwnNm    Max35Text `xml:"TwnNm,omitempty" json:"TwnNm,omitempty"`
}

// TODO: pacs.008.001.08, json tag

type GroupHeader93 struct {
	MsgId             *Max35Text                                    `xml:"MsgId" json:"MsgId"`
	CreDtTm           *ISODateTime                                  `xml:"CreDtTm" json:"CreDtTm"`
	BtchBookg         *TrueFalseIndicator                           `xml:"BtchBookg,omitempty" json:"BtchBookg,omitempty"`
	NbOfTxs           *Max15NumericText                             `xml:"NbOfTxs" json:"NbOfTxs"`
	CtrlSum           *DecimalNumber                                `xml:"CtrlSum,omitempty" json:"CtrlSum,omitempty"`
	TtlIntrBkSttlmAmt *ActiveCurrencyAndAmount                      `xml:"TtlIntrBkSttlmAmt,omitempty" json:"TtlIntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt     *ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmInf          *SettlementInstruction7                       `xml:"SttlmInf" json:"SttlmInf"`
	PmtTpInf          *PaymentTypeInformation28                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
}

// res.AddMessage(). AddGroupHeader().SetMsgId("val")
func (g *GroupHeader93) SetMsgId(value string) {
	g.MsgId = (*Max35Text)(&value)
}

func (g *GroupHeader93) SetCreDtTm(value string) {
	g.CreDtTm = (*ISODateTime)(&value)
}

func (g *GroupHeader93) SetBtchBookg(value string) {
	g.BtchBookg = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader93) SetNbOfTxs(value string) {
	g.NbOfTxs = (*Max15NumericText)(&value)
}

func (g *GroupHeader93) SetCtrlSum(value string) {
	g.CtrlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader93) AddTtlIntrBkSttlmAmt() *ActiveCurrencyAndAmount {
	g.TtlIntrBkSttlmAmt = new(ActiveCurrencyAndAmount)
	return g.TtlIntrBkSttlmAmt
}

func (g *GroupHeader93) SetIntrBkSttlmDt(value string) {
	g.IntrBkSttlmDt = (*ISODate)(&value)
}

func (g *GroupHeader93) AddSttlmInf() *SettlementInstruction7 {
	g.SttlmInf = new(SettlementInstruction7)
	return g.SttlmInf
}

func (g *GroupHeader93) AddPmtTpInf() *PaymentTypeInformation28 {
	g.PmtTpInf = new(PaymentTypeInformation28)
	return g.PmtTpInf
}

func (g *GroupHeader93) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	g.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstgAgt
}

func (g *GroupHeader93) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	g.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstdAgt
}

type PaymentTypeInformation28 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:"InstrPrty,omitempty"`
	ClrChanl  *ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:"ClrChanl,omitempty"`
	SvcLvl    []*ServiceLevel8Choice  `xml:"SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation28) AddCtgyPurp() *CategoryPurpose1Choice {
	p.CtgyPurp = new(CategoryPurpose1Choice)
	return p.CtgyPurp
}

func (p *PaymentTypeInformation28) AddInstrPrty() *Priority2Code {
	p.InstrPrty = new(Priority2Code)
	return p.InstrPrty
}

func (p *PaymentTypeInformation28) SetClrChanl(value string) {
	p.ClrChanl = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation28) AddSvcLvl() *ServiceLevel8Choice {
	newValue := new(ServiceLevel8Choice)
	p.SvcLvl = append(p.SvcLvl, newValue)
	return newValue
}

func (p *PaymentTypeInformation28) AddLclInstrm() *LocalInstrument2Choice {
	p.LclInstrm = new(LocalInstrument2Choice)
	return p.LclInstrm
}

func (p *PaymentTypeInformation28) AddCtgyPurp() *CategoryPurpose1Choice {
	p.CtgyPurp = new(CategoryPurpose1Choice)
	return p.CtgyPurp
}

type CreditTransferTransaction39 struct {
	PmtId             *PaymentIdentification7                       `xml:"PmtId" json:"PmtId"`
	PmtTpInf          *PaymentTypeInformation28                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	IntrBkSttlmAmt    *ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt" json:"IntrBkSttlmAmt"`
	IntrBkSttlmDt     *ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmPrty         *Priority3Code                                `xml:"SttlmPrty,omitempty" json:"SttlmPrty,omitempty"`
	SttlmTmIndctn     *SettlementDateTimeIndication1                `xml:"SttlmTmIndctn,omitempty" json:"SttlmTmIndctn,omitempty"`
	SttlmTmReq        *SettlementTimeRequest2                       `xml:"SttlmTmReq,omitempty" json:"SttlmTmReq,omitempty"`
	AccptncDtTm       *ISODateTime                                  `xml:"AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	PoolgAdjstmntDt   *ISODate                                      `xml:"PoolgAdjstmntDt,omitempty" json:"PoolgAdjstmntDt,omitempty"`
	InstdAmt          *ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty" json:"InstdAmt,omitempty"`
	XchgRate          *DecimalNumber                                `xml:"XchgRate,omitempty" json:"XchgRate,omitempty"`
	ChrgBr            *ChargeBearerType1Code                        `xml:"ChrgBr" json:"ChrgBr"`
	ChrgsInf          []*Charges7                                   `xml:"ChrgsInf,omitempty" json:"ChrgsInf,omitempty"`
	PrvsInstgAgt1     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct *CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct *CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct *CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	IntrmyAgt1        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct    *CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct    *CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct    *CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	UltmtDbtr         *PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	InitgPty          *PartyIdentification135                       `xml:"InitgPty,omitempty" json:"InitgPty,omitempty"`
	Dbtr              *PartyIdentification135                       `xml:"Dbtr" json:"Dbtr"`
	DbtrAcct          *CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt" json:"DbtrAgt"`
	DbtrAgtAcct       *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt" json:"CdtrAgt"`
	CdtrAgtAcct       *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr              *PartyIdentification135                       `xml:"Cdtr" json:"Cdtr"`
	CdtrAcct          *CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr         *PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt   []*InstructionForCreditorAgent1               `xml:"InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt    []*InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Purp              *Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
	RgltryRptg        []*RegulatoryReporting3                       `xml:"RgltryRptg,omitempty" json:"RgltryRptg,omitempty"`
	Tax               *TaxInformation8                              `xml:"Tax,omitempty" json:"Tax,omitempty"`
	RltdRmtInf        []*RemittanceLocation7                        `xml:"RltdRmtInf,omitempty" json:"RltdRmtInf,omitempty"`
	RmtInf            *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	SplmtryData       []*BI_SupplementaryData1                      `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

// res.AddMessage(). AddCreditTransferTransactionInformation().SetChrgBr("val")
func (c *CreditTransferTransaction39) SetChrgBr(value string) {
	c.ChrgBr = (*ChargeBearerType1Code)(&value)
}

type PaymentIdentification7 struct {
	InstrId    Max35Text        `xml:"InstrId,omitempty" json:"InstrId,omitempty"`
	EndToEndId Max35Text        `xml:"EndToEndId" json:"EndToEndId"`
	TxId       Max35Text        `xml:"TxId,omitempty" json:"TxId,omitempty"`
	UETR       UUIDv4Identifier `xml:"UETR,omitempty" json:"UETR,omitempty"`
	ClrSysRef  Max35Text        `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
}

type TaxInformation8 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	AdmstnZone      Max35Text                         `xml:"AdmstnZone,omitempty" json:"AdmstnZone,omitempty"`
	RefNb           Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"Mtd,omitempty" json:"Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	SeqNb           DecimalNumber                     `xml:"SeqNb,omitempty" json:"SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"Rcrd,omitempty" json:"Rcrd,omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       Max35Text                 `xml:"RmtId,omitempty" json:"RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty" json:"RmtLctnDtls,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd" json:"Mtd"`
	ElctrncAdr Max2048Text                   `xml:"ElctrncAdr,omitempty" json:"ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

type NameAndAddress16 struct {
	Nm  Max140Text      `xml:"Nm" json:"Nm"`
	Adr PostalAddress24 `xml:"Adr" json:"Adr"`
}

// TODO: pacs.009.001.09, json tag

type CreditTransferTransaction44 struct {
	PmtId              PaymentIdentification13                      `xml:"PmtId" json:"PmtId"`
	PmtTpInf           PaymentTypeInformation28                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	IntrBkSttlmAmt     ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt" json:"IntrBkSttlmAmt"`
	IntrBkSttlmDt      ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmPrty          Priority3Code                                `xml:"SttlmPrty,omitempty" json:"SttlmPrty,omitempty"`
	SttlmTmIndctn      SettlementDateTimeIndication1                `xml:"SttlmTmIndctn,omitempty" json:"SttlmTmIndctn,omitempty"`
	SttlmTmReq         SettlementTimeRequest2                       `xml:"SttlmTmReq,omitempty" json:"SttlmTmReq,omitempty"`
	PrvsInstgAgt1      BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct  CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2      BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct  CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3      BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct  CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	InstgAgt           BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt           BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	IntrmyAgt1         BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct     CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2         BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct     CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3         BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct     CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	UltmtDbtr          BranchAndFinancialInstitutionIdentification6 `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr               BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr" json:"Dbtr"`
	DbtrAcct           CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt            BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct        CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt            BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct        CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr               BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr" json:"Cdtr"`
	CdtrAcct           CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr          BranchAndFinancialInstitutionIdentification6 `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt    []InstructionForCreditorAgent3               `xml:"InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt     []InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Purp               Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
	RmtInf             RemittanceInformation2                       `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UndrlygCstmrCdtTrf CreditTransferTransaction45                  `xml:"UndrlygCstmrCdtTrf,omitempty" json:"UndrlygCstmrCdtTrf,omitempty"`
	SplmtryData        []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

func (f *CreditTransferTransaction44) AddPmtId() *PaymentIdentification13 {
	f.PmtId = new(PaymentIdentification13)
	return f.PmtId
}

func (f *CreditTransferTransaction44) AddPmtTpInf() *PaymentTypeInformation28 {
	f.PmtTpInf = new(PaymentTypeInformation28)
	return f.PmtTpInf
}

func (f *CreditTransferTransaction44) AddIntrBkSttlmAmt() *ActiveCurrencyAndAmount {
	f.IntrBkSttlmAmt = new(ActiveCurrencyAndAmount)
	return f.IntrBkSttlmAmt
}

func (f *CreditTransferTransaction44) SetIntrBkSttlmDt(value string) {
	f.IntrBkSttlmDt = (*ISODate)(&value)
}

func (f *CreditTransferTransaction44) AddDbtr() *BranchAndFinancialInstitutionIdentification6 {
	f.Dbtr = new(BranchAndFinancialInstitutionIdentification6)
	return f.Dbtr
}

func (f *CreditTransferTransaction44) AddCdtr() *BranchAndFinancialInstitutionIdentification6 {
	f.Cdtr = new(BranchAndFinancialInstitutionIdentification6)
	return f.Cdtr
}

func (f *CreditTransferTransaction44) AddRmtInf() *RemittanceInformation2 {
	f.RmtInf = new(RemittanceInformation2)
	return f.RmtInf
}

type PaymentIdentification13 struct {
	InstrId    Max35Text        `xml:"InstrId,omitempty" json:"InstrId,omitempty"`
	EndToEndId Max35Text        `xml:"EndToEndId" json:"EndToEndId"`
	TxId       Max35Text        `xml:"TxId,omitempty" json:"TxId,omitempty"`
	UETR       UUIDv4Identifier `xml:"UETR,omitempty" json:"UETR,omitempty"`
	ClrSysRef  Max35Text        `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
}

func (p *PaymentIdentification13) SetEndToEndId(value string) {
	p.EndToEndId = (*Max35Text)(&value)
}

func (p *PaymentIdentification13) SetTxId(value string) {
	p.TxId = (*Max35Text)(&value)
}

type InstructionForCreditorAgent3 struct {
	Cd       ExternalCreditorAgentInstruction1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	InstrInf Max140Text                            `xml:"InstrInf,omitempty" json:"InstrInf,omitempty"`
}

// May be no more than 4 items long
type ExternalCreditorAgentInstruction1Code string

type CreditTransferTransaction45 struct {
	UltmtDbtr         PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	InitgPty          PartyIdentification135                       `xml:"InitgPty,omitempty" json:"InitgPty,omitempty"`
	Dbtr              PartyIdentification135                       `xml:"Dbtr" json:"Dbtr"`
	DbtrAcct          CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt           BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt" json:"DbtrAgt"`
	DbtrAgtAcct       CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	PrvsInstgAgt1     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	IntrmyAgt1        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct    CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct    CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct    CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	CdtrAgt           BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt" json:"CdtrAgt"`
	CdtrAgtAcct       CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr              PartyIdentification135                       `xml:"Cdtr" json:"Cdtr"`
	CdtrAcct          CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr         PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt   []InstructionForCreditorAgent3               `xml:"InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt    []InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Tax               TaxInformation8                              `xml:"Tax,omitempty" json:"Tax,omitempty"`
	RmtInf            RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	InstdAmt          ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty" json:"InstdAmt,omitempty"`
}

// TODO: pacs.028.001.04, json tag

type OriginalGroupInformation27 struct {
	OrgnlMsgId   Max35Text        `xml:"OrgnlMsgId" json:"OrgnlMsgId"`
	OrgnlMsgNmId Max35Text        `xml:"OrgnlMsgNmId" json:"OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime      `xml:"OrgnlCreDtTm,omitempty" json:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs Max15NumericText `xml:"OrgnlNbOfTxs,omitempty" json:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum DecimalNumber    `xml:"OrgnlCtrlSum,omitempty" json:"OrgnlCtrlSum,omitempty"`
}

type PaymentTransaction121 struct {
	StsReqId        Max35Text                                    `xml:"StsReqId,omitempty" json:"StsReqId,omitempty"`
	OrgnlGrpInf     OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId    Max35Text                                    `xml:"OrgnlInstrId,omitempty" json:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId Max35Text                                    `xml:"OrgnlEndToEndId,omitempty" json:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId       Max35Text                                    `xml:"OrgnlTxId,omitempty" json:"OrgnlTxId,omitempty"`
	OrgnlUETR       UUIDv4Identifier                             `xml:"OrgnlUETR,omitempty" json:"OrgnlUETR,omitempty"`
	AccptncDtTm     ISODateTime                                  `xml:"AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	ClrSysRef       Max35Text                                    `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
	InstgAgt        BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt        BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	OrgnlTxRef      OriginalTransactionReference31               `xml:"OrgnlTxRef,omitempty" json:"OrgnlTxRef,omitempty"`
	SplmtryData     []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

//func (g *GroupHeader93) SetCtrlSum (value string) {
//	g.CtrlSum = (*DecimalNumber)(&value)
//}
//
//func (g *GroupHeader93) AddTtlIntrBkSttlmAmt() *ActiveCurrencyAndAmount {
//	g.TtlIntrBkSttlmAmt = new(ActiveCurrencyAndAmount)
//	return g.TtlIntrBkSttlmAmt
//}

func (p *PaymentTransaction121) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:"IntrBkSttlmAmt,omitempty"`
	Amt            AmountType4Choice                            `xml:"Amt,omitempty" json:"Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                      `xml:"ReqdColltnDt,omitempty" json:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:"CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:"SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:"PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedData1Choice                    `xml:"MndtRltdInf,omitempty" json:"MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UltmtDbtr      Party40Choice                                `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr           Party40Choice                                `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	DbtrAcct       CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr           Party40Choice                                `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	CdtrAcct       CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr      Party40Choice                                `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	Purp           Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
}

type MandateRelatedData1Choice struct {
	DrctDbtMndt MandateRelatedInformation14 `xml:"DrctDbtMndt,omitempty" json:"DrctDbtMndt,omitempty"`
	CdtTrfMndt  CreditTransferMandateData1  `xml:"CdtTrfMndt,omitempty" json:"CdtTrfMndt,omitempty"`
}

type CreditTransferMandateData1 struct {
	MndtId       Max35Text                 `xml:"MndtId,omitempty" json:"MndtId,omitempty"`
	Tp           MandateTypeInformation2   `xml:"Tp,omitempty" json:"Tp,omitempty"`
	DtOfSgntr    ISODate                   `xml:"DtOfSgntr,omitempty" json:"DtOfSgntr,omitempty"`
	DtOfVrfctn   ISODateTime               `xml:"DtOfVrfctn,omitempty" json:"DtOfVrfctn,omitempty"`
	ElctrncSgntr Max10KBinary              `xml:"ElctrncSgntr,omitempty" json:"ElctrncSgntr,omitempty"`
	FrstPmtDt    ISODate                   `xml:"FrstPmtDt,omitempty" json:"FrstPmtDt,omitempty"`
	FnlPmtDt     ISODate                   `xml:"FnlPmtDt,omitempty" json:"FnlPmtDt,omitempty"`
	Frqcy        Frequency36Choice         `xml:"Frqcy,omitempty" json:"Frqcy,omitempty"`
	Rsn          MandateSetupReason1Choice `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"Clssfctn,omitempty" json:"Clssfctn,omitempty"`
}

type MandateClassification1Choice struct {
	Cd    MandateClassification1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                  `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// May be one of FIXE, USGB, VARI
type MandateClassification1Code string

// TODO: prxy.001.001.01, json tag

type ProxyRegistration1 struct {
	RegnTp    ProxyRegistrationType1Code    `xml:"RegnTp" json:"RegnTp"`
	RegnSubTp ProxyRegistrationSubType1Code `xml:"RegnSubTp,omitempty" json:"RegnSubTp,omitempty"`
	Prxy      ProxyDefinition1              `xml:"Prxy,omitempty" json:"Prxy,omitempty"`
	PrxyRegn  ProxyRegistrationAccount1     `xml:"PrxyRegn" json:"PrxyRegn"`
}

// May be one of AMND, DEAC, NEWR, SUSP, ACTV, PORT
type ProxyRegistrationType1Code string

// May be no more than 4 items long
type ProxyRegistrationSubType1Code string

type ProxyDefinition1 struct {
	Tp  Max12Text  `xml:"Tp" json:"Tp"`
	Val Max140Text `xml:"Val" json:"Val"`
}

type ProxyRegistrationAccount1 struct {
	RegnId      Max35Text                                    `xml:"RegnId,omitempty" json:"RegnId,omitempty"`
	DsplNm      Max140Text                                   `xml:"DsplNm,omitempty" json:"DsplNm,omitempty"`
	Agt         BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty" json:"Agt,omitempty"`
	Acct        CashAccount40                                `xml:"Acct,omitempty" json:"Acct,omitempty"`
	AcctHldr    Party30Choice                                `xml:"AcctHldr,omitempty" json:"AcctHldr,omitempty"`
	ScndId      ScndIdDefinition1                            `xml:"ScndId,omitempty" json:"ScndId,omitempty"`
	RegnSts     ProxyRegistrationStatusCode                  `xml:"RegnSts,omitempty" json:"RegnSts,omitempty"`
	PreAuthrsd  TrueFalseIndicator                           `xml:"PreAuthrsd,omitempty" json:"PreAuthrsd,omitempty"`
	SplmtryData []BI_SupplementaryData1                      `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type CashAccount40 struct {
	Id         AccountIdentification4Choice `xml:"Id" json:"Id"`
	Tp         CashAccountType2ChoiceProxy  `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Nm         Max140Text                   `xml:"Nm,omitempty" json:"Nm,omitempty"`
	AcctHldrTp Max35Text                    `xml:"AcctHldrTp,omitempty" json:"AcctHldrTp,omitempty"`
}

type CashAccountType2ChoiceProxy struct {
	Prtry ProxyAccountType `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// May be no more than 4 items long
type ProxyAccountType string

type ScndIdDefinition1 struct {
	Tp  Max12Text `xml:"Tp" json:"Tp"`
	Val Max35Text `xml:"Val" json:"Val"`
}

// May be no more than 4 items long
type ProxyRegistrationStatusCode string

// TODO: prxy.002.001.01, json tag

type ProxyRegistrationResponse1 struct {
	PrxRspnSts     ProxyStatusCode               `xml:"PrxRspnSts" json:"PrxRspnSts"`
	StsRsnInf      ProxyStatusChoice             `xml:"StsRsnInf,omitempty" json:"StsRsnInf,omitempty"`
	OrgnlRegnTp    ProxyRegistrationType1Code    `xml:"OrgnlRegnTp" json:"OrgnlRegnTp"`
	OrgnlRegnSubTp ProxyRegistrationSubType1Code `xml:"OrgnlRegnSubTp,omitempty" json:"OrgnlRegnSubTp,omitempty"`
	OrgnlPrxy      ProxyDefinition1              `xml:"OrgnlPrxy,omitempty" json:"OrgnlPrxy,omitempty"`
	PrxyRegn       []ProxyRegistrationAccount1   `xml:"PrxyRegn,omitempty" json:"PrxyRegn,omitempty"`
}

// May be one of ACTC, RJCT
type ProxyStatusCode string

type ProxyStatusChoice struct {
	Cd    ExternalStatusReason1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                 `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// TODO: prxy.003.001.01, json tag
type ProxyLookUpChoice1 struct {
	PrxyOnly ProxyLookUp11 `xml:"PrxyOnly" json:"PrxyOnly"`
}

type ProxyLookUp11 struct {
	LkUpTp    ProxyLookUpType1Code `xml:"LkUpTp" json:"LkUpTP"`
	Id        Max35Text            `xml:"Id" json:"Id"`
	PrxyRtrvl ProxyDefinition1     `xml:"PrxyRtrvl" json:"PrxyRtrvl"`
}

type ProxyLookUpType1Code string

// TODO: prxy.004.001.01, json tag
type ProxyLookUpResponse1 struct {
	OrgnlId        Max35Text                `xml:"OrgnlId" json:"OrgnlId"`
	OrgnlPrxyRtrvl ProxyDefinition1         `xml:"OrgnlPrxyRtrvl" json:"OrgnlPrxyRtrvl"`
	RegnRspn       ProxyLookUpRegistration1 `xml:"RegnRspn" json:"RegnRspn"`
}

type ProxyLookUpRegistration1 struct {
	PrxRspnSts ProxyStatusCode     `xml:"PrxRspnSts" json:"PrxRspnSts"`
	StsRsnInf  ProxyStatusChoice   `xml:"StsRsnInf" json:"StsRsnInf"`
	Prxy       ProxyDefinition1    `xml:"Prxy" json:"Prxy"`
	Regn       ProxyLookUpAccount1 `xml:"Regn" json:"Regn"`
}

type ProxyLookUpAccount1 struct {
	RegnId Max35Text                                    `xml:"RegnId" json:"RegnId"`
	DsplNm Max140Text                                   `xml:"DsplNm" json:"DsplNm"`
	Agt    BranchAndFinancialInstitutionIdentification5 `xml:"Agt" json:"Agt"`
	Acct   CashAccount40                                `xml:"Acct" json:"Acct"`
}

// TODO: prxy.005.001.01, json tag
type ProxyEnquiryChoice1 struct {
	RegnId Max35Text         `xml:"RegnId" json:"RegnId"`
	ScndId ScndIdDefinition1 `xml:"ScndId" json:"ScndId"`
}

// TODO: prxy.006.001.01, json tag
type ProxyEnquiryResponse1 struct {
	PrxRspnSts ProxyStatusCode          `xml:"PrxRspnSts" json:"PrxRspnSts"`
	StsRsnInf  ProxyStatusChoice        `xml:"StsRsnInf" json:"StsRsnInf"`
	Rspn       ProxyEnquiryInformation1 `xml:"Rspn" json:"Rspn"`
}

type ProxyEnquiryInformation1 struct {
	RegnId      Max35Text                                    `xml:"RegnId" json:"RegnId"`
	DsplNm      Max140Text                                   `xml:"DsplNm" json:"DsplNm"`
	Ptcpt       BranchAndFinancialInstitutionIdentification5 `xml:"Ptcpt" json:"Ptcpt"`
	PrxyInf     ProxyEnquiryDefinition1                      `xml:"PrxyInf" json:"PrxyInf"`
	AcctInf     ProxyEnquiryAccount1                         `xml:"AcctInf" json:"AcctInf"`
	SplmtryData BI_SupplementaryData1                        `xml:"SplmtryData" json:"SplmtryData"`
}

type ProxyEnquiryDefinition1 struct {
	Tp  Max12Text              `xml:"Tp" json:"Tp"`
	Val Max140Text             `xml:"Val" json:"Val"`
	Sts ProxyEnquiryStatusCode `xml:"Sts" json:"Sts"`
}

type ProxyEnquiryAccount1 struct {
	Agt BranchAndFinancialInstitutionIdentification5 `xml:"Agt" json:"Agt"`
}

type ProxyEnquiryStatusCode string

// TODO: prxy.901.001.01, json tag
type ProxyNtfctn1 struct {
	OrgnlId   Max35Text        `xml:"OrgnlId" json:"OrgnlId"`
	OrgnlPrxy ProxyDefinition1 `xml:"OrgnlPrxy" json:"OrgnlPrxy"`
	OrgnlAcct ProxyAccount1    `xml:"OrgnlAcct" json:"OrgnlAcct"`
	NewAcct   ProxyAccount1    `xml:"NewAcct" json:"NewAcct"`
}

type ProxyAccount1 struct {
	RegnId      Max35Text                                    `xml:"RegnId" json:"RegnId"`
	DsplNm      Max140Text                                   `xml:"DsplNm" json:"DsplNm"`
	Agt         BranchAndFinancialInstitutionIdentification5 `xml:"Agt" json:"Agt"`
	Acct        CashAccount40                                `xml:"Acct" json:"Acct"`
	SplmtryData BI_SupplementaryData1                        `xml:"SplmtryData" json:"SplmtryData"`
}

type AdminTransactionInformation struct {
	FnctnCd  Max4NumericText                              `xml:"FnctnCd" json:"FnctnCd"`
	InstrId  Max35Text                                    `xml:"InstrId" json:"InstrId"`
	InstgAgt BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt" json:"InstgAgt"`
}

type AdminResponse struct {
	InstgAgt     BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt" json:"InstgAgt"`
	OrgnlInstrId Max35Text                                    `xml:"OrgnlInstrId" json:"OrgnlInstrId"`
	FnctnCd      Max4NumericText                              `xml:"FnctnCd" json:"FnctnCd"`
	TxSts        TransactionIndividualStatus3Code             `xml:"TxSts" json:"TxSts"`
}
