package templates

import _ "embed"

//go:embed ticket-action-pack.md
var ticketActionPack string

// RenderTicketActionPack returns the canonical ticket action pack template.
func RenderTicketActionPack() string {
	return ticketActionPack
}
