package cli

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/splorg/urbandict/internal/command"
	"github.com/splorg/urbandict/internal/data"
	"github.com/splorg/urbandict/internal/message"
)

type model struct {
	title     string
	textinput textinput.Model
	terms     data.Terms
	err       error
}

func Start() {
	m := newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err := p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func newModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter search term"
	ti.Focus()

	return model{
		title:     "Hello world",
		textinput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			v := m.textinput.Value()
			return m, command.HandleQuerySearch(v)
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case message.TermsResponseMsg:
		if msg.Err != nil {
			m.err = msg.Err
		}

		m.terms = msg.Terms
		return m, nil
	}

	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := m.textinput.View() + "\n\n"

	if len(m.terms.List) > 0 {
		s += m.terms.List[0].Definition + "\n\n"
		s += m.terms.List[0].Example + "\n\n"
		s += fmt.Sprintf(
			"Upvotes: %d\nDownvotes: %d\n\n",
			m.terms.List[0].ThumbsUp,
			m.terms.List[0].ThumbsDown,
		)
	}

	return s
}
