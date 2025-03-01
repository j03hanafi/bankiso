package seev

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document00800104 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.04 Document"`
	Message *MeetingResultDisseminationV04 `xml:"MtgRsltDssmntn"`
}

func (d *Document00800104) AddMessage() *MeetingResultDisseminationV04 {
	d.Message = new(MeetingResultDisseminationV04)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingResultDissemination message to another intermediary, to a party holding the right to vote, to a registered security holder or to a beneficial holder to provide information on the voting results of a shareholders meeting.
// Usage
// The MeetingResultDissemination message is used to provide the vote results per resolution. It may also provide information on the level of participation.
// This message is also used to notify an update or amendment to a previously sent MeetingResultDissemination message.
type MeetingResultDisseminationV04 struct {

	// Identifies the meeting dissemination notification message.
	Identification *iso20022.MessageIdentification1 `xml:"Id"`

	// Information specific to an amemdment.
	Amendment *iso20022.AmendInformation2 `xml:"Amdmnt,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference4 `xml:"MtgRef"`

	// Party reporting the meeting results.
	ReportingParty *iso20022.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	Security []*iso20022.SecurityPosition6 `xml:"Scty"`

	// Results per resolution.
	VoteResult []*iso20022.Vote5 `xml:"VoteRslt"`

	// Information about the participation to the voting process.
	Participation *iso20022.Participation3 `xml:"Prtcptn,omitempty"`

	// Information on where additionnal information can be received.
	AdditionalInformation *iso20022.CommunicationAddress4 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingResultDisseminationV04) AddIdentification() *iso20022.MessageIdentification1 {
	m.Identification = new(iso20022.MessageIdentification1)
	return m.Identification
}

func (m *MeetingResultDisseminationV04) AddAmendment() *iso20022.AmendInformation2 {
	m.Amendment = new(iso20022.AmendInformation2)
	return m.Amendment
}

func (m *MeetingResultDisseminationV04) AddMeetingReference() *iso20022.MeetingReference4 {
	m.MeetingReference = new(iso20022.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingResultDisseminationV04) AddReportingParty() *iso20022.PartyIdentification9Choice {
	m.ReportingParty = new(iso20022.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingResultDisseminationV04) AddSecurity() *iso20022.SecurityPosition6 {
	newValue := new(iso20022.SecurityPosition6)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV04) AddVoteResult() *iso20022.Vote5 {
	newValue := new(iso20022.Vote5)
	m.VoteResult = append(m.VoteResult, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV04) AddParticipation() *iso20022.Participation3 {
	m.Participation = new(iso20022.Participation3)
	return m.Participation
}

func (m *MeetingResultDisseminationV04) AddAdditionalInformation() *iso20022.CommunicationAddress4 {
	m.AdditionalInformation = new(iso20022.CommunicationAddress4)
	return m.AdditionalInformation
}

func (m *MeetingResultDisseminationV04) AddExtension() *iso20022.Extension2 {
	newValue := new(iso20022.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
func (d *Document00800104) String() (result string, ok bool) { return }
