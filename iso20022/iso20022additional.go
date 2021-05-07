package iso20022

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
//func (s *SettlementInstruction4) SetSettlementMethod(value string) {
//	s.SettlementMethod = (*SettlementMethod1Code)(&value)
//}
//
//func (s *SettlementInstruction4) AddSettlementAccount() *CashAccount38 {
//	s.SettlementAccount = new(CashAccount38)
//	return s.SettlementAccount
//}
//
//func (s *SettlementInstruction4) AddClearingSystem() *ClearingSystemIdentification3Choice {
//	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
//	return s.ClearingSystem
//}
//
//func (s *SettlementInstruction4) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification6 {
//	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification6)
//	return s.InstructingReimbursementAgent
//}
//
//func (s *SettlementInstruction4) AddInstructingReimbursementAgentAccount() *CashAccount38 {
//	s.InstructingReimbursementAgentAccount = new(CashAccount38)
//	return s.InstructingReimbursementAgentAccount
//}
//
//func (s *SettlementInstruction4) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification6 {
//	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification6)
//	return s.InstructedReimbursementAgent
//}
//
//func (s *SettlementInstruction4) AddInstructedReimbursementAgentAccount() *CashAccount38 {
//	s.InstructedReimbursementAgentAccount = new(CashAccount38)
//	return s.InstructedReimbursementAgentAccount
//}
//
//func (s *SettlementInstruction4) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification6 {
//	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification6)
//	return s.ThirdReimbursementAgent
//}
//
//func (s *SettlementInstruction4) AddThirdReimbursementAgentAccount() *CashAccount38 {
//	s.ThirdReimbursementAgentAccount = new(CashAccount38)
//	return s.ThirdReimbursementAgentAccount
//}
