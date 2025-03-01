package camt

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document03100104 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.031.001.04 Document"`
	Message *RejectInvestigationV04 `xml:"RjctInvstgtn"`
}

func (d *Document03100104) AddMessage() *RejectInvestigationV04 {
	d.Message = new(RejectInvestigationV04)
	return d.Message
}

// Scope
// The Reject Investigation message is sent by a case assignee to a case creator or case assigner to reject a case given to him.
// Usage
// The Reject Investigation message is used to notify the case creator or case assigner the rejection of an assignment by the case assignee in a:
// - request to cancel payment case
// - request to modify payment case
// - unable to apply case
// - claim non receipt case
// Rejecting a case assignment occurs when:
// - the case assignee is unable to trace the original payment instruction
// - the case assignee is unable, or does not have authority, to process the assigned case (indicate "You have by-passed a party",
// - the case assignee has received a non expected message, and rejects the message with a wrong message indicator
// - the case assignee has not yet received the Resolution Of Investigation message and the case has already been reopened
// - the case assignee has rejects an non-cash related query
// The Reject Investigation message covers one and only one case at a time. If the case assignee needs to reject several case assignments, then multiple Reject Investigation messages must be sent.
// The Reject Investigation message must be forwarded by all subsequent case assignee(s) until it reaches the case assigner and must not be used in place of a Resolution Of Investigation or Case Status Report message.
type RejectInvestigationV04 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *iso20022.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *iso20022.Case3 `xml:"Case"`

	// Specifies the reason for the rejection of an investigation.
	Justification *iso20022.InvestigationRejectionJustification1 `xml:"Justfn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RejectInvestigationV04) AddAssignment() *iso20022.CaseAssignment3 {
	r.Assignment = new(iso20022.CaseAssignment3)
	return r.Assignment
}

func (r *RejectInvestigationV04) AddCase() *iso20022.Case3 {
	r.Case = new(iso20022.Case3)
	return r.Case
}

func (r *RejectInvestigationV04) AddJustification() *iso20022.InvestigationRejectionJustification1 {
	r.Justification = new(iso20022.InvestigationRejectionJustification1)
	return r.Justification
}

func (r *RejectInvestigationV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
func (d *Document03100104) String() (result string, ok bool) { return }
