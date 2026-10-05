package styles

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// agentSmithFont is the 5x7 bitmap font used to render the AGENT-SMITH logo.
// Each "on" cell (1) is drawn with the glyph's own character so the big "A"
// is built from smaller "A" pixels, the big "G" from smaller "G" pixels, etc.
var agentSmithFont = map[rune][]string{
	'A': {"01110", "10001", "10001", "11111", "10001", "10001", "10001"},
	'G': {"01111", "10000", "10000", "10111", "10001", "10001", "01110"},
	'E': {"11111", "10000", "10000", "11110", "10000", "10000", "11111"},
	'N': {"10001", "11001", "10101", "10101", "10101", "10011", "10001"},
	'T': {"11111", "00100", "00100", "00100", "00100", "00100", "00100"},
	'-': {"00000", "00000", "00000", "11111", "00000", "00000", "00000"},
	'S': {"01110", "10001", "10000", "01110", "00001", "10001", "01110"},
	'M': {"10001", "11011", "10101", "10101", "10001", "10001", "10001"},
	'I': {"11111", "00100", "00100", "00100", "00100", "00100", "11111"},
	'H': {"10001", "10001", "10001", "11111", "10001", "10001", "10001"},
}

const agentSmithText = "AGENT-SMITH"

// logoLines holds the AGENT-SMITH bitmap art, computed once at package init.
var logoLines = buildAgentSmithArt()

// buildAgentSmithArt renders agentSmithText through agentSmithFont so every
// glyph becomes a column of its own smaller character.
func buildAgentSmithArt() []string {
	const rows = 7
	lines := make([]string, rows)
	runes := []rune(agentSmithText)
	for i, ch := range runes {
		pattern, ok := agentSmithFont[ch]
		if !ok {
			pattern = agentSmithFont['-']
		}
		for r := 0; r < rows; r++ {
			row := make([]byte, 5)
			for c := 0; c < 5; c++ {
				if pattern[r][c] == '1' {
					row[c] = byte(ch)
				} else {
					row[c] = ' '
				}
			}
			if i > 0 {
				lines[r] += " "
			}
			lines[r] += string(row)
		}
	}
	return lines
}

// gradientColors defines the top-to-bottom Matrix-green gradient for the logo.
var gradientColors = []lipgloss.Color{
	ColorMatrixDim,    // band 1 (top)
	ColorMatrixMid,    // band 2
	ColorMatrixBright, // band 3 (middle)
	ColorMatrixMid,    // band 4
	ColorMatrixDim,    // band 5 (bottom)
}

// RenderLogo returns the AGENT-SMITH ASCII logo with a Matrix-green gradient.
func RenderLogo() string {
	total := len(logoLines)
	if total == 0 {
		return ""
	}

	bands := len(gradientColors)
	var b strings.Builder

	for i, line := range logoLines {
		bandIdx := (i * bands) / total
		if bandIdx >= bands {
			bandIdx = bands - 1
		}
		style := lipgloss.NewStyle().Foreground(gradientColors[bandIdx]).Bold(true)
		b.WriteString(style.Render(line))
		if i < total-1 {
			b.WriteByte('\n')
		}
	}

	return b.String()
}