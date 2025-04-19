package lang

import (
    "strings"
    "strconv"
    "github.com/MikhailoSafronov/architecture-lab-3/painter"
)

type Parser struct{}

func (p *Parser) Parse(cmd string) ([]painter.Operation, error) {
    parts := strings.Fields(cmd)
    switch parts[0] {
    case "bgrect":
        return p.parseBgRect(parts[1:])
    case "figure":
        return p.parseFigure(parts[1:])
    case "move":
        return p.parseMove(parts[1:])
    case "update":
        return []painter.Operation{painter.UpdateOp{}}, nil
    default:
        return nil, fmt.Errorf("unknown command: %s", parts[0])
    }
}

func (p *Parser) parseBgRect(args []string) ([]painter.Operation, error) {
    if len(args) != 4 {
        return nil, fmt.Errorf("bgrect requires 4 arguments")
    }
    
    coords := make([]float32, 4)
    for i, arg := range args {
        val, _ := strconv.ParseFloat(arg, 32)
        coords[i] = float32(val)
    }
    
    return []painter.Operation{painter.BgRect{
        X1: coords[0], Y1: coords[1],
        X2: coords[2], Y2: coords[3],
    }}, nil
}
