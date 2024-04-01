package command

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/splorg/urbandict/internal/data"
	"github.com/splorg/urbandict/internal/message"
)

func HandleQuerySearch(q string) tea.Cmd {
	return func() tea.Msg {
		apiUrl := fmt.Sprintf("https://api.urbandictionary.com/v0/define?term=%s", url.QueryEscape(q))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
		if err != nil {
			return message.TermsResponseMsg{
				Err: err,
			}
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return message.TermsResponseMsg{
				Err: err,
			}
		}

		defer res.Body.Close()

		var terms data.Terms

		err = json.NewDecoder(res.Body).Decode(&terms)
		if err != nil {
			return message.TermsResponseMsg{
				Err: err,
			}
		}

		return message.TermsResponseMsg{
			Terms: terms,
		}
	}
}
