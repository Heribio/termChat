package cli

import (	
    "fmt"
	"log"
    "net/http"
    "time"
"strings"
    "io"
    "encoding/json"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

    "github.com/Heribio/termChat/internal/discord"
    "github.com/Heribio/termChat/pkg/discordWebhook"
)

func Run() {
	p := tea.NewProgram(initialModel())
    
    go func() {
        for {
            messages := CheckForMessages()
            time.Sleep( 1 * time.Millisecond * 200)
            p.Send(messageList(messages))
        }
    }()

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

type messageList []discordMessage

type discordMessage struct {
    Content string `json:"content"`
    Username string `json:"username"`
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 30)
	vp.SetContent(`Chat with Friends !
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

func (m model) Init() tea.Cmd {
    return tea.Batch(
        textarea.Blink,
        //Refresh(),
    )
}

func CheckForMessages() messageList{
    url := "http://localhost:8080"
    client := &http.Client{
        Timeout: time.Second * 2,
    }

    req, err := client.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    req.Header.Set("Accept", "application/json")
    resp, err := io.ReadAll(req.Body)
    defer req.Body.Close()
    if err != nil {
        fmt.Println(err)
    }
    var messages messageList

    json.Unmarshal(resp, &messages)
    if err != nil {
        fmt.Println(err)
    }
    return messages
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
            discord.SendMessage(discordWebhook.Message{Content: m.textarea.Value(), Username: "Heribio"})
			m.textarea.Reset()
		}
        case messageList:
        m.messages = []string{}
            for _, message := range msg {
                m.messages = append(m.messages, m.senderStyle.Render(message.Username)+ ": " +message.Content)
                m.viewport.SetContent(strings.Join(m.messages, "\n"))
                m.viewport.GotoBottom()
        }
	case errMsg:
		m.err = msg
		return m, nil
	}
	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
