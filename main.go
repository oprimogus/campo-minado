package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	// Configuração para dificuldade "easy"
	rows  = 8
	cols  = 8
	mines = 10
)

// Estilos
var (
	cellStyle = lipgloss.NewStyle().
		Width(3).
		Height(1).
		Align(lipgloss.Center)

	hiddenStyle = cellStyle.Copy().
		Background(lipgloss.Color("240")).
		Foreground(lipgloss.Color("255"))

	revealedStyle = cellStyle.Copy().
		Background(lipgloss.Color("236")).
		Foreground(lipgloss.Color("255"))

	mineStyle = cellStyle.Copy().
		Background(lipgloss.Color("1")).
		Foreground(lipgloss.Color("255"))

	flagStyle = cellStyle.Copy().
		Background(lipgloss.Color("240")).
		Foreground(lipgloss.Color("9"))

	cursorStyle = cellStyle.Copy().
		Background(lipgloss.Color("39")).
		Foreground(lipgloss.Color("255"))

	infoStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("170"))

	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("99"))

	footerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))
)

// KeyMap define os atalhos de teclado para o jogo
type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Reveal key.Binding
	Flag   key.Binding
	Quit   key.Binding
}

// Definindo os atalhos de teclado
var keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "mover para cima"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "mover para baixo"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "mover para esquerda"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "mover para direita"),
	),
	Reveal: key.NewBinding(
		key.WithKeys("enter", "r", " "),
		key.WithHelp("enter/r/space", "revelar célula"),
	),
	Flag: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "marcar/desmarcar bandeira"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "sair"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Reveal, k.Flag, k.Up, k.Down, k.Left, k.Right, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.Reveal, k.Flag, k.Quit},
	}
}

// Cell representa uma célula no tabuleiro
type Cell struct {
	hasMine       bool
	isRevealed    bool
	isFlagged     bool
	adjacentMines int
}

// Model contém o estado do jogo
type Model struct {
	board       [][]Cell
	cursorRow   int
	cursorCol   int
	minesLeft   int
	cellsToReveal int
	isGameOver  bool
	isWin       bool
	message     string
	keymap      KeyMap
	help        help.Model
}

func initialModel() Model {
	rand.Seed(time.Now().UnixNano())
	
	// Criar tabuleiro vazio
	board := make([][]Cell, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]Cell, cols)
	}
	
	// Colocar minas aleatoriamente
	minesPlaced := 0
	for minesPlaced < mines {
		row := rand.Intn(rows)
		col := rand.Intn(cols)
		
		if !board[row][col].hasMine {
			board[row][col].hasMine = true
			minesPlaced++
		}
	}
	
	// Calcular minas adjacentes
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !board[r][c].hasMine {
				count := 0
				// Verificar as 8 células adjacentes
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						nr, nc := r+dr, c+dc
						if nr >= 0 && nr < rows && nc >= 0 && nc < cols && board[nr][nc].hasMine {
							count++
						}
					}
				}
				board[r][c].adjacentMines = count
			}
		}
	}
	
	help := help.New()
	help.Width = 50
	
	return Model{
		board:       board,
		cursorRow:   0,
		cursorCol:   0,
		minesLeft:   mines,
		cellsToReveal: rows*cols - mines,
		isGameOver:  false,
		isWin:       false,
		message:     "Bem-vindo ao Campo Minado! Use as setas para navegar.",
		keymap:      keys,
		help:        help,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Processar teclas somente se o jogo não acabou
		if !m.isGameOver {
			switch {
			case key.Matches(msg, m.keymap.Up):
				if m.cursorRow > 0 {
					m.cursorRow--
				}
			case key.Matches(msg, m.keymap.Down):
				if m.cursorRow < rows-1 {
					m.cursorRow++
				}
			case key.Matches(msg, m.keymap.Left):
				if m.cursorCol > 0 {
					m.cursorCol--
				}
			case key.Matches(msg, m.keymap.Right):
				if m.cursorCol < cols-1 {
					m.cursorCol++
				}
			case key.Matches(msg, m.keymap.Reveal):
				m.revealCell(m.cursorRow, m.cursorCol)
				m.checkGameStatus()
			case key.Matches(msg, m.keymap.Flag):
				m.flagCell(m.cursorRow, m.cursorCol)
			}
		}
		
		// Sempre verifica para sair
		if key.Matches(msg, m.keymap.Quit) {
			return m, tea.Quit
		}
	}
	
	return m, nil
}

func (m Model) View() string {
	if m.isGameOver {
		if m.isWin {
			m.message = "Parabéns! Você venceu!"
		} else {
			m.message = "BOOM! Você encontrou uma mina!"
		}
	}
	
	// Cabeçalho
	s := headerStyle.Render("=== CAMPO MINADO - MODO FÁCIL ===\n\n")
	
	// Informação do jogo
	s += infoStyle.Render(fmt.Sprintf("Minas restantes: %d\n", m.minesLeft))
	s += infoStyle.Render(m.message + "\n\n")
	
	// Índices das colunas
	s += "   "
	for c := 0; c < cols; c++ {
		s += fmt.Sprintf(" %d ", c)
	}
	s += "\n"
	
	// Tabuleiro
	for r := 0; r < rows; r++ {
		s += fmt.Sprintf(" %d ", r)
		
		for c := 0; c < cols; c++ {
			cell := m.board[r][c]
			isCursor := r == m.cursorRow && c == m.cursorCol && !m.isGameOver
			
			cellValue := " "
			style := hiddenStyle
			
			if m.isGameOver && cell.hasMine {
				// Mostra todas as minas no final do jogo
				cellValue = "*"
				if cell.isRevealed {
					style = mineStyle
				} else {
					style = revealedStyle
				}
			} else if cell.isFlagged {
				cellValue = "F"
				style = flagStyle
			} else if cell.isRevealed {
				style = revealedStyle
				if cell.adjacentMines > 0 {
					cellValue = fmt.Sprintf("%d", cell.adjacentMines)
				} else {
					cellValue = " "
				}
			} else {
				cellValue = "■"
				style = hiddenStyle
			}
			
			// Aplicar estilo do cursor se for a célula atual
			if isCursor {
				style = cursorStyle
			}
			
			s += style.Render(cellValue)
		}
		s += "\n"
	}
	
	// Ajuda
	s += "\n" + m.help.View(m.keymap)
	
	// Rodapé
	s += "\n" + footerStyle.Render("Pressione q para sair")
	
	return s
}

// Revelar uma célula no tabuleiro
func (m *Model) revealCell(row, col int) {
	cell := &m.board[row][col]
	
	// Se a célula já estiver revelada ou marcada com uma bandeira, não faz nada
	if cell.isRevealed || cell.isFlagged {
		return
	}
	
	// Revelar a célula
	cell.isRevealed = true
	
	// Se for uma mina, fim de jogo
	if cell.hasMine {
		m.isGameOver = true
		return
	}
	
	// Decrementar o contador de células para revelar
	m.cellsToReveal--
	
	// Se for uma célula vazia (sem minas adjacentes), revelar células vizinhas automaticamente
	if cell.adjacentMines == 0 {
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				nr, nc := row+dr, col+dc
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !m.board[nr][nc].isRevealed {
					m.revealCell(nr, nc)
				}
			}
		}
	}
}

// Marcar/desmarcar uma célula com bandeira
func (m *Model) flagCell(row, col int) {
	cell := &m.board[row][col]
	
	// Se a célula já estiver revelada, não pode marcar
	if cell.isRevealed {
		return
	}
	
	// Alternar o estado da bandeira
	if cell.isFlagged {
		cell.isFlagged = false
		m.minesLeft++
	} else {
		cell.isFlagged = true
		m.minesLeft--
	}
}

// Verificar o status do jogo (vitória/derrota)
func (m *Model) checkGameStatus() {
	// Verificar se o jogador revelou todas as células seguras
	if m.cellsToReveal == 0 {
		m.isGameOver = true
		m.isWin = true
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Erro ao iniciar Bubble Tea: %v\n", err)
		os.Exit(1)
	}
}