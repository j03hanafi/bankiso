package camt

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022"
)

type Document05900104 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.04 Document"`
	Message *NotificationToReceiveStatusReportV04 `xml:"NtfctnToRcvStsRpt"`
}

func (d *Document05900104) AddMessage() *NotificationToReceiveStatusReportV04 {
	d.Message = new(NotificationToReceiveStatusReportV04)
	return d.Message
}

// Scope
// The NotificationToReceiveStatusReport message is sent by an account servicing institution to an account owner or to a party acting on the account owner's behalf. It is used to notify the account owner about the status of one or more expected payments that were advised in a previous NotificationToReceive message.
// Usage
// The NotificationToReceiveStatusReport message is sent in response to a NotificationToReceive message and can be used in either a direct or a relay scenario. It is used to advise the account owner of receipt of the funds as expected. It is also used to notify the account owner of non-receipt of funds or of discrepancies in the funds received.
type NotificationToReceiveStatusReportV04 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *iso20022.GroupHeader60 `xml:"GrpHdr"`

	// Set of elements used to identify the original notification and to provide the status.
	OriginalNotificationAndStatus *iso20022.OriginalNotification7 `xml:"OrgnlNtfctnAndSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NotificationToReceiveStatusReportV04) AddGroupHeader() *iso20022.GroupHeader60 {
	n.GroupHeader = new(iso20022.GroupHeader60)
	return n.GroupHeader
}

func (n *NotificationToReceiveStatusReportV04) AddOriginalNotificationAndStatus() *iso20022.OriginalNotification7 {
	n.OriginalNotificationAndStatus = new(iso20022.OriginalNotification7)
	return n.OriginalNotificationAndStatus
}

func (n *NotificationToReceiveStatusReportV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}
func (d *Document05900104) String() (result string, ok bool) { return }
