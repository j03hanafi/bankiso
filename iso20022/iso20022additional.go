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

func (f *FinancialInstitutionIdentification18) AddOthr() *GenericFinancialIdentification1 {
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
	Orgtr    *PartyIdentification135 `xml:"Orgtr,omitempty" json:"Orgtr,omitempty"`
	Rsn      *StatusReason6Choice    `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
	AddtlInf []*Max105Text           `xml:"AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

func (s *StatusReasonInformation12) AddOrgtr() *PartyIdentification135 {
	s.Orgtr = new(PartyIdentification135)
	return s.Orgtr
}

func (s *StatusReasonInformation12) AddRsn() *StatusReason6Choice {
	s.Rsn = new(StatusReason6Choice)
	return s.Rsn
}

func (s *StatusReasonInformation12) AddAddtlInf(value string) {
	s.AddtlInf = append(s.AddtlInf, (*Max105Text)(&value))
}

type PartyIdentification135 struct {
	Nm        *Max140Text      `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr   *PostalAddress24 `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Id        *Party38Choice   `xml:"Id,omitempty" json:"Id,omitempty"`
	CtryOfRes *CountryCode     `xml:"CtryOfRes,omitempty" json:"CtryOfRes,omitempty"`
	CtctDtls  *Contact4        `xml:"CtctDtls,omitempty" json:"CtctDtls,omitempty"`
}

func (p *PartyIdentification135) SetNm(value string) {
	p.Nm = (*Max140Text)(&value)
}

func (p *PartyIdentification135) AddPstlAdr() *PostalAddress24 {
	p.PstlAdr = new(PostalAddress24)
	return p.PstlAdr
}

func (p *PartyIdentification135) AddId() *Party38Choice {
	p.Id = new(Party38Choice)
	return p.Id
}

func (p *PartyIdentification135) SetCtryOfRes(value string) {
	p.CtryOfRes = (*CountryCode)(&value)
}

func (p *PartyIdentification135) AddCtctDtls() *Contact4 {
	p.CtctDtls = new(Contact4)
	return p.CtctDtls
}

type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"OrgId,omitempty" json:"OrgId,omitempty"`
	PrvtId *PersonIdentification13       `xml:"PrvtId,omitempty" json:"PrvtId,omitempty"`
}

func (p *Party38Choice) AddOrgId() *OrganisationIdentification29 {
	p.OrgId = new(OrganisationIdentification29)
	return p.OrgId
}

func (p *Party38Choice) AddPrvtId() *PersonIdentification13 {
	p.PrvtId = new(PersonIdentification13)
	return p.PrvtId
}

type OrganisationIdentification29 struct {
	AnyBIC *AnyBICIdentifier                     `xml:"AnyBIC,omitempty" json:"AnyBIC,omitempty"`
	LEI    *LEIIdentifier                        `xml:"LEI,omitempty" json:"LEI,omitempty"`
	Othr   []*GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

func (o *OrganisationIdentification29) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification29) SetLEI(value string) {
	o.LEI = (*LEIIdentifier)(&value)
}

func (o *OrganisationIdentification29) AddOthr() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Othr = append(o.Othr, newValue)
	return newValue
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1           `xml:"DtAndPlcOfBirth,omitempty" json:"DtAndPlcOfBirth,omitempty"`
	Othr            []*GenericPersonIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

func (p *PersonIdentification13) AddDtAndPlcOfBirth() *DateAndPlaceOfBirth1 {
	p.DtAndPlcOfBirth = new(DateAndPlaceOfBirth1)
	return p.DtAndPlcOfBirth
}

func (p *PersonIdentification13) AddOthr() *GenericPersonIdentification1 {
	newValue := new(GenericPersonIdentification1)
	p.Othr = append(p.Othr, newValue)
	return newValue
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     *ISODate     `xml:"BirthDt" json:"BirthDt"`
	PrvcOfBirth *Max35Text   `xml:"PrvcOfBirth,omitempty" json:"PrvcOfBirth,omitempty"`
	CityOfBirth *Max35Text   `xml:"CityOfBirth" json:"CityOfBirth"`
	CtryOfBirth *CountryCode `xml:"CtryOfBirth" json:"CtryOfBirth"`
}

func (d *DateAndPlaceOfBirth1) SetBirthDt(value string) {
	d.BirthDt = (*ISODate)(&value)
}

func (d *DateAndPlaceOfBirth1) SetPrvcOfBirth(value string) {
	d.PrvcOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth1) SetCityOfBirth(value string) {
	d.CityOfBirth = (*Max35Text)(&value)
}

type Contact4 struct {
	NmPrfx    *NamePrefix2Code             `xml:"NmPrfx,omitempty" json:"NmPrfx,omitempty"`
	Nm        *Max140Text                  `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PhneNb    *PhoneNumber                 `xml:"PhneNb,omitempty" json:"PhneNb,omitempty"`
	MobNb     *PhoneNumber                 `xml:"MobNb,omitempty" json:"MobNb,omitempty"`
	FaxNb     *PhoneNumber                 `xml:"FaxNb,omitempty" json:"FaxNb,omitempty"`
	EmailAdr  *Max2048Text                 `xml:"EmailAdr,omitempty" json:"EmailAdr,omitempty"`
	EmailPurp *Max35Text                   `xml:"EmailPurp,omitempty" json:"EmailPurp,omitempty"`
	JobTitl   *Max35Text                   `xml:"JobTitl,omitempty" json:"JobTitl,omitempty"`
	Rspnsblty *Max35Text                   `xml:"Rspnsblty,omitempty" json:"Rspnsblty,omitempty"`
	Dept      *Max70Text                   `xml:"Dept,omitempty" json:"Dept,omitempty"`
	Othr      []*OtherContact1             `xml:"Othr,omitempty" json:"Othr,omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:"PrefrdMtd,omitempty"`
}

func (c *Contact4) SetNmPrfx(value string) {
	c.NmPrfx = (*NamePrefix2Code)(&value)
}

func (c *Contact4) SetNm(value string) {
	c.Nm = (*Max140Text)(&value)
}

func (c *Contact4) SetPhneNb(value string) {
	c.PhneNb = (*PhoneNumber)(&value)
}

func (c *Contact4) SetMobNb(value string) {
	c.MobNb = (*PhoneNumber)(&value)
}

func (c *Contact4) SetFaxNb(value string) {
	c.FaxNb = (*PhoneNumber)(&value)
}

func (c *Contact4) SetEmailAdr(value string) {
	c.EmailAdr = (*Max2048Text)(&value)
}

func (c *Contact4) SetJobTitl(value string) {
	c.JobTitl = (*Max35Text)(&value)
}
func (c *Contact4) SetRspnsblty(value string) {
	c.Rspnsblty = (*Max35Text)(&value)
}

func (c *Contact4) SetDept(value string) {
	c.Dept = (*Max70Text)(&value)
}

// Othr      []*OtherContact1
func (c *Contact4) AddOthr() *OtherContact1 {
	newValue := new(OtherContact1)
	c.Othr = append(c.Othr, newValue)
	return newValue
}

func (c *Contact4) SetPrefrdMtd(value string) {
	c.PrefrdMtd = (*PreferredContactMethod1Code)(&value)
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs *Max15NumericText                      `xml:"DtldNbOfTxs" json:"DtldNbOfTxs"`
	DtldSts     *ExternalPaymentTransactionStatus1Code `xml:"DtldSts" json:"DtldSts"`
	DtldCtrlSum *DecimalNumber                         `xml:"DtldCtrlSum,omitempty" json:"DtldCtrlSum,omitempty"`
}

func (n *NumberOfTransactionsPerStatus5) SetDtldNbOfTxs(value string) {
	n.DtldNbOfTxs = (*Max15NumericText)(&value)
}

func (n *NumberOfTransactionsPerStatus5) SetDtldSts(value string) {
	n.DtldSts = (*ExternalPaymentTransactionStatus1Code)(&value)
}

func (n *NumberOfTransactionsPerStatus5) SetDtldCtrlSum(value string) {
	n.DtldCtrlSum = (*DecimalNumber)(&value)
}

// May be no more than 4 items long
type ExternalPaymentTransactionStatus1Code string

type PaymentTransaction110 struct {
	StsId             *Max35Text                                    `xml:"StsId,omitempty" json:"StsId,omitempty"`
	OrgnlGrpInf       *OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId      *Max35Text                                    `xml:"OrgnlInstrId,omitempty" json:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId   *Max35Text                                    `xml:"OrgnlEndToEndId,omitempty" json:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId         *Max35Text                                    `xml:"OrgnlTxId,omitempty" json:"OrgnlTxId,omitempty"`
	OrgnlUETR         *UUIDv4Identifier                             `xml:"OrgnlUETR,omitempty" json:"OrgnlUETR,omitempty"`
	TxSts             *ExternalPaymentTransactionStatus1Code        `xml:"TxSts,omitempty" json:"TxSts,omitempty"`
	StsRsnInf         []*StatusReasonInformation12                  `xml:"StsRsnInf,omitempty" json:"StsRsnInf,omitempty"`
	ChrgsInf          []*Charges7                                   `xml:"ChrgsInf,omitempty" json:"ChrgsInf,omitempty"`
	AccptncDtTm       *ISODateTime                                  `xml:"AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	FctvIntrBkSttlmDt *DateAndDateTime2Choice                       `xml:"FctvIntrBkSttlmDt,omitempty" json:"FctvIntrBkSttlmDt,omitempty"`
	AcctSvcrRef       *Max35Text                                    `xml:"AcctSvcrRef,omitempty" json:"AcctSvcrRef,omitempty"`
	ClrSysRef         *Max35Text                                    `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	OrgnlTxRef        *OriginalTransactionReference28               `xml:"OrgnlTxRef,omitempty" json:"OrgnlTxRef,omitempty"`
	SplmtryData       []*BI_SupplementaryData1                      `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction110) SetStsId(value string) {
	p.StsId = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) AddC() *OriginalGroupInformation29 {
	p.OrgnlGrpInf = new(OriginalGroupInformation29)
	return p.OrgnlGrpInf
}

func (p *PaymentTransaction110) SetOrgnlInstrId(value string) {
	p.OrgnlInstrId = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) SetOrgnlTxId(value string) {
	p.OrgnlTxId = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) SetOrgnlUETR(value string) {
	p.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentTransaction110) SetTxSts(value string) {
	p.TxSts = (*ExternalPaymentTransactionStatus1Code)(&value)
}

func (p *PaymentTransaction110) AddStsRsnInf() *StatusReasonInformation12 {
	newValue := new(StatusReasonInformation12)
	p.StsRsnInf = append(p.StsRsnInf, newValue)
	return newValue
}

func (p *PaymentTransaction110) AddChrgsInf() *Charges7 {
	newValue := new(Charges7)
	p.ChrgsInf = append(p.ChrgsInf, newValue)
	return newValue
}

func (p *PaymentTransaction110) SetAccptncDtTm(value string) {
	p.AccptncDtTm = (*ISODateTime)(&value)
}

func (p *PaymentTransaction110) AddFctvIntrBkSttlmDt() *DateAndDateTime2Choice {
	p.FctvIntrBkSttlmDt = new(DateAndDateTime2Choice)
	return p.FctvIntrBkSttlmDt
}

func (p *PaymentTransaction110) SetAcctSvcrRef(value string) {
	p.AcctSvcrRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) SetClrSysRef(value string) {
	p.ClrSysRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction110) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	p.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return p.InstgAgt
}

func (p *PaymentTransaction110) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	p.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return p.InstdAgt
}

func (p *PaymentTransaction110) AddOrgnlTxRef() *OriginalTransactionReference28 {
	p.OrgnlTxRef = new(OriginalTransactionReference28)
	return p.OrgnlTxRef
}

func (p *PaymentTransaction110) AddSplmtryData() *BI_SupplementaryData1 {
	newValue := new(BI_SupplementaryData1)
	p.SplmtryData = append(p.SplmtryData, newValue)
	return newValue
}

type OriginalGroupInformation29 struct {
	OrgnlMsgId   *Max35Text   `xml:"OrgnlMsgId" json:"OrgnlMsgId"`
	OrgnlMsgNmId *Max35Text   `xml:"OrgnlMsgNmId" json:"OrgnlMsgNmId"`
	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"OrgnlCreDtTm,omitempty"`
}

func (o *OriginalGroupInformation29) SetOrgnlMsgId(value string) {
	o.OrgnlMsgId = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation29) SetOrgnlMsgNmId(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation29) SetOrgnlCreDtTm(value string) {
	o.OrgnlCreDtTm = (*ISODateTime)(&value)
}

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

type Charges7 struct {
	Amt *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt" json:"Amt"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt" json:"Agt"`
}

func (c *Charges7) AddAmt() *ActiveOrHistoricCurrencyAndAmount {
	c.Amt = new(ActiveOrHistoricCurrencyAndAmount)
	return c.Amt
}

func (c *Charges7) AddAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.Agt = new(BranchAndFinancialInstitutionIdentification6)
	return c.Agt
}

type DateAndDateTime2Choice struct {
	Dt   *ISODate     `xml:"Dt,omitempty" json:"Dt,omitempty"`
	DtTm *ISODateTime `xml:"DtTm,omitempty" json:"DtTm,omitempty"`
}

func (d *DateAndDateTime2Choice) SetDt(value string) {
	d.Dt = (*ISODate)(&value)
}

func (d *DateAndDateTime2Choice) SetDtTm(value string) {
	d.DtTm = (*ISODateTime)(&value)
}

type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:"IntrBkSttlmAmt,omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty" json:"Amt,omitempty"`
	IntrBkSttlmDt  *ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   *ISODate                                      `xml:"ReqdColltnDt,omitempty" json:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:"CdtrSchmeId,omitempty"`
	SttlmInf       *SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:"SttlmInf,omitempty"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:"PmtMtd,omitempty"`
	MndtRltdInf    *MandateRelatedInformation14                  `xml:"MndtRltdInf,omitempty" json:"MndtRltdInf,omitempty"`
	RmtInf         *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	DbtrAcct       *CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct    *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct    *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr           *Party40Choice                                `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	CdtrAcct       *CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	Purp           *Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
}

func (o *OriginalTransactionReference28) AddIntrBkSttlmAmt() *ActiveOrHistoricCurrencyAndAmount {
	o.IntrBkSttlmAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return o.IntrBkSttlmAmt
}

func (o *OriginalTransactionReference28) AddAmt() *AmountType4Choice {
	o.Amt = new(AmountType4Choice)
	return o.Amt
}

func (o *OriginalTransactionReference28) SetIntrBkSttlmDt(value string) {
	o.IntrBkSttlmDt = (*ISODate)(&value)
}

func (o *OriginalTransactionReference28) SetReqdColltnDt(value string) {
	o.ReqdColltnDt = (*ISODate)(&value)
}

func (o *OriginalTransactionReference28) AddReqdExctnDt() *DateAndDateTime2Choice {
	o.ReqdExctnDt = new(DateAndDateTime2Choice)
	return o.ReqdExctnDt
}

func (o *OriginalTransactionReference28) AddCdtrSchmeId() *PartyIdentification135 {
	o.CdtrSchmeId = new(PartyIdentification135)
	return o.CdtrSchmeId
}

func (o *OriginalTransactionReference28) AddSttlmInf() *SettlementInstruction7 {
	o.SttlmInf = new(SettlementInstruction7)
	return o.SttlmInf
}

func (o *OriginalTransactionReference28) AddPmtTpInf() *PaymentTypeInformation27 {
	o.PmtTpInf = new(PaymentTypeInformation27)
	return o.PmtTpInf
}

func (o *OriginalTransactionReference28) SetPmtMtd(value string) {
	o.PmtMtd = (*PaymentMethod4Code)(&value)
}

func (o *OriginalTransactionReference28) AddMndtRltdInf() *MandateRelatedInformation14 {
	o.MndtRltdInf = new(MandateRelatedInformation14)
	return o.MndtRltdInf
}
func (o *OriginalTransactionReference28) AddRmtInf() *RemittanceInformation16 {
	o.RmtInf = new(RemittanceInformation16)
	return o.RmtInf
}

func (o *OriginalTransactionReference28) AddUltmtDbtr() *Party40Choice {
	o.UltmtDbtr = new(Party40Choice)
	return o.UltmtDbtr
}

func (o *OriginalTransactionReference28) AddDbtr() *Party40Choice {
	o.Dbtr = new(Party40Choice)
	return o.Dbtr
}

func (o *OriginalTransactionReference28) AddDbtrAcct() *CashAccount38 {
	o.DbtrAcct = new(CashAccount38)
	return o.DbtrAcct
}

func (o *OriginalTransactionReference28) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	o.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return o.DbtrAgt
}

func (o *OriginalTransactionReference28) AddDbtrAgtAcct() *CashAccount38 {
	o.DbtrAgtAcct = new(CashAccount38)
	return o.DbtrAgtAcct
}

func (o *OriginalTransactionReference28) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	o.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return o.CdtrAgt
}

func (o *OriginalTransactionReference28) AddCdtrAgtAcct() *CashAccount38 {
	o.CdtrAgtAcct = new(CashAccount38)
	return o.CdtrAgtAcct
}

func (o *OriginalTransactionReference28) AddCdtr() *Party40Choice {
	o.Cdtr = new(Party40Choice)
	return o.Cdtr
}

func (o *OriginalTransactionReference28) AddCdtrAcct() *CashAccount38 {
	o.CdtrAcct = new(CashAccount38)
	return o.CdtrAcct
}

func (o *OriginalTransactionReference28) AddUltmtCdtr() *Party40Choice {
	o.UltmtCdtr = new(Party40Choice)
	return o.UltmtCdtr
}

func (o *OriginalTransactionReference28) AddPurp() *Purpose2Choice {
	o.Purp = new(Purpose2Choice)
	return o.Purp
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

func (s *SettlementInstruction7) AddSttlmAcct() *CashAccount38 {
	s.SttlmAcct = new(CashAccount38)
	return s.SttlmAcct
}

func (s *SettlementInstruction7) AddClrSys() *ClearingSystemIdentification3Choice {
	s.ClrSys = new(ClearingSystemIdentification3Choice)
	return s.ClrSys
}

func (s *SettlementInstruction7) AddInstgRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification6 {
	s.InstgRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification6)
	return s.InstgRmbrsmntAgt
}

func (s *SettlementInstruction7) AddInstgRmbrsmntAgtAcct() *CashAccount38 {
	s.InstgRmbrsmntAgtAcct = new(CashAccount38)
	return s.InstgRmbrsmntAgtAcct
}

func (s *SettlementInstruction7) AddInstdRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification6 {
	s.InstdRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification6)
	return s.InstdRmbrsmntAgt
}

func (s *SettlementInstruction7) AddInstdRmbrsmntAgtAcct() *CashAccount38 {
	s.InstdRmbrsmntAgtAcct = new(CashAccount38)
	return s.InstdRmbrsmntAgtAcct
}

func (s *SettlementInstruction7) AddThrdRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification6 {
	s.ThrdRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification6)
	return s.ThrdRmbrsmntAgt
}

func (s *SettlementInstruction7) AddThrdRmbrsmntAgtAcct() *CashAccount38 {
	s.ThrdRmbrsmntAgtAcct = new(CashAccount38)
	return s.ThrdRmbrsmntAgtAcct
}

type CashAccount38 struct {
	Id   *AccountIdentification4Choice `xml:"Id" json:"Id"`
	Tp   *CashAccountType2Choice       `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Ccy  *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:"Ccy,omitempty"`
	Nm   *Max70Text                    `xml:"Nm,omitempty" json:"Nm,omitempty"`
	Prxy *ProxyAccountIdentification1  `xml:"Prxy,omitempty" json:"Prxy,omitempty"`
}

func (c *CashAccount38) AddId() *AccountIdentification4Choice {
	c.Id = new(AccountIdentification4Choice)
	return c.Id
}

func (c *CashAccount38) AddTp() *CashAccountType2Choice {
	c.Tp = new(CashAccountType2Choice)
	return c.Tp
}

func (c *CashAccount38) SetCcy(value string) {
	c.Ccy = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount38) SetNm(value string) {
	c.Nm = (*Max70Text)(&value)
}

func (c *CashAccount38) AddPrxy() *ProxyAccountIdentification1 {
	c.Prxy = new(ProxyAccountIdentification1)
	return c.Prxy
}

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
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:"InstrPrty,omitempty"`
	ClrChanl  *ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:"ClrChanl,omitempty"`
	SvcLvl    []*ServiceLevel8Choice  `xml:"SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	SeqTp     *SequenceType3Code      `xml:"SeqTp,omitempty" json:"SeqTp,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation27) SetTp(value string) {
	p.InstrPrty = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation27) SetlrChanl(value string) {
	p.ClrChanl = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation27) AddSvcLvl() *ServiceLevel8Choice {
	newValue := new(ServiceLevel8Choice)
	p.SvcLvl = append(p.SvcLvl, newValue)
	return newValue
}

func (p *PaymentTypeInformation27) AddLclInstrm() *LocalInstrument2Choice {
	p.LclInstrm = new(LocalInstrument2Choice)
	return p.LclInstrm
}

func (p *PaymentTypeInformation27) SetSeqTp(value string) {
	p.SeqTp = (*SequenceType3Code)(&value)
}

func (p *PaymentTypeInformation27) AddCtgyPurp() *CategoryPurpose1Choice {
	p.CtgyPurp = new(CategoryPurpose1Choice)
	return p.CtgyPurp
}

type MandateRelatedInformation14 struct {
	MndtId        *Max35Text                     `xml:"MndtId,omitempty" json:"MndtId,omitempty"`
	DtOfSgntr     *ISODate                       `xml:"DtOfSgntr,omitempty" json:"DtOfSgntr,omitempty"`
	AmdmntInd     *TrueFalseIndicator            `xml:"AmdmntInd,omitempty" json:"AmdmntInd,omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty" json:"AmdmntInfDtls,omitempty"`
	ElctrncSgntr  *Max1025Text                   `xml:"ElctrncSgntr,omitempty" json:"ElctrncSgntr,omitempty"`
	FrstColltnDt  *ISODate                       `xml:"FrstColltnDt,omitempty" json:"FrstColltnDt,omitempty"`
	FnlColltnDt   *ISODate                       `xml:"FnlColltnDt,omitempty" json:"FnlColltnDt,omitempty"`
	Frqcy         *Frequency36Choice             `xml:"Frqcy,omitempty" json:"Frqcy,omitempty"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
	TrckgDays     *Exact2NumericText             `xml:"TrckgDays,omitempty" json:"TrckgDays,omitempty"`
}

func (m *MandateRelatedInformation14) SetMndtId(value string) {
	m.MndtId = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation14) SetDtOfSgntr(value string) {
	m.DtOfSgntr = (*ISODate)(&value)
}

func (m *MandateRelatedInformation14) SetAmdmntInd(value string) {
	m.AmdmntInd = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation14) AddAmdmntInfDtls() *AmendmentInformationDetails13 {
	m.AmdmntInfDtls = new(AmendmentInformationDetails13)
	return m.AmdmntInfDtls
}

func (m *MandateRelatedInformation14) SetElctrncSgntr(value string) {
	m.ElctrncSgntr = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation14) SetFrstColltnDt(value string) {
	m.FrstColltnDt = (*ISODate)(&value)
}

func (m *MandateRelatedInformation14) SetFnlColltnDt(value string) {
	m.FnlColltnDt = (*ISODate)(&value)
}

func (m *MandateRelatedInformation14) AddFrqcy() *Frequency36Choice {
	m.Frqcy = new(Frequency36Choice)
	return m.Frqcy
}
func (m *MandateRelatedInformation14) AddRsn() *MandateSetupReason1Choice {
	m.Rsn = new(MandateSetupReason1Choice)
	return m.Rsn
}

func (m *MandateRelatedInformation14) SetTrckgDays(value string) {
	m.TrckgDays = (*Exact2NumericText)(&value)
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      *Max35Text                                    `xml:"OrgnlMndtId,omitempty" json:"OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty" json:"OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty" json:"OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct *CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:"OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        *PartyIdentification135                       `xml:"OrgnlDbtr,omitempty" json:"OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    *CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty" json:"OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty" json:"OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct *CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:"OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt *ISODate                                      `xml:"OrgnlFnlColltnDt,omitempty" json:"OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       *Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty" json:"OrgnlFrqcy,omitempty"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:"OrgnlRsn,omitempty"`
	OrgnlTrckgDays   *Exact2NumericText                            `xml:"OrgnlTrckgDays,omitempty" json:"OrgnlTrckgDays,omitempty"`
}

func (a *AmendmentInformationDetails13) SetOrgnlMndtId(value string) {
	a.OrgnlMndtId = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails13) AddOrgnlCdtrSchmeId() *PartyIdentification135 {
	a.OrgnlCdtrSchmeId = new(PartyIdentification135)
	return a.OrgnlCdtrSchmeId
}

func (a *AmendmentInformationDetails13) AddOrgnlCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	a.OrgnlCdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return a.OrgnlCdtrAgt
}

func (a *AmendmentInformationDetails13) AddOrgnlCdtrAgtAcct() *CashAccount38 {
	a.OrgnlCdtrAgtAcct = new(CashAccount38)
	return a.OrgnlCdtrAgtAcct
}

func (a *AmendmentInformationDetails13) AddOrgnlDbtr() *PartyIdentification135 {
	a.OrgnlDbtr = new(PartyIdentification135)
	return a.OrgnlDbtr
}

func (a *AmendmentInformationDetails13) AddOrgnlDbtrAcct() *CashAccount38 {
	a.OrgnlDbtrAcct = new(CashAccount38)
	return a.OrgnlDbtrAcct
}

func (a *AmendmentInformationDetails13) AddOrgnlDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	a.OrgnlDbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return a.OrgnlDbtrAgt
}

func (a *AmendmentInformationDetails13) AddOrgnlDbtrAgtAcct() *CashAccount38 {
	a.OrgnlDbtrAcct = new(CashAccount38)
	return a.OrgnlDbtrAcct
}

func (a *AmendmentInformationDetails13) SetOrgnlFnlColltnDt(value string) {
	a.OrgnlFnlColltnDt = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails13) AddOrgnlFrqcy() *Frequency36Choice {
	a.OrgnlFrqcy = new(Frequency36Choice)
	return a.OrgnlFrqcy
}

func (a *AmendmentInformationDetails13) AddOrgnlRsn() *MandateSetupReason1Choice {
	a.OrgnlRsn = new(MandateSetupReason1Choice)
	return a.OrgnlRsn
}

func (a *AmendmentInformationDetails13) SetTrckgDays(value string) {
	a.OrgnlTrckgDays = (*Exact2NumericText)(&value)
}

type Frequency36Choice struct {
	Tp     *Frequency6Code      `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Prd    *FrequencyPeriod1    `xml:"Prd,omitempty" json:"Prd,omitempty"`
	PtInTm *FrequencyAndMoment1 `xml:"PtInTm,omitempty" json:"PtInTm,omitempty"`
}

func (f *Frequency36Choice) SetTp(value string) {
	f.Tp = (*Frequency6Code)(&value)
}

func (f *Frequency36Choice) AddPrd() *FrequencyPeriod1 {
	f.Prd = new(FrequencyPeriod1)
	return f.Prd
}

func (f *Frequency36Choice) AddPtInTm() *FrequencyAndMoment1 {
	f.PtInTm = new(FrequencyAndMoment1)
	return f.PtInTm
}

type FrequencyAndMoment1 struct {
	Tp     *Frequency6Code    `xml:"Tp" json:"Tp"`
	PtInTm *Exact2NumericText `xml:"PtInTm" json:"PtInTm"`
}

func (f *FrequencyAndMoment1) SetTp(value string) {
	f.Tp = (*Frequency6Code)(&value)
}

func (f *FrequencyAndMoment1) SetPtInTm(value string) {
	f.PtInTm = (*Exact2NumericText)(&value)
}

// Must match the pattern [0-9]{2}
type Exact2NumericText string

type RemittanceInformation16 struct {
	Ustrd []*Max140Text                        `xml:"Ustrd,omitempty" json:"Ustrd,omitempty"`
	Strd  []*StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:"Strd,omitempty"`
}

func (r *RemittanceInformation16) AddUstrd(value string) {
	r.Ustrd = append(r.Ustrd, (*Max140Text)(&value))
}

func (r *RemittanceInformation16) AddStrd() *StructuredRemittanceInformation16 {
	newValue := new(StructuredRemittanceInformation16)
	r.Strd = append(r.Strd, newValue)
	return newValue
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []*ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:"RfrdDocInf,omitempty"`
	RfrdDocAmt  *RemittanceAmount2              `xml:"RfrdDocAmt,omitempty" json:"RfrdDocAmt,omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2  `xml:"CdtrRefInf,omitempty" json:"CdtrRefInf,omitempty"`
	Invcr       *PartyIdentification135         `xml:"Invcr,omitempty" json:"Invcr,omitempty"`
	Invcee      *PartyIdentification135         `xml:"Invcee,omitempty" json:"Invcee,omitempty"`
	TaxRmt      *TaxInformation7                `xml:"TaxRmt,omitempty" json:"TaxRmt,omitempty"`
	GrnshmtRmt  *Garnishment3                   `xml:"GrnshmtRmt,omitempty" json:"GrnshmtRmt,omitempty"`
	AddtlRmtInf []*Max140Text                   `xml:"AddtlRmtInf,omitempty" json:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation16) AddRfrdDocInf() *ReferredDocumentInformation7 {
	newValue := new(ReferredDocumentInformation7)
	s.RfrdDocInf = append(s.RfrdDocInf, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation16) AddRfrdDocAmt() *RemittanceAmount2 {
	s.RfrdDocAmt = new(RemittanceAmount2)
	return s.RfrdDocAmt
}

func (s *StructuredRemittanceInformation16) AddCdtrRefInf() *CreditorReferenceInformation2 {
	s.CdtrRefInf = new(CreditorReferenceInformation2)
	return s.CdtrRefInf
}

func (s *StructuredRemittanceInformation16) AddInvcr() *PartyIdentification135 {
	s.Invcr = new(PartyIdentification135)
	return s.Invcr
}

func (s *StructuredRemittanceInformation16) AddInvcee() *PartyIdentification135 {
	s.Invcee = new(PartyIdentification135)
	return s.Invcee
}

func (s *StructuredRemittanceInformation16) AddTaxRmt() *TaxInformation7 {
	s.TaxRmt = new(TaxInformation7)
	return s.TaxRmt
}

func (s *StructuredRemittanceInformation16) AddGrnshmtRmt() *Garnishment3 {
	s.GrnshmtRmt = new(Garnishment3)
	return s.GrnshmtRmt
}

func (s *StructuredRemittanceInformation16) AddAddtlRmtInf(value string) {
	s.AddtlRmtInf = append(s.AddtlRmtInf, (*Max140Text)(&value))
}

type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	UltmtDbtr       *TaxParty2                         `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	AdmstnZone      *Max35Text                         `xml:"AdmstnZone,omitempty" json:"AdmstnZone,omitempty"`
	RefNb           *Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Mtd             *Max35Text                         `xml:"Mtd,omitempty" json:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"TtlTaxAmt,omitempty"`
	Dt              *ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	SeqNb           *DecimalNumber                     `xml:"SeqNb,omitempty" json:"SeqNb,omitempty"`
	Rcrd            []*TaxRecord2                      `xml:"Rcrd,omitempty" json:"Rcrd,omitempty"`
}

// Cdtr            *TaxParty1
func (t *TaxInformation7) AddCdtr() *TaxParty1 {
	t.Cdtr = new(TaxParty1)
	return t.Cdtr
}

// Dbtr            *TaxParty2
func (t *TaxInformation7) AddDbtr() *TaxParty2 {
	t.Dbtr = new(TaxParty2)
	return t.Dbtr
}

// UltmtDbtr       *TaxParty2
func (t *TaxInformation7) AddUltmtDbtr() *TaxParty2 {
	t.UltmtDbtr = new(TaxParty2)
	return t.UltmtDbtr
}

// AdmstnZone      *Max35Text
func (t *TaxInformation7) SetAdmstnZone(value string) {
	t.AdmstnZone = (*Max35Text)(&value)
}

// RefNb           *Max140Text
func (t *TaxInformation7) SetRefNb(value string) {
	t.RefNb = (*Max140Text)(&value)
}

// Mtd             *Max35Text
func (t *TaxInformation7) SetMtd(value string) {
	t.Mtd = (*Max35Text)(&value)
}

// TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount
func (t *TaxInformation7) AddTtlTaxblBaseAmt() *ActiveOrHistoricCurrencyAndAmount {
	t.TtlTaxblBaseAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return t.TtlTaxblBaseAmt
}

// TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount
func (t *TaxInformation7) AddTtlTaxAmt() *ActiveOrHistoricCurrencyAndAmount {
	t.TtlTaxAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return t.TtlTaxAmt
}

// Dt              *ISODate
func (t *TaxInformation7) SetDt(value string) {
	t.Dt = (*ISODate)(&value)
}

// SeqNb           *DecimalNumber
func (t *TaxInformation7) SetSeqNb(value string) {
	t.SeqNb = (*DecimalNumber)(&value)
}

// Rcrd            []*TaxRecord2
func (t *TaxInformation7) AddC() *TaxRecord2 {
	newValue := new(TaxRecord2)
	t.Rcrd = append(t.Rcrd, newValue)
	return newValue
}

type TaxRecord2 struct {
	Tp       *Max35Text  `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Ctgy     *Max35Text  `xml:"Ctgy,omitempty" json:"Ctgy,omitempty"`
	CtgyDtls *Max35Text  `xml:"CtgyDtls,omitempty" json:"CtgyDtls,omitempty"`
	DbtrSts  *Max35Text  `xml:"DbtrSts,omitempty" json:"DbtrSts,omitempty"`
	CertId   *Max35Text  `xml:"CertId,omitempty" json:"CertId,omitempty"`
	FrmsCd   *Max35Text  `xml:"FrmsCd,omitempty" json:"FrmsCd,omitempty"`
	Prd      *TaxPeriod2 `xml:"Prd,omitempty" json:"Prd,omitempty"`
	TaxAmt   *TaxAmount2 `xml:"TaxAmt,omitempty" json:"TaxAmt,omitempty"`
	AddtlInf *Max140Text `xml:"AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

func (t *TaxRecord2) SetTp(value string) {
	t.Tp = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetCtgy(value string) {
	t.Ctgy = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetCtgyDtls(value string) {
	t.CtgyDtls = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetDbtrSts(value string) {
	t.DbtrSts = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetCertId(value string) {
	t.CertId = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetFrmsCd(value string) {
	t.FrmsCd = (*Max35Text)(&value)
}

func (t *TaxRecord2) AddPrd() *TaxPeriod2 {
	t.Prd = new(TaxPeriod2)
	return t.Prd
}

func (t *TaxRecord2) AddTtlTaxAmt() *TaxAmount2 {
	t.TaxAmt = new(TaxAmount2)
	return t.TaxAmt
}

func (t *TaxRecord2) SetAddtlInf(value string) {
	t.AddtlInf = (*Max140Text)(&value)
}

type TaxPeriod2 struct {
	Yr     *ISODate              `xml:"Yr,omitempty" json:"Yr,omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:"Tp,omitempty"`
	FrToDt *DatePeriod2          `xml:"FrToDt,omitempty" json:"FrToDt,omitempty"`
}

func (t *TaxPeriod2) SetYr(value string) {
	t.Yr = (*ISODate)(&value)
}

func (t *TaxPeriod2) SetTp(value string) {
	t.Tp = (*TaxRecordPeriod1Code)(&value)
}

func (t *TaxPeriod2) AddFrToDt() *DatePeriod2 {
	t.FrToDt = new(DatePeriod2)
	return t.FrToDt
}

type DatePeriod2 struct {
	FrDt *ISODate `xml:"FrDt" json:"FrDt"`
	ToDt *ISODate `xml:"ToDt" json:"ToDt"`
}

func (d *DatePeriod2) SetFrDt(value string) {
	d.FrDt = (*ISODate)(&value)
}

func (d *DatePeriod2) SetToDt(value string) {
	d.ToDt = (*ISODate)(&value)
}

type TaxAmount2 struct {
	Rate         *DecimalNumber                     `xml:"Rate,omitempty" json:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:"TtlAmt,omitempty"`
	Dtls         []*TaxRecordDetails2               `xml:"Dtls,omitempty" json:"Dtls,omitempty"`
}

func (t *TaxAmount2) SetRate(value string) {
	t.Rate = (*DecimalNumber)(&value)
}

func (t *TaxAmount2) AddTaxblBaseAmt() *ActiveOrHistoricCurrencyAndAmount {
	t.TaxblBaseAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return t.TaxblBaseAmt
}

func (t *TaxAmount2) AddTtlAmt() *ActiveOrHistoricCurrencyAndAmount {
	t.TtlAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return t.TtlAmt
}

func (t *TaxAmount2) AddDtls() *TaxRecordDetails2 {
	newValue := new(TaxRecordDetails2)
	t.Dtls = append(t.Dtls, newValue)
	return newValue
}

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                        `xml:"Prd,omitempty" json:"Prd,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"Amt"`
}

func (t *TaxRecordDetails2) AddPrd() *TaxPeriod2 {
	t.Prd = new(TaxPeriod2)
	return t.Prd
}

func (t *TaxRecordDetails2) AddAmt() *ActiveOrHistoricCurrencyAndAmount {
	t.Amt = new(ActiveOrHistoricCurrencyAndAmount)
	return t.Amt
}

type Garnishment3 struct {
	Tp                *GarnishmentType1                  `xml:"Tp" json:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee,omitempty" json:"Grnshee,omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:"GrnshmtAdmstr,omitempty"`
	RefNb             *Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Dt                *ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"RmtdAmt,omitempty"`
	FmlyMdclInsrncInd *TrueFalseIndicator                `xml:"FmlyMdclInsrncInd,omitempty" json:"FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  *TrueFalseIndicator                `xml:"MplyeeTermntnInd,omitempty" json:"MplyeeTermntnInd,omitempty"`
}

func (g *Garnishment3) AddTp() *GarnishmentType1 {
	g.Tp = new(GarnishmentType1)
	return g.Tp
}

func (g *Garnishment3) AddGrnshee() *PartyIdentification135 {
	g.Grnshee = new(PartyIdentification135)
	return g.Grnshee
}

func (g *Garnishment3) AddGrnshmtAdmstr() *PartyIdentification135 {
	g.GrnshmtAdmstr = new(PartyIdentification135)
	return g.GrnshmtAdmstr
}

func (g *Garnishment3) SetRefNb(value string) {
	g.RefNb = (*Max140Text)(&value)
}

func (g *Garnishment3) SetDt(value string) {
	g.Dt = (*ISODate)(&value)
}

func (g *Garnishment3) AddRmtdAmt() *ActiveOrHistoricCurrencyAndAmount {
	g.RmtdAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return g.RmtdAmt
}

func (g *Garnishment3) SetFmlyMdclInsrncInd(value string) {
	g.FmlyMdclInsrncInd = (*TrueFalseIndicator)(&value)
}

func (g *Garnishment3) SetMplyeeTermntnInd(value string) {
	g.MplyeeTermntnInd = (*TrueFalseIndicator)(&value)
}

type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"Pty,omitempty" json:"Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty" json:"Agt,omitempty"`
}

func (p *Party40Choice) AddPty() *PartyIdentification135 {
	p.Pty = new(PartyIdentification135)
	return p.Pty
}

func (p *Party40Choice) AddAgt() *BranchAndFinancialInstitutionIdentification6 {
	p.Agt = new(BranchAndFinancialInstitutionIdentification6)
	return p.Agt
}

type BI_SupplementaryData1 struct {
	PlcAndNm *Max350Text                    `xml:"PlcAndNm,omitempty" json:"PlcAndNm,omitempty"`
	Envlp    *BI_SupplementaryDataEnvelope1 `xml:"Envlp" json:"Envlp"`
}

func (b *BI_SupplementaryData1) SetPlcAndNm(value string) {
	b.PlcAndNm = (*Max350Text)(&value)
}

func (b *BI_SupplementaryData1) AddEnvlp() *BI_SupplementaryDataEnvelope1 {
	b.Envlp = new(BI_SupplementaryDataEnvelope1)
	return b.Envlp
}

type BI_SupplementaryDataEnvelope1 struct {
	Dbtr        *BI_AddtlCstmrInf `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	Cdtr        *BI_AddtlCstmrInf `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	DbtrAgtAcct *CashAccount38    `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgtAcct *CashAccount38    `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cstmr       *BI_AddtlCstmrInf `xml:"Cstmr,omitempty" json:"Cstmr,omitempty"`
}

func (b *BI_SupplementaryDataEnvelope1) AddDbtr() *BI_AddtlCstmrInf {
	b.Dbtr = new(BI_AddtlCstmrInf)
	return b.Dbtr
}

func (b *BI_SupplementaryDataEnvelope1) AddCdtr() *BI_AddtlCstmrInf {
	b.Cdtr = new(BI_AddtlCstmrInf)
	return b.Cdtr
}

func (b *BI_SupplementaryDataEnvelope1) AddDbtrAgtAcct() *CashAccount38 {
	b.DbtrAgtAcct = new(CashAccount38)
	return b.DbtrAgtAcct
}

func (b *BI_SupplementaryDataEnvelope1) AddCdtrAgtAcct() *CashAccount38 {
	b.CdtrAgtAcct = new(CashAccount38)
	return b.CdtrAgtAcct
}

func (b *BI_SupplementaryDataEnvelope1) AddCstmr() *BI_AddtlCstmrInf {
	b.Cstmr = new(BI_AddtlCstmrInf)
	return b.Cstmr
}

type BI_AddtlCstmrInf struct {
	Tp       *Max35Text `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Id       *Max35Text `xml:"Id,omitempty" json:"Id,omitempty"`
	RsdntSts *Max35Text `xml:"RsdntSts,omitempty" json:"RsdntSts,omitempty"`
	TwnNm    *Max35Text `xml:"TwnNm,omitempty" json:"TwnNm,omitempty"`
}

func (b *BI_AddtlCstmrInf) SetTp(value string) {
	b.Tp = (*Max35Text)(&value)
}

func (b *BI_AddtlCstmrInf) SetId(value string) {
	b.Id = (*Max35Text)(&value)
}

func (b *BI_AddtlCstmrInf) SetRsdntSts(value string) {
	b.RsdntSts = (*Max35Text)(&value)
}

func (b *BI_AddtlCstmrInf) SetTwnNm(value string) {
	b.TwnNm = (*Max35Text)(&value)
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

func (c *CreditTransferTransaction39) AddPmtId() *PaymentIdentification7 {
	c.PmtId = new(PaymentIdentification7)
	return c.PmtId
}

func (c *CreditTransferTransaction39) AddPmtTpInf() *PaymentTypeInformation28 {
	c.PmtTpInf = new(PaymentTypeInformation28)
	return c.PmtTpInf
}

func (c *CreditTransferTransaction39) SetInterbankSettlementAmount(value, currency string) {
	c.IntrBkSttlmAmt = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransaction39) SetIntrBkSttlmDt(value string) {
	c.IntrBkSttlmDt = (*ISODate)(&value)
}

func (c *CreditTransferTransaction39) SetSttlmPrty(value string) {
	c.SttlmPrty = (*Priority3Code)(&value)
}

func (c *CreditTransferTransaction39) AddSttlmTmIndctn() *SettlementDateTimeIndication1 {
	c.SttlmTmIndctn = new(SettlementDateTimeIndication1)
	return c.SttlmTmIndctn
}

func (c *CreditTransferTransaction39) AddSttlmTmReq() *SettlementTimeRequest2 {
	c.SttlmTmReq = new(SettlementTimeRequest2)
	return c.SttlmTmReq
}

func (c *CreditTransferTransaction39) SetAccptncDtTm(value string) {
	c.AccptncDtTm = (*ISODateTime)(&value)
}

func (c *CreditTransferTransaction39) SetPoolgAdjstmntDt(value string) {
	c.PoolgAdjstmntDt = (*ISODate)(&value)
}

func (c *CreditTransferTransaction39) AddInstdAmt() *ActiveOrHistoricCurrencyAndAmount {
	c.InstdAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return c.InstdAmt
}

func (c *CreditTransferTransaction39) SetXchgRate(value string) {
	c.XchgRate = (*DecimalNumber)(&value)
}

func (c *CreditTransferTransaction39) SetChrgBr(value string) {
	c.ChrgBr = (*ChargeBearerType1Code)(&value)
}

func (c *CreditTransferTransaction39) AddChrgsInf() *Charges7 {
	newValue := new(Charges7)
	c.ChrgsInf = append(c.ChrgsInf, newValue)
	return newValue
}

func (c *CreditTransferTransaction39) AddPrvsInstgAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt1
}

func (c *CreditTransferTransaction39) AddPrvsInstgAgt1Acct() *CashAccount38 {
	c.PrvsInstgAgt1Acct = new(CashAccount38)
	return c.PrvsInstgAgt1Acct
}

func (c *CreditTransferTransaction39) AddPrvsInstgAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt2
}

func (c *CreditTransferTransaction39) AddPrvsInstgAgt2Acct() *CashAccount38 {
	c.PrvsInstgAgt2Acct = new(CashAccount38)
	return c.PrvsInstgAgt2Acct
}

func (c *CreditTransferTransaction39) AddPrvsInstgAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt3
}

func (c *CreditTransferTransaction39) AddPrvsInstgAgt3Acct() *CashAccount38 {
	c.PrvsInstgAgt3Acct = new(CashAccount38)
	return c.PrvsInstgAgt3Acct
}

func (c *CreditTransferTransaction39) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.InstgAgt
}

func (c *CreditTransferTransaction39) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.InstdAgt
}

func (c *CreditTransferTransaction39) AddIntrmyAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt1
}

func (c *CreditTransferTransaction39) AddIntrmyAgt1Acct() *CashAccount38 {
	c.IntrmyAgt1Acct = new(CashAccount38)
	return c.IntrmyAgt1Acct
}

func (c *CreditTransferTransaction39) AddIntrmyAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt2
}

func (c *CreditTransferTransaction39) AddIntrmyAgt2Acct() *CashAccount38 {
	c.IntrmyAgt2Acct = new(CashAccount38)
	return c.IntrmyAgt2Acct
}

func (c *CreditTransferTransaction39) AddIntrmyAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt3
}

func (c *CreditTransferTransaction39) AddIntrmyAgt3Acct() *CashAccount38 {
	c.IntrmyAgt3Acct = new(CashAccount38)
	return c.IntrmyAgt3Acct
}

func (c *CreditTransferTransaction39) AddUltmtDbtr() *PartyIdentification135 {
	c.UltmtDbtr = new(PartyIdentification135)
	return c.UltmtDbtr
}

func (c *CreditTransferTransaction39) AddInitgPty() *PartyIdentification135 {
	c.InitgPty = new(PartyIdentification135)
	return c.InitgPty
}

func (c *CreditTransferTransaction39) AddDbtr() *PartyIdentification135 {
	c.Dbtr = new(PartyIdentification135)
	return c.Dbtr
}

func (c *CreditTransferTransaction39) AddDbtrAcct() *CashAccount38 {
	c.DbtrAcct = new(CashAccount38)
	return c.DbtrAcct
}

func (c *CreditTransferTransaction39) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.DbtrAgt
}

func (c *CreditTransferTransaction39) AddDbtrAgtAcct() *CashAccount38 {
	c.DbtrAgtAcct = new(CashAccount38)
	return c.DbtrAgtAcct
}

func (c *CreditTransferTransaction39) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.CdtrAgt
}

func (c *CreditTransferTransaction39) AddCdtrAgtAcct() *CashAccount38 {
	c.CdtrAgtAcct = new(CashAccount38)
	return c.CdtrAgtAcct
}

func (c *CreditTransferTransaction39) AddCdtr() *PartyIdentification135 {
	c.Cdtr = new(PartyIdentification135)
	return c.Cdtr
}

func (c *CreditTransferTransaction39) AddCdtrAcct() *CashAccount38 {
	c.CdtrAcct = new(CashAccount38)
	return c.CdtrAcct
}

func (c *CreditTransferTransaction39) AddUltmtCdtr() *PartyIdentification135 {
	c.UltmtCdtr = new(PartyIdentification135)
	return c.UltmtCdtr
}

func (c *CreditTransferTransaction39) AddInstrForCdtrAgt() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstrForCdtrAgt = append(c.InstrForCdtrAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction39) AddInstrForNxtAgt() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstrForNxtAgt = append(c.InstrForNxtAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction39) AddPurp() *Purpose2Choice {
	c.Purp = new(Purpose2Choice)
	return c.Purp
}

func (c *CreditTransferTransaction39) AddRgltryRptg() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	c.RgltryRptg = append(c.RgltryRptg, newValue)
	return newValue
}

func (c *CreditTransferTransaction39) AddTax() *TaxInformation8 {
	c.Tax = new(TaxInformation8)
	return c.Tax
}

func (c *CreditTransferTransaction39) AddRltdRmtInf() *RemittanceLocation7 {
	newValue := new(RemittanceLocation7)
	c.RltdRmtInf = append(c.RltdRmtInf, newValue)
	return newValue
}

func (c *CreditTransferTransaction39) AddRmtInf() *RemittanceInformation16 {
	c.RmtInf = new(RemittanceInformation16)
	return c.RmtInf
}

func (c *CreditTransferTransaction39) AddSplmtryData() *BI_SupplementaryData1 {
	newValue := new(BI_SupplementaryData1)
	c.SplmtryData = append(c.SplmtryData, newValue)
	return newValue
}

type PaymentIdentification7 struct {
	InstrId    *Max35Text        `xml:"InstrId,omitempty" json:"InstrId,omitempty"`
	EndToEndId *Max35Text        `xml:"EndToEndId" json:"EndToEndId"`
	TxId       *Max35Text        `xml:"TxId,omitempty" json:"TxId,omitempty"`
	UETR       *UUIDv4Identifier `xml:"UETR,omitempty" json:"UETR,omitempty"`
	ClrSysRef  *Max35Text        `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
}

func (p *PaymentIdentification7) SetInstrId(value string) {
	p.InstrId = (*Max35Text)(&value)
}

func (p *PaymentIdentification7) SetEndToEndId(value string) {
	p.EndToEndId = (*Max35Text)(&value)
}

func (p *PaymentIdentification7) SetTxId(value string) {
	p.TxId = (*Max35Text)(&value)
}

func (p *PaymentIdentification7) SetUETR(value string) {
	p.UETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentIdentification7) SetClrSysRef(value string) {
	p.ClrSysRef = (*Max35Text)(&value)
}

type TaxInformation8 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	AdmstnZone      *Max35Text                         `xml:"AdmstnZone,omitempty" json:"AdmstnZone,omitempty"`
	RefNb           *Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Mtd             *Max35Text                         `xml:"Mtd,omitempty" json:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"TtlTaxAmt,omitempty"`
	Dt              *ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	SeqNb           *DecimalNumber                     `xml:"SeqNb,omitempty" json:"SeqNb,omitempty"`
	Rcrd            []*TaxRecord2                      `xml:"Rcrd,omitempty" json:"Rcrd,omitempty"`
}

func (t *TaxInformation8) AddCdtr() *TaxParty1 {
	t.Cdtr = new(TaxParty1)
	return t.Cdtr
}

func (t *TaxInformation8) AddDbtr() *TaxParty2 {
	t.Dbtr = new(TaxParty2)
	return t.Dbtr
}

func (t *TaxInformation8) SetAdmstnZone(value string) {
	t.AdmstnZone = (*Max35Text)(&value)
}

func (t *TaxInformation8) SetRefNb(value string) {
	t.RefNb = (*Max140Text)(&value)
}

func (t *TaxInformation8) SetMtd(value string) {
	t.Mtd = (*Max35Text)(&value)
}

func (t *TaxInformation8) AddTtlTaxblBaseAmt() *ActiveOrHistoricCurrencyAndAmount {
	t.TtlTaxblBaseAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return t.TtlTaxblBaseAmt
}

func (t *TaxInformation8) AddTtlTaxAmt() *ActiveOrHistoricCurrencyAndAmount {
	t.TtlTaxAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return t.TtlTaxAmt
}

func (t *TaxInformation8) SetDt(value string) {
	t.Dt = (*ISODate)(&value)
}

func (t *TaxInformation8) SetSeqNb(value string) {
	t.SeqNb = (*DecimalNumber)(&value)
}

func (t *TaxInformation8) AddRcrd() *TaxRecord2 {
	newValue := new(TaxRecord2)
	t.Rcrd = append(t.Rcrd, newValue)
	return newValue
}

type RemittanceLocation7 struct {
	RmtId       *Max35Text                 `xml:"RmtId,omitempty" json:"RmtId,omitempty"`
	RmtLctnDtls []*RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty" json:"RmtLctnDtls,omitempty"`
}

func (r *RemittanceLocation7) SetRmtId(value string) {
	r.RmtId = (*Max35Text)(&value)
}

func (r *RemittanceLocation7) AddRmtLctnDtls() *RemittanceLocationData1 {
	newValue := new(RemittanceLocationData1)
	r.RmtLctnDtls = append(r.RmtLctnDtls, newValue)
	return newValue
}

type RemittanceLocationData1 struct {
	Mtd        *RemittanceLocationMethod2Code `xml:"Mtd" json:"Mtd"`
	ElctrncAdr *Max2048Text                   `xml:"ElctrncAdr,omitempty" json:"ElctrncAdr,omitempty"`
	PstlAdr    *NameAndAddress16              `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

func (r *RemittanceLocationData1) SetMtd(value string) {
	r.Mtd = (*RemittanceLocationMethod2Code)(&value)
}

func (r *RemittanceLocationData1) SetElctrncAdr(value string) {
	r.ElctrncAdr = (*Max2048Text)(&value)
}

func (r *RemittanceLocationData1) AddPstlAdr() *NameAndAddress16 {
	r.PstlAdr = new(NameAndAddress16)
	return r.PstlAdr
}

type NameAndAddress16 struct {
	Nm  *Max140Text      `xml:"Nm" json:"Nm"`
	Adr *PostalAddress24 `xml:"Adr" json:"Adr"`
}

func (n *NameAndAddress16) SetNm(value string) {
	n.Nm = (*Max140Text)(&value)
}

func (n *NameAndAddress16) AddPstlAdr() *PostalAddress24 {
	n.Adr = new(PostalAddress24)
	return n.Adr
}

// TODO: pacs.009.001.09, json tag

type CreditTransferTransaction44 struct {
	PmtId              *PaymentIdentification13                      `xml:"PmtId" json:"PmtId"`
	PmtTpInf           *PaymentTypeInformation28                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	IntrBkSttlmAmt     *ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt" json:"IntrBkSttlmAmt"`
	IntrBkSttlmDt      *ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmPrty          *Priority3Code                                `xml:"SttlmPrty,omitempty" json:"SttlmPrty,omitempty"`
	SttlmTmIndctn      *SettlementDateTimeIndication1                `xml:"SttlmTmIndctn,omitempty" json:"SttlmTmIndctn,omitempty"`
	SttlmTmReq         *SettlementTimeRequest2                       `xml:"SttlmTmReq,omitempty" json:"SttlmTmReq,omitempty"`
	PrvsInstgAgt1      *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct  *CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2      *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct  *CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3      *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct  *CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	InstgAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	IntrmyAgt1         *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct     *CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2         *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct     *CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3         *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct     *CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	UltmtDbtr          *BranchAndFinancialInstitutionIdentification6 `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr               *BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr" json:"Dbtr"`
	DbtrAcct           *CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt            *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct        *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt            *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct        *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr               *BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr" json:"Cdtr"`
	CdtrAcct           *CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr          *BranchAndFinancialInstitutionIdentification6 `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt    []*InstructionForCreditorAgent3               `xml:"InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt     []*InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Purp               *Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
	RmtInf             *RemittanceInformation2                       `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UndrlygCstmrCdtTrf *CreditTransferTransaction45                  `xml:"UndrlygCstmrCdtTrf,omitempty" json:"UndrlygCstmrCdtTrf,omitempty"`
	SplmtryData        []*SupplementaryData1                         `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

func (c *CreditTransferTransaction44) AddPmtId() *PaymentIdentification13 {
	c.PmtId = new(PaymentIdentification13)
	return c.PmtId
}

func (c *CreditTransferTransaction44) AddPmtTpInf() *PaymentTypeInformation28 {
	c.PmtTpInf = new(PaymentTypeInformation28)
	return c.PmtTpInf
}

func (c *CreditTransferTransaction44) AddIntrBkSttlmAmt() *ActiveCurrencyAndAmount {
	c.IntrBkSttlmAmt = new(ActiveCurrencyAndAmount)
	return c.IntrBkSttlmAmt
}

func (c *CreditTransferTransaction44) SetIntrBkSttlmDt(value string) {
	c.IntrBkSttlmDt = (*ISODate)(&value)
}

func (c *CreditTransferTransaction44) SetSttlmPrty(value string) {
	c.SttlmPrty = (*Priority3Code)(&value)
}

func (c *CreditTransferTransaction44) AddSttlmTmIndctn() *SettlementDateTimeIndication1 {
	c.SttlmTmIndctn = new(SettlementDateTimeIndication1)
	return c.SttlmTmIndctn
}

func (c *CreditTransferTransaction44) AddSttlmTmReq() *SettlementTimeRequest2 {
	c.SttlmTmReq = new(SettlementTimeRequest2)
	return c.SttlmTmReq
}

func (c *CreditTransferTransaction44) AddPrvsInstgAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt1
}

func (c *CreditTransferTransaction44) AddPrvsInstgAgt1Acct() *CashAccount38 {
	c.PrvsInstgAgt1Acct = new(CashAccount38)
	return c.PrvsInstgAgt1Acct
}

func (c *CreditTransferTransaction44) AddPrvsInstgAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt2
}

func (c *CreditTransferTransaction44) AddPrvsInstgAgt2Acct() *CashAccount38 {
	c.PrvsInstgAgt2Acct = new(CashAccount38)
	return c.PrvsInstgAgt2Acct
}

func (c *CreditTransferTransaction44) AddPrvsInstgAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt3
}

func (c *CreditTransferTransaction44) AddPrvsInstgAgt3Acct() *CashAccount38 {
	c.PrvsInstgAgt3Acct = new(CashAccount38)
	return c.PrvsInstgAgt3Acct
}

func (c *CreditTransferTransaction44) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.InstgAgt
}

func (c *CreditTransferTransaction44) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.InstdAgt
}

func (c *CreditTransferTransaction44) AddIntrmyAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt1
}

func (c *CreditTransferTransaction44) AddIntrmyAgt1Acct() *CashAccount38 {
	c.IntrmyAgt1Acct = new(CashAccount38)
	return c.IntrmyAgt1Acct
}

func (c *CreditTransferTransaction44) AddIntrmyAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt2
}

func (c *CreditTransferTransaction44) AddIntrmyAgt2Acct() *CashAccount38 {
	c.IntrmyAgt2Acct = new(CashAccount38)
	return c.IntrmyAgt2Acct
}

func (c *CreditTransferTransaction44) AddIntrmyAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt3
}

func (c *CreditTransferTransaction44) AddIntrmyAgt3Acct() *CashAccount38 {
	c.IntrmyAgt3Acct = new(CashAccount38)
	return c.IntrmyAgt3Acct
}

func (c *CreditTransferTransaction44) AddUltmtDbtr() *BranchAndFinancialInstitutionIdentification6 {
	c.UltmtDbtr = new(BranchAndFinancialInstitutionIdentification6)
	return c.UltmtDbtr
}

func (c *CreditTransferTransaction44) AddDbtr() *BranchAndFinancialInstitutionIdentification6 {
	c.Dbtr = new(BranchAndFinancialInstitutionIdentification6)
	return c.Dbtr
}

func (c *CreditTransferTransaction44) AddDbtrAcct() *CashAccount38 {
	c.DbtrAcct = new(CashAccount38)
	return c.DbtrAcct
}

func (c *CreditTransferTransaction44) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.DbtrAgt
}

func (c *CreditTransferTransaction44) AddDbtrAgtAcct() *CashAccount38 {
	c.DbtrAgtAcct = new(CashAccount38)
	return c.DbtrAgtAcct
}

func (c *CreditTransferTransaction44) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.CdtrAgt
}

func (c *CreditTransferTransaction44) AddCdtrAgtAcct() *CashAccount38 {
	c.CdtrAgtAcct = new(CashAccount38)
	return c.CdtrAgtAcct
}

func (c *CreditTransferTransaction44) AddCdtr() *BranchAndFinancialInstitutionIdentification6 {
	c.Cdtr = new(BranchAndFinancialInstitutionIdentification6)
	return c.Cdtr
}

func (c *CreditTransferTransaction44) AddCdtrAcct() *CashAccount38 {
	c.CdtrAcct = new(CashAccount38)
	return c.CdtrAcct
}

func (c *CreditTransferTransaction44) AddUltmtCdtr() *BranchAndFinancialInstitutionIdentification6 {
	c.UltmtCdtr = new(BranchAndFinancialInstitutionIdentification6)
	return c.UltmtCdtr
}

func (c *CreditTransferTransaction44) AddInstrForCdtrAgt() *InstructionForCreditorAgent3 {
	newValue := new(InstructionForCreditorAgent3)
	c.InstrForCdtrAgt = append(c.InstrForCdtrAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction44) AddInstrForNxtAgt() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstrForNxtAgt = append(c.InstrForNxtAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction44) AddPurp() *Purpose2Choice {
	c.Purp = new(Purpose2Choice)
	return c.Purp
}

func (c *CreditTransferTransaction44) AddRmtInf() *RemittanceInformation2 {
	c.RmtInf = new(RemittanceInformation2)
	return c.RmtInf
}

func (c *CreditTransferTransaction44) AddUndrlygCstmrCdtTrf() *CreditTransferTransaction45 {
	c.UndrlygCstmrCdtTrf = new(CreditTransferTransaction45)
	return c.UndrlygCstmrCdtTrf
}

func (c *CreditTransferTransaction44) AddSplmtryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SplmtryData = append(c.SplmtryData, newValue)
	return newValue
}

type PaymentIdentification13 struct {
	InstrId    *Max35Text        `xml:"InstrId,omitempty" json:"InstrId,omitempty"`
	EndToEndId *Max35Text        `xml:"EndToEndId" json:"EndToEndId"`
	TxId       *Max35Text        `xml:"TxId,omitempty" json:"TxId,omitempty"`
	UETR       *UUIDv4Identifier `xml:"UETR,omitempty" json:"UETR,omitempty"`
	ClrSysRef  *Max35Text        `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
}

func (p *PaymentIdentification13) SetInstrId(value string) {
	p.InstrId = (*Max35Text)(&value)
}

func (p *PaymentIdentification13) SetEndToEndId(value string) {
	p.EndToEndId = (*Max35Text)(&value)
}

func (p *PaymentIdentification13) SetTxId(value string) {
	p.TxId = (*Max35Text)(&value)
}

func (p *PaymentIdentification13) SetUETR(value string) {
	p.UETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentIdentification13) SetClrSysRef(value string) {
	p.ClrSysRef = (*Max35Text)(&value)
}

type InstructionForCreditorAgent3 struct {
	Cd       ExternalCreditorAgentInstruction1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	InstrInf Max140Text                            `xml:"InstrInf,omitempty" json:"InstrInf,omitempty"`
}

// May be no more than 4 items long
type ExternalCreditorAgentInstruction1Code string

type CreditTransferTransaction45 struct {
	UltmtDbtr         *PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	InitgPty          *PartyIdentification135                       `xml:"InitgPty,omitempty" json:"InitgPty,omitempty"`
	Dbtr              *PartyIdentification135                       `xml:"Dbtr" json:"Dbtr"`
	DbtrAcct          *CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt" json:"DbtrAgt"`
	DbtrAgtAcct       *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	PrvsInstgAgt1     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct *CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct *CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct *CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	IntrmyAgt1        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct    *CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct    *CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct    *CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	CdtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt" json:"CdtrAgt"`
	CdtrAgtAcct       *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr              *PartyIdentification135                       `xml:"Cdtr" json:"Cdtr"`
	CdtrAcct          *CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr         *PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt   []*InstructionForCreditorAgent3               `xml:"InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt    []*InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Tax               *TaxInformation8                              `xml:"Tax,omitempty" json:"Tax,omitempty"`
	RmtInf            *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	InstdAmt          *ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty" json:"InstdAmt,omitempty"`
}

func (c *CreditTransferTransaction45) AddUltmtDbtr() *PartyIdentification135 {
	c.UltmtDbtr = new(PartyIdentification135)
	return c.UltmtDbtr
}

func (c *CreditTransferTransaction45) AddInitgPty() *PartyIdentification135 {
	c.InitgPty = new(PartyIdentification135)
	return c.InitgPty
}

func (c *CreditTransferTransaction45) AddDbtr() *PartyIdentification135 {
	c.Dbtr = new(PartyIdentification135)
	return c.Dbtr
}

func (c *CreditTransferTransaction45) AddDbtrAcct() *CashAccount38 {
	c.DbtrAcct = new(CashAccount38)
	return c.DbtrAcct
}

func (c *CreditTransferTransaction45) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.DbtrAgt
}

func (c *CreditTransferTransaction45) AddDbtrAgtAcct() *CashAccount38 {
	c.DbtrAgtAcct = new(CashAccount38)
	return c.DbtrAgtAcct
}

func (c *CreditTransferTransaction45) AddPrvsInstgAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt1
}

func (c *CreditTransferTransaction45) AddPrvsInstgAgt1Acct() *CashAccount38 {
	c.PrvsInstgAgt1Acct = new(CashAccount38)
	return c.PrvsInstgAgt1Acct
}

func (c *CreditTransferTransaction45) AddPrvsInstgAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt2
}

func (c *CreditTransferTransaction45) AddPrvsInstgAgt2Acct() *CashAccount38 {
	c.PrvsInstgAgt2Acct = new(CashAccount38)
	return c.PrvsInstgAgt2Acct
}

func (c *CreditTransferTransaction45) AddPrvsInstgAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt3
}

func (c *CreditTransferTransaction45) AddPrvsInstgAgt3Acct() *CashAccount38 {
	c.PrvsInstgAgt3Acct = new(CashAccount38)
	return c.PrvsInstgAgt3Acct
}

func (c *CreditTransferTransaction45) AddIntrmyAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt1
}

func (c *CreditTransferTransaction45) AddIntrmyAgt1Acct() *CashAccount38 {
	c.IntrmyAgt1Acct = new(CashAccount38)
	return c.IntrmyAgt1Acct
}

func (c *CreditTransferTransaction45) AddIntrmyAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt2
}

func (c *CreditTransferTransaction45) AddIntrmyAgt2Acct() *CashAccount38 {
	c.IntrmyAgt2Acct = new(CashAccount38)
	return c.IntrmyAgt2Acct
}

func (c *CreditTransferTransaction45) AddIntrmyAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt3
}

func (c *CreditTransferTransaction45) AddIntrmyAgt3Acct() *CashAccount38 {
	c.IntrmyAgt3Acct = new(CashAccount38)
	return c.IntrmyAgt3Acct
}

func (c *CreditTransferTransaction45) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.CdtrAgt
}

func (c *CreditTransferTransaction45) AddCdtrAgtAcct() *CashAccount38 {
	c.CdtrAgtAcct = new(CashAccount38)
	return c.CdtrAgtAcct
}

func (c *CreditTransferTransaction45) AddCdtr() *PartyIdentification135 {
	c.Cdtr = new(PartyIdentification135)
	return c.Cdtr
}

func (c *CreditTransferTransaction45) AddCdtrAcct() *CashAccount38 {
	c.CdtrAcct = new(CashAccount38)
	return c.CdtrAcct
}

func (c *CreditTransferTransaction45) AddUltmtCdtr() *PartyIdentification135 {
	c.UltmtCdtr = new(PartyIdentification135)
	return c.UltmtCdtr
}

func (c *CreditTransferTransaction45) AddInstrForCdtrAgt() *InstructionForCreditorAgent3 {
	newValue := new(InstructionForCreditorAgent3)
	c.InstrForCdtrAgt = append(c.InstrForCdtrAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction45) AddInstrForNxtAgt() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstrForNxtAgt = append(c.InstrForNxtAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction45) AddTax() *TaxInformation8 {
	c.Tax = new(TaxInformation8)
	return c.Tax
}

func (c *CreditTransferTransaction45) AddRmtInf() *RemittanceInformation16 {
	c.RmtInf = new(RemittanceInformation16)
	return c.RmtInf
}

func (c *CreditTransferTransaction45) AddInstdAmt() *ActiveOrHistoricCurrencyAndAmount {
	c.InstdAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return c.InstdAmt
}

// TODO: pacs.028.001.04, json tag

type OriginalGroupInformation27 struct {
	OrgnlMsgId   *Max35Text        `xml:"OrgnlMsgId" json:"OrgnlMsgId"`
	OrgnlMsgNmId *Max35Text        `xml:"OrgnlMsgNmId" json:"OrgnlMsgNmId"`
	OrgnlCreDtTm *ISODateTime      `xml:"OrgnlCreDtTm,omitempty" json:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty" json:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum *DecimalNumber    `xml:"OrgnlCtrlSum,omitempty" json:"OrgnlCtrlSum,omitempty"`
}

func (o *OriginalGroupInformation27) SetOrgnlMsgId(value string) {
	o.OrgnlMsgId = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation27) SetOrgnlMsgNmId(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation27) SetOrgnlCreDtTm(value string) {
	o.OrgnlCreDtTm = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation27) SetOrgnlNbOfTxs(value string) {
	o.OrgnlNbOfTxs = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation27) SetOrgnlCtrlSum(value string) {
	o.OrgnlCtrlSum = (*DecimalNumber)(&value)
}

type PaymentTransaction121 struct {
	StsReqId        *Max35Text                                    `xml:"StsReqId,omitempty" json:"StsReqId,omitempty"`
	OrgnlGrpInf     *OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId    *Max35Text                                    `xml:"OrgnlInstrId,omitempty" json:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId *Max35Text                                    `xml:"OrgnlEndToEndId,omitempty" json:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId       *Max35Text                                    `xml:"OrgnlTxId,omitempty" json:"OrgnlTxId,omitempty"`
	OrgnlUETR       *UUIDv4Identifier                             `xml:"OrgnlUETR,omitempty" json:"OrgnlUETR,omitempty"`
	AccptncDtTm     *ISODateTime                                  `xml:"AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	ClrSysRef       *Max35Text                                    `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
	InstgAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	OrgnlTxRef      *OriginalTransactionReference31               `xml:"OrgnlTxRef,omitempty" json:"OrgnlTxRef,omitempty"`
	SplmtryData     []*SupplementaryData1                         `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

// StsReqId        *Max35Text
func (p *PaymentTransaction121) SetStsReqId(value string) {
	p.StsReqId = (*Max35Text)(&value)
}

// OrgnlGrpInf     *OriginalGroupInformation29
func (p *PaymentTransaction121) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	p.OrgnlGrpInf = new(OriginalGroupInformation29)
	return p.OrgnlGrpInf
}

// OrgnlInstrId    *Max35Text
func (p *PaymentTransaction121) SetOrgnlInstrId(value string) {
	p.OrgnlInstrId = (*Max35Text)(&value)
}

// OrgnlEndToEndId *Max35Text
func (p *PaymentTransaction121) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

// OrgnlTxId       *Max35Text
func (p *PaymentTransaction121) SetOrgnlTxId(value string) {
	p.OrgnlTxId = (*Max35Text)(&value)
}

// OrgnlUETR       *UUIDv4Identifier
func (p *PaymentTransaction121) SetOrgnlUETR(value string) {
	p.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

// AccptncDtTm     *ISODateTime
func (p *PaymentTransaction121) SetAccptncDtTm(value string) {
	p.AccptncDtTm = (*ISODateTime)(&value)
}

// ClrSysRef       *Max35Text
func (p *PaymentTransaction121) SetClrSysRef(value string) {
	p.ClrSysRef = (*Max35Text)(&value)
}

// InstgAgt        *BranchAndFinancialInstitutionIdentification6
func (p *PaymentTransaction121) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	p.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return p.InstgAgt
}

// InstdAgt        *BranchAndFinancialInstitutionIdentification6
func (p *PaymentTransaction121) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	p.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return p.InstdAgt
}

// OrgnlTxRef      *OriginalTransactionReference31
func (p *PaymentTransaction121) AddOrgnlTxRef() *OriginalTransactionReference31 {
	p.OrgnlTxRef = new(OriginalTransactionReference31)
	return p.OrgnlTxRef
}

// SplmtryData     []*SupplementaryData1
func (p *PaymentTransaction121) AddSplmtryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SplmtryData = append(p.SplmtryData, newValue)
	return newValue
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:"IntrBkSttlmAmt,omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty" json:"Amt,omitempty"`
	IntrBkSttlmDt  *ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   *ISODate                                      `xml:"ReqdColltnDt,omitempty" json:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:"CdtrSchmeId,omitempty"`
	SttlmInf       *SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:"SttlmInf,omitempty"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:"PmtMtd,omitempty"`
	MndtRltdInf    *MandateRelatedData1Choice                    `xml:"MndtRltdInf,omitempty" json:"MndtRltdInf,omitempty"`
	RmtInf         *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	DbtrAcct       *CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct    *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct    *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr           *Party40Choice                                `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	CdtrAcct       *CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	Purp           *Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
}

func (o *OriginalTransactionReference31) AddIntrBkSttlmAmt() *ActiveOrHistoricCurrencyAndAmount {
	o.IntrBkSttlmAmt = new(ActiveOrHistoricCurrencyAndAmount)
	return o.IntrBkSttlmAmt
}

func (o *OriginalTransactionReference31) AddAmt() *AmountType4Choice {
	o.Amt = new(AmountType4Choice)
	return o.Amt
}

func (o *OriginalTransactionReference31) SetIntrBkSttlmDt(value string) {
	o.IntrBkSttlmDt = (*ISODate)(&value)
}

func (o *OriginalTransactionReference31) SetReqdColltnDt(value string) {
	o.ReqdColltnDt = (*ISODate)(&value)
}

func (o *OriginalTransactionReference31) AddReqdExctnDt() *DateAndDateTime2Choice {
	o.ReqdExctnDt = new(DateAndDateTime2Choice)
	return o.ReqdExctnDt
}

func (o *OriginalTransactionReference31) AddCdtrSchmeId() *PartyIdentification135 {
	o.CdtrSchmeId = new(PartyIdentification135)
	return o.CdtrSchmeId
}

func (o *OriginalTransactionReference31) AddSttlmInf() *SettlementInstruction7 {
	o.SttlmInf = new(SettlementInstruction7)
	return o.SttlmInf
}

func (o *OriginalTransactionReference31) AddPmtTpInf() *PaymentTypeInformation27 {
	o.PmtTpInf = new(PaymentTypeInformation27)
	return o.PmtTpInf
}

func (o *OriginalTransactionReference31) SetPmtMtd(value string) {
	o.PmtMtd = (*PaymentMethod4Code)(&value)
}

func (o *OriginalTransactionReference31) AddMndtRltdInf() *MandateRelatedData1Choice {
	o.MndtRltdInf = new(MandateRelatedData1Choice)
	return o.MndtRltdInf
}

func (o *OriginalTransactionReference31) AddRmtInf() *RemittanceInformation16 {
	o.RmtInf = new(RemittanceInformation16)
	return o.RmtInf
}

func (o *OriginalTransactionReference31) AddUltmtDbtr() *Party40Choice {
	o.UltmtDbtr = new(Party40Choice)
	return o.UltmtDbtr
}

func (o *OriginalTransactionReference31) AddDbtr() *Party40Choice {
	o.Dbtr = new(Party40Choice)
	return o.Dbtr
}

func (o *OriginalTransactionReference31) AddDbtrAcct() *CashAccount38 {
	o.DbtrAcct = new(CashAccount38)
	return o.DbtrAcct
}

func (o *OriginalTransactionReference31) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	o.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return o.DbtrAgt
}

func (o *OriginalTransactionReference31) AddDbtrAgtAcct() *CashAccount38 {
	o.DbtrAgtAcct = new(CashAccount38)
	return o.DbtrAgtAcct
}

func (o *OriginalTransactionReference31) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	o.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return o.CdtrAgt
}

func (o *OriginalTransactionReference31) AddCdtrAgtAcct() *CashAccount38 {
	o.CdtrAgtAcct = new(CashAccount38)
	return o.CdtrAgtAcct
}

func (o *OriginalTransactionReference31) AddCdtr() *Party40Choice {
	o.Cdtr = new(Party40Choice)
	return o.Cdtr
}

func (o *OriginalTransactionReference31) AddCdtrAcct() *CashAccount38 {
	o.CdtrAcct = new(CashAccount38)
	return o.CdtrAcct
}

func (o *OriginalTransactionReference31) AddUltmtCdtr() *Party40Choice {
	o.UltmtCdtr = new(Party40Choice)
	return o.UltmtCdtr
}

func (o *OriginalTransactionReference31) AddPurp() *Purpose2Choice {
	o.Purp = new(Purpose2Choice)
	return o.Purp
}

type MandateRelatedData1Choice struct {
	DrctDbtMndt *MandateRelatedInformation14 `xml:"DrctDbtMndt,omitempty" json:"DrctDbtMndt,omitempty"`
	CdtTrfMndt  *CreditTransferMandateData1  `xml:"CdtTrfMndt,omitempty" json:"CdtTrfMndt,omitempty"`
}

func (m *MandateRelatedData1Choice) AddDrctDbtMndt() *MandateRelatedInformation14 {
	m.DrctDbtMndt = new(MandateRelatedInformation14)
	return m.DrctDbtMndt
}

func (m *MandateRelatedData1Choice) AddCdtTrfMndt() *CreditTransferMandateData1 {
	m.CdtTrfMndt = new(CreditTransferMandateData1)
	return m.CdtTrfMndt
}

type CreditTransferMandateData1 struct {
	MndtId       *Max35Text                 `xml:"MndtId,omitempty" json:"MndtId,omitempty"`
	Tp           *MandateTypeInformation2   `xml:"Tp,omitempty" json:"Tp,omitempty"`
	DtOfSgntr    *ISODate                   `xml:"DtOfSgntr,omitempty" json:"DtOfSgntr,omitempty"`
	DtOfVrfctn   *ISODateTime               `xml:"DtOfVrfctn,omitempty" json:"DtOfVrfctn,omitempty"`
	ElctrncSgntr *Max10KBinary              `xml:"ElctrncSgntr,omitempty" json:"ElctrncSgntr,omitempty"`
	FrstPmtDt    *ISODate                   `xml:"FrstPmtDt,omitempty" json:"FrstPmtDt,omitempty"`
	FnlPmtDt     *ISODate                   `xml:"FnlPmtDt,omitempty" json:"FnlPmtDt,omitempty"`
	Frqcy        *Frequency36Choice         `xml:"Frqcy,omitempty" json:"Frqcy,omitempty"`
	Rsn          *MandateSetupReason1Choice `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
}

func (c *CreditTransferMandateData1) SetMndtId(value string) {
	c.MndtId = (*Max35Text)(&value)
}

func (c *CreditTransferMandateData1) AddTp() *MandateTypeInformation2 {
	c.Tp = new(MandateTypeInformation2)
	return c.Tp
}

func (c *CreditTransferMandateData1) SetDtOfSgntr(value string) {
	c.DtOfSgntr = (*ISODate)(&value)
}

func (c *CreditTransferMandateData1) SetDtOfVrfctn(value string) {
	c.DtOfVrfctn = (*ISODateTime)(&value)
}

func (c *CreditTransferMandateData1) SetElctrncSgntr(value string) {
	c.ElctrncSgntr = (*Max10KBinary)(&value)
}

func (c *CreditTransferMandateData1) SetFrstPmtDt(value string) {
	c.FrstPmtDt = (*ISODate)(&value)
}

func (c *CreditTransferMandateData1) SetFnlPmtDt(value string) {
	c.FnlPmtDt = (*ISODate)(&value)
}

func (c *CreditTransferMandateData1) AddFrqcy() *Frequency36Choice {
	c.Frqcy = new(Frequency36Choice)
	return c.Frqcy
}

func (c *CreditTransferMandateData1) AddRsn() *MandateSetupReason1Choice {
	c.Rsn = new(MandateSetupReason1Choice)
	return c.Rsn
}

type MandateTypeInformation2 struct {
	SvcLvl    *ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
	Clssfctn  *MandateClassification1Choice `xml:"Clssfctn,omitempty" json:"Clssfctn,omitempty"`
}

func (m *MandateTypeInformation2) AddSvcLvl() *ServiceLevel8Choice {
	m.SvcLvl = new(ServiceLevel8Choice)
	return m.SvcLvl
}

func (m *MandateTypeInformation2) AddLclInstrm() *LocalInstrument2Choice {
	m.LclInstrm = new(LocalInstrument2Choice)
	return m.LclInstrm
}

func (m *MandateTypeInformation2) AddCtgyPurp() *CategoryPurpose1Choice {
	m.CtgyPurp = new(CategoryPurpose1Choice)
	return m.CtgyPurp
}

func (m *MandateTypeInformation2) AddClssfctn() *MandateClassification1Choice {
	m.Clssfctn = new(MandateClassification1Choice)
	return m.Clssfctn
}

type MandateClassification1Choice struct {
	Cd    *MandateClassification1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                  `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

func (m *MandateClassification1Choice) SetCd(value string) {
	m.Cd = (*MandateClassification1Code)(&value)
}

func (m *MandateClassification1Choice) SetPrtry(value string) {
	m.Prtry = (*Max35Text)(&value)
}

// May be one of FIXE, USGB, VARI
type MandateClassification1Code string

// TODO: prxy.001.001.01, json tag

type ProxyRegistration1 struct {
	RegnTp    *ProxyRegistrationType1Code    `xml:"RegnTp,omitempty" json:"RegnTp,omitempty"`
	RegnSubTp *ProxyRegistrationSubType1Code `xml:"RegnSubTp,omitempty" json:"RegnSubTp,omitempty"`
	Prxy      *ProxyDefinition1              `xml:"Prxy,omitempty" json:"Prxy,omitempty"`
	PrxyRegn  *ProxyRegistrationAccount1     `xml:"PrxyRegn,omitempty" json:"PrxyRegn,omitempty"`
}

func (p *ProxyRegistration1) SetRegnTp(value string) {
	p.RegnTp = (*ProxyRegistrationType1Code)(&value)
}

func (p *ProxyRegistration1) SetRegnSubTp(value string) {
	p.RegnSubTp = (*ProxyRegistrationSubType1Code)(&value)
}

func (p *ProxyRegistration1) AddPrxy() *ProxyDefinition1 {
	p.Prxy = new(ProxyDefinition1)
	return p.Prxy
}

func (p *ProxyRegistration1) AddPrxyRegn() *ProxyRegistrationAccount1 {
	p.PrxyRegn = new(ProxyRegistrationAccount1)
	return p.PrxyRegn
}

// May be one of AMND, DEAC, NEWR, SUSP, ACTV, PORT
type ProxyRegistrationType1Code string

// May be no more than 4 items long
type ProxyRegistrationSubType1Code string

type ProxyDefinition1 struct {
	Tp  *Max12Text  `xml:"Tp" json:"Tp"`
	Val *Max140Text `xml:"Val" json:"Val"`
}

func (p *ProxyDefinition1) SetTp(value string) {
	p.Tp = (*Max12Text)(&value)
}

func (p *ProxyDefinition1) SetVal(value string) {
	p.Val = (*Max140Text)(&value)
}

type ProxyRegistrationAccount1 struct {
	RegnId      *Max35Text                                    `xml:"RegnId,omitempty" json:"RegnId,omitempty"`
	DsplNm      *Max140Text                                   `xml:"DsplNm,omitempty" json:"DsplNm,omitempty"`
	Agt         *BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty" json:"Agt,omitempty"`
	Acct        *CashAccount40                                `xml:"Acct,omitempty" json:"Acct,omitempty"`
	AcctHldr    *Party30Choice                                `xml:"AcctHldr,omitempty" json:"AcctHldr,omitempty"`
	ScndId      *ScndIdDefinition1                            `xml:"ScndId,omitempty" json:"ScndId,omitempty"`
	RegnSts     *ProxyRegistrationStatusCode                  `xml:"RegnSts,omitempty" json:"RegnSts,omitempty"`
	PreAuthrsd  *TrueFalseIndicator                           `xml:"PreAuthrsd,omitempty" json:"PreAuthrsd,omitempty"`
	SplmtryData []*BI_SupplementaryData1                      `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

func (p *ProxyRegistrationAccount1) SetRegnId(value string) {
	p.RegnId = (*Max35Text)(&value)
}

func (p *ProxyRegistrationAccount1) SetDsplNm(value string) {
	p.DsplNm = (*Max140Text)(&value)
}

func (p *ProxyRegistrationAccount1) AddAgt() *BranchAndFinancialInstitutionIdentification5 {
	p.Agt = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agt
}

func (p *ProxyRegistrationAccount1) AddAcct() *CashAccount40 {
	p.Acct = new(CashAccount40)
	return p.Acct
}

func (p *ProxyRegistrationAccount1) AddAcctHldr() *Party30Choice {
	p.AcctHldr = new(Party30Choice)
	return p.AcctHldr
}

func (p *ProxyRegistrationAccount1) AddScndId() *ScndIdDefinition1 {
	p.ScndId = new(ScndIdDefinition1)
	return p.ScndId
}

func (p *ProxyRegistrationAccount1) SetRegnSts(value string) {
	p.RegnSts = (*ProxyRegistrationStatusCode)(&value)
}

func (p *ProxyRegistrationAccount1) SetPreAuthrsd(value string) {
	p.PreAuthrsd = (*TrueFalseIndicator)(&value)
}

func (p *ProxyRegistrationAccount1) AddSplmtryData() *BI_SupplementaryData1 {
	newValue := new(BI_SupplementaryData1)
	p.SplmtryData = append(p.SplmtryData, newValue)
	return newValue
}

type CashAccount40 struct {
	Id         *AccountIdentification4Choice `xml:"Id" json:"Id"`
	Tp         *CashAccountType2ChoiceProxy  `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Nm         *Max140Text                   `xml:"Nm,omitempty" json:"Nm,omitempty"`
	AcctHldrTp *Max35Text                    `xml:"AcctHldrTp,omitempty" json:"AcctHldrTp,omitempty"`
}

func (c *CashAccount40) AddId() *AccountIdentification4Choice {
	c.Id = new(AccountIdentification4Choice)
	return c.Id
}

func (c *CashAccount40) AddTp() *CashAccountType2ChoiceProxy {
	c.Tp = new(CashAccountType2ChoiceProxy)
	return c.Tp
}

func (c *CashAccount40) SetNm(value string) {
	c.Nm = (*Max140Text)(&value)
}

func (c *CashAccount40) SetAcctHldrTp(value string) {
	c.AcctHldrTp = (*Max35Text)(&value)
}

type CashAccountType2ChoiceProxy struct {
	Prtry *ProxyAccountType `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

func (c *CashAccountType2ChoiceProxy) SetPrtry(value string) {
	c.Prtry = (*ProxyAccountType)(&value)
}

// May be no more than 4 items long
type ProxyAccountType string

type ScndIdDefinition1 struct {
	Tp  *Max12Text `xml:"Tp" json:"Tp"`
	Val *Max35Text `xml:"Val" json:"Val"`
}

func (s *ScndIdDefinition1) SetTp(value string) {
	s.Tp = (*Max12Text)(&value)
}
func (s *ScndIdDefinition1) SetVal(value string) {
	s.Val = (*Max35Text)(&value)
}

// May be no more than 4 items long
type ProxyRegistrationStatusCode string

// TODO: prxy.002.001.01, json tag

type ProxyRegistrationResponse1 struct {
	PrxRspnSts     *ProxyStatusCode               `xml:"PrxRspnSts" json:"PrxRspnSts"`
	StsRsnInf      *ProxyStatusChoice             `xml:"StsRsnInf,omitempty" json:"StsRsnInf,omitempty"`
	OrgnlRegnTp    *ProxyRegistrationType1Code    `xml:"OrgnlRegnTp" json:"OrgnlRegnTp"`
	OrgnlRegnSubTp *ProxyRegistrationSubType1Code `xml:"OrgnlRegnSubTp,omitempty" json:"OrgnlRegnSubTp,omitempty"`
	OrgnlPrxy      *ProxyDefinition1              `xml:"OrgnlPrxy,omitempty" json:"OrgnlPrxy,omitempty"`
	PrxyRegn       []*ProxyRegistrationAccount1   `xml:"PrxyRegn,omitempty" json:"PrxyRegn,omitempty"`
}

func (p *ProxyRegistrationResponse1) SetPrxRspnSts(value string) {
	p.PrxRspnSts = (*ProxyStatusCode)(&value)
}

func (p *ProxyRegistrationResponse1) AddStsRsnInf() *ProxyStatusChoice {
	p.StsRsnInf = new(ProxyStatusChoice)
	return p.StsRsnInf
}

func (p *ProxyRegistrationResponse1) SetOrgnlRegnTp(value string) {
	p.OrgnlRegnTp = (*ProxyRegistrationType1Code)(&value)
}

func (p *ProxyRegistrationResponse1) SetOrgnlRegnSubTp(value string) {
	p.OrgnlRegnSubTp = (*ProxyRegistrationSubType1Code)(&value)
}

func (p *ProxyRegistrationResponse1) AddOrgnlPrxy() *ProxyDefinition1 {
	p.OrgnlPrxy = new(ProxyDefinition1)
	return p.OrgnlPrxy
}

func (p *ProxyRegistrationResponse1) AddPrxyRegn() *ProxyRegistrationAccount1 {
	newValue := new(ProxyRegistrationAccount1)
	p.PrxyRegn = append(p.PrxyRegn, newValue)
	return newValue
}

// May be one of ACTC, RJCT
type ProxyStatusCode string

type ProxyStatusChoice struct {
	Cd    *ExternalStatusReason1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                 `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

func (p *ProxyStatusChoice) SetCd(value string) {
	p.Cd = (*ExternalStatusReason1Code)(&value)
}

func (p *ProxyStatusChoice) SetPrtry(value string) {
	p.Prtry = (*Max35Text)(&value)
}

// TODO: prxy.003.001.01, json tag
type ProxyLookUpChoice1 struct {
	PrxyOnly *ProxyLookUp11 `xml:"PrxyOnly" json:"PrxyOnly"`
}

func (p *ProxyLookUpChoice1) AddPrxyOnly() *ProxyLookUp11 {
	p.PrxyOnly = new(ProxyLookUp11)
	return p.PrxyOnly
}

type ProxyLookUp11 struct {
	LkUpTp    *ProxyLookUpType1Code `xml:"LkUpTp" json:"LkUpTP"`
	Id        *Max35Text            `xml:"Id" json:"Id"`
	PrxyRtrvl *ProxyDefinition1     `xml:"PrxyRtrvl" json:"PrxyRtrvl"`
}

func (p *ProxyLookUp11) SetLkUpTp(value string) {
	p.LkUpTp = (*ProxyLookUpType1Code)(&value)
}

func (p *ProxyLookUp11) SetId(value string) {
	p.Id = (*Max35Text)(&value)
}

func (p *ProxyLookUp11) AddPrxyRtrvl() *ProxyDefinition1 {
	p.PrxyRtrvl = new(ProxyDefinition1)
	return p.PrxyRtrvl
}

type ProxyLookUpType1Code string

// TODO: prxy.004.001.01, json tag
type ProxyLookUpResponse1 struct {
	OrgnlId        *Max35Text                `xml:"OrgnlId" json:"OrgnlId"`
	OrgnlPrxyRtrvl *ProxyDefinition1         `xml:"OrgnlPrxyRtrvl" json:"OrgnlPrxyRtrvl"`
	RegnRspn       *ProxyLookUpRegistration1 `xml:"RegnRspn" json:"RegnRspn"`
}

func (p *ProxyLookUpResponse1) SetOrgnlId(value string) {
	p.OrgnlId = (*Max35Text)(&value)
}

func (p *ProxyLookUpResponse1) AddOrgnlPrxyRtrvl() *ProxyDefinition1 {
	p.OrgnlPrxyRtrvl = new(ProxyDefinition1)
	return p.OrgnlPrxyRtrvl
}

func (p *ProxyLookUpResponse1) AddRegnRspn() *ProxyLookUpRegistration1 {
	p.RegnRspn = new(ProxyLookUpRegistration1)
	return p.RegnRspn
}

type ProxyLookUpRegistration1 struct {
	PrxRspnSts *ProxyStatusCode     `xml:"PrxRspnSts" json:"PrxRspnSts"`
	StsRsnInf  *ProxyStatusChoice   `xml:"StsRsnInf" json:"StsRsnInf"`
	Prxy       *ProxyDefinition1    `xml:"Prxy" json:"Prxy"`
	Regn       *ProxyLookUpAccount1 `xml:"Regn" json:"Regn"`
}

func (p *ProxyLookUpRegistration1) SetPrxRspnSts(value string) {
	p.PrxRspnSts = (*ProxyStatusCode)(&value)
}

func (p *ProxyLookUpRegistration1) AddStsRsnInf() *ProxyStatusChoice {
	p.StsRsnInf = new(ProxyStatusChoice)
	return p.StsRsnInf
}

func (p *ProxyLookUpRegistration1) AddPrxy() *ProxyDefinition1 {
	p.Prxy = new(ProxyDefinition1)
	return p.Prxy
}

func (p *ProxyLookUpRegistration1) AddRegn() *ProxyLookUpAccount1 {
	p.Regn = new(ProxyLookUpAccount1)
	return p.Regn
}

type ProxyLookUpAccount1 struct {
	RegnId *Max35Text                                    `xml:"RegnId" json:"RegnId"`
	DsplNm *Max140Text                                   `xml:"DsplNm" json:"DsplNm"`
	Agt    *BranchAndFinancialInstitutionIdentification5 `xml:"Agt" json:"Agt"`
	Acct   *CashAccount40                                `xml:"Acct" json:"Acct"`
}

func (p *ProxyLookUpAccount1) SetRegnId(value string) {
	p.RegnId = (*Max35Text)(&value)
}

func (p *ProxyLookUpAccount1) SetDsplNm(value string) {
	p.DsplNm = (*Max140Text)(&value)
}

func (p *ProxyLookUpAccount1) AddAgt() *BranchAndFinancialInstitutionIdentification5 {
	p.Agt = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agt
}

func (p *ProxyLookUpAccount1) AddAcct() *CashAccount40 {
	p.Acct = new(CashAccount40)
	return p.Acct
}

// TODO: prxy.005.001.01, json tag
type ProxyEnquiryChoice1 struct {
	RegnId *Max35Text         `xml:"RegnId" json:"RegnId"`
	ScndId *ScndIdDefinition1 `xml:"ScndId" json:"ScndId"`
}

func (p *ProxyEnquiryChoice1) SetRegnId(value string) {
	p.RegnId = (*Max35Text)(&value)
}

func (p *ProxyEnquiryChoice1) AddScndId() *ScndIdDefinition1 {
	p.ScndId = new(ScndIdDefinition1)
	return p.ScndId
}

// TODO: prxy.006.001.01, json tag
type ProxyEnquiryResponse1 struct {
	PrxRspnSts *ProxyStatusCode          `xml:"PrxRspnSts" json:"PrxRspnSts"`
	StsRsnInf  *ProxyStatusChoice        `xml:"StsRsnInf" json:"StsRsnInf"`
	Rspn       *ProxyEnquiryInformation1 `xml:"Rspn" json:"Rspn"`
}

func (p *ProxyEnquiryResponse1) SetPrxRspnSts(value string) {
	p.PrxRspnSts = (*ProxyStatusCode)(&value)
}

func (p *ProxyEnquiryResponse1) AddStsRsnInf() *ProxyStatusChoice {
	p.StsRsnInf = new(ProxyStatusChoice)
	return p.StsRsnInf
}

func (p *ProxyEnquiryResponse1) AddRspn() *ProxyEnquiryInformation1 {
	p.Rspn = new(ProxyEnquiryInformation1)
	return p.Rspn
}

type ProxyEnquiryInformation1 struct {
	RegnId      *Max35Text                                    `xml:"RegnId" json:"RegnId"`
	DsplNm      *Max140Text                                   `xml:"DsplNm" json:"DsplNm"`
	Ptcpt       *BranchAndFinancialInstitutionIdentification5 `xml:"Ptcpt" json:"Ptcpt"`
	PrxyInf     *ProxyEnquiryDefinition1                      `xml:"PrxyInf" json:"PrxyInf"`
	AcctInf     *ProxyEnquiryAccount1                         `xml:"AcctInf" json:"AcctInf"`
	SplmtryData *BI_SupplementaryData1                        `xml:"SplmtryData" json:"SplmtryData"`
}

func (p *ProxyEnquiryInformation1) SetRegnId(value string) {
	p.RegnId = (*Max35Text)(&value)
}
func (p *ProxyEnquiryInformation1) SetDsplNm(value string) {
	p.DsplNm = (*Max140Text)(&value)
}
func (p *ProxyEnquiryInformation1) AddPtcpt() *BranchAndFinancialInstitutionIdentification5 {
	p.Ptcpt = new(BranchAndFinancialInstitutionIdentification5)
	return p.Ptcpt
}
func (p *ProxyEnquiryInformation1) AddPrxyInf() *ProxyEnquiryDefinition1 {
	p.PrxyInf = new(ProxyEnquiryDefinition1)
	return p.PrxyInf
}
func (p *ProxyEnquiryInformation1) AddAcctInf() *ProxyEnquiryAccount1 {
	p.AcctInf = new(ProxyEnquiryAccount1)
	return p.AcctInf
}
func (p *ProxyEnquiryInformation1) AddSplmtryData() *BI_SupplementaryData1 {
	p.SplmtryData = new(BI_SupplementaryData1)
	return p.SplmtryData
}

type ProxyEnquiryDefinition1 struct {
	Tp  *Max12Text              `xml:"Tp" json:"Tp"`
	Val *Max140Text             `xml:"Val" json:"Val"`
	Sts *ProxyEnquiryStatusCode `xml:"Sts" json:"Sts"`
}

func (p *ProxyEnquiryDefinition1) SetTp(value string) {
	p.Tp = (*Max12Text)(&value)
}

func (p *ProxyEnquiryDefinition1) SetVal(value string) {
	p.Val = (*Max140Text)(&value)
}

func (p *ProxyEnquiryDefinition1) SetSts(value string) {
	p.Sts = (*ProxyEnquiryStatusCode)(&value)
}

type ProxyEnquiryAccount1 struct {
	Agt *BranchAndFinancialInstitutionIdentification5 `xml:"Agt" json:"Agt"`
}

func (p *ProxyEnquiryAccount1) AddAgt() *BranchAndFinancialInstitutionIdentification5 {
	p.Agt = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agt
}

type ProxyEnquiryStatusCode string

// TODO: prxy.901.001.01, json tag
type ProxyNtfctn1 struct {
	OrgnlId   *Max35Text        `xml:"OrgnlId" json:"OrgnlId"`
	OrgnlPrxy *ProxyDefinition1 `xml:"OrgnlPrxy" json:"OrgnlPrxy"`
	OrgnlAcct *ProxyAccount1    `xml:"OrgnlAcct" json:"OrgnlAcct"`
	NewAcct   *ProxyAccount1    `xml:"NewAcct" json:"NewAcct"`
}

func (p *ProxyNtfctn1) SetOrgnlId(value string) {
	p.OrgnlId = (*Max35Text)(&value)
}

func (p *ProxyNtfctn1) AddOrgnlPrxy() *ProxyDefinition1 {
	p.OrgnlPrxy = new(ProxyDefinition1)
	return p.OrgnlPrxy
}

func (p *ProxyNtfctn1) AddOrgnlAcct() *ProxyAccount1 {
	p.OrgnlAcct = new(ProxyAccount1)
	return p.OrgnlAcct
}

func (p *ProxyNtfctn1) AddNewAcct() *ProxyAccount1 {
	p.NewAcct = new(ProxyAccount1)
	return p.NewAcct
}

type ProxyAccount1 struct {
	RegnId      *Max35Text                                    `xml:"RegnId" json:"RegnId"`
	DsplNm      *Max140Text                                   `xml:"DsplNm" json:"DsplNm"`
	Agt         *BranchAndFinancialInstitutionIdentification5 `xml:"Agt" json:"Agt"`
	Acct        *CashAccount40                                `xml:"Acct" json:"Acct"`
	SplmtryData *BI_SupplementaryData1                        `xml:"SplmtryData" json:"SplmtryData"`
}

func (p *ProxyAccount1) SetRegnId(value string) {
	p.RegnId = (*Max35Text)(&value)
}

func (p *ProxyAccount1) SetDsplNm(value string) {
	p.DsplNm = (*Max140Text)(&value)
}

func (p *ProxyAccount1) AddAgt() *BranchAndFinancialInstitutionIdentification5 {
	p.Agt = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agt
}
func (p *ProxyAccount1) AddAcct() *CashAccount40 {
	p.Acct = new(CashAccount40)
	return p.Acct
}

func (p *ProxyAccount1) AddSplmtryData() *BI_SupplementaryData1 {
	p.SplmtryData = new(BI_SupplementaryData1)
	return p.SplmtryData
}

type AdminTransactionInformation struct {
	FnctnCd  *Max4NumericText                              `xml:"FnctnCd" json:"FnctnCd"`
	InstrId  *Max35Text                                    `xml:"InstrId" json:"InstrId"`
	InstgAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt" json:"InstgAgt"`
}

func (a *AdminTransactionInformation) SetFnctnCd(value string) {
	a.FnctnCd = (*Max4NumericText)(&value)
}

func (a *AdminTransactionInformation) SetInstrId(value string) {
	a.InstrId = (*Max35Text)(&value)
}

type AdminResponse struct {
	InstgAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt" json:"InstgAgt"`
	OrgnlInstrId *Max35Text                                    `xml:"OrgnlInstrId" json:"OrgnlInstrId"`
	FnctnCd      *Max4NumericText                              `xml:"FnctnCd" json:"FnctnCd"`
	TxSts        *TransactionIndividualStatus3Code             `xml:"TxSts" json:"TxSts"`
}

func (a *AdminResponse) AddInstgAgt() *BranchAndFinancialInstitutionIdentification5 {
	a.InstgAgt = new(BranchAndFinancialInstitutionIdentification5)
	return a.InstgAgt
}

func (a *AdminResponse) SetOrgnlInstrId(value string) {
	a.OrgnlInstrId = (*Max35Text)(&value)
}

func (a *AdminResponse) SetFnctnCd(value string) {
	a.FnctnCd = (*Max4NumericText)(&value)
}

func (a *AdminResponse) SetTxSts(value string) {
	a.TxSts = (*TransactionIndividualStatus3Code)(&value)
}
