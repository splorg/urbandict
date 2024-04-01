package message

import "github.com/splorg/urbandict/internal/data"

type TermsResponseMsg struct {
	Terms data.Terms
	Err   error
}
