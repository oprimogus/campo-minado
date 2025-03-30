# 💣 GO-MINESWEEPER 💣

Um jogo de Campo Minado interativo implementado em Go com uma elegante interface de texto (TUI).

![Campo Minado Screenshot](https://via.placeholder.com/800x400?text=Go+Minesweeper+Screenshot)

## 🚀 Características

- Interface de texto totalmente interativa usando [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- Design visual elegante com [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- Controles de teclado intuitivos (estilo vim ou setas direcionais)
- Revelação automática de células vazias
- Sistema de marcação com bandeiras
- Exibição de minas ao final do jogo
- Indicação visual de vitória/derrota
- Dificuldade inicial "Easy" (8x8 com 10 minas)

## 🎮 Como Jogar

### Controles
- **Movimento**: Setas direcionais ou teclas Vim (`h`, `j`, `k`, `l`)
- **Revelar célula**: `Enter`, `r` ou `Espaço`
- **Marcar/desmarcar bandeira**: `f`
- **Sair do jogo**: `q` ou `Ctrl+C`

### Regras
1. Revelar todas as células que não contêm minas para vencer.
2. Se você revelar uma célula com mina, o jogo termina.
3. Os números indicam quantas minas estão adjacentes àquela célula.
4. Use as bandeiras para marcar onde você acredita que há minas.

## 📦 Instalação

### Pré-requisitos
- Go 1.16+ instalado

### Passos
1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/go-minesweeper.git
   cd go-minesweeper
   ```

2. Instale as dependências:
   ```bash
   go get github.com/charmbracelet/bubbletea
   go get github.com/charmbracelet/bubbles/help
   go get github.com/charmbracelet/bubbles/key
   go get github.com/charmbracelet/lipgloss
   ```

3. Compile o jogo:
   ```bash
   go build -o minesweeper
   ```

4. Execute:
   ```bash
   ./minesweeper
   ```

## 🛠️ Tecnologias Utilizadas

- [Go](https://golang.org/) - Linguagem de programação
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Framework TUI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Estilização de texto
- [Bubbles](https://github.com/charmbracelet/bubbles) - Componentes TUI reutilizáveis

## 📋 Recursos Futuros

- [ ] Múltiplos níveis de dificuldade (Médio, Difícil)
- [ ] Sistema de pontuação e tempo
- [ ] Salvar recordes
- [ ] Customização do tabuleiro
- [ ] Temas de cores
- [ ] Sons (com BubbleTea Tea Commands)

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

1. Fork o projeto
2. Crie sua branch de feature (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

## 📜 Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👏 Agradecimentos

- [Charmbracelet](https://github.com/charmbracelet) pela biblioteca Bubble Tea
- Todos os contribuidores e entusiastas de TUIs em Go

---

Feito com ❤️ e Go