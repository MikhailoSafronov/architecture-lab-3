package lang

import (
	"bufio"
	"errors"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

type Parser struct{}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	var ops []painter.Operation

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := parts[0]
		args := parts[1:]

		switch cmd {
		case "white":
			ops = append(ops, painter.BgColorOp{Color: color.White})
		case "green":
			ops = append(ops, painter.BgColorOp{Color: color.RGBA{G: 0xff, A: 0xff}})
		case "bgrect":
			if len(args) != 4 {
				return nil, errors.New("bgrect requires 4 arguments")
			}
			x1, _ := strconv.ParseFloat(args[0], 64)
			y1, _ := strconv.ParseFloat(args[1], 64)
			x2, _ := strconv.ParseFloat(args[2], 64)
			y2, _ := strconv.ParseFloat(args[3], 64)
			ops = append(ops, painter.BgRectOp{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "figure":
			if len(args) != 2 {
				return nil, errors.New("figure requires 2 arguments")
			}
			x, _ := strconv.ParseFloat(args[0], 64)
			y, _ := strconv.ParseFloat(args[1], 64)
			ops = append(ops, painter.FigureOp{X: x, Y: y})
		case "move":
			if len(args) != 2 {
				return nil, errors.New("move requires 2 arguments")
			}
			x, _ := strconv.ParseFloat(args[0], 64)
			y, _ := strconv.ParseFloat(args[1], 64)
			ops = append(ops, painter.MoveOp{X: x, Y: y})
		case "update":
			ops = append(ops, painter.UpdateOp{})
		case "reset":
			ops = append(ops, painter.ResetOp{})
		default:
			return nil, errors.New("unknown command: " + cmd)
		}
	}

	return ops, nil
}
