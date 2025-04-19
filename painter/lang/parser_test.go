package lang

import (
    "fmt"
    "io"
    "io/ioutil"
    "strings"
    "strconv"
    "github.com/roman-mazur/architecture-lab-3/painter" // Виправлений імпорт
)

type Parser struct{}

// Змінено сигнатуру функції, щоб вона приймала io.Reader
func (p *Parser) Parse(r io.Reader) ([]painter.Operation, error) {
    data, err := ioutil.ReadAll(r)
    if err != nil {
        return nil, err
    }
    
    cmd := string(data)
    parts := strings.Fields(cmd)
    if len(parts) == 0 {
        return nil, fmt.Errorf("empty command")
    }
    
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
    
    coords := make([]float64, 4)
    for i, arg := range args {
        val, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            return nil, fmt.Errorf("invalid coordinate %s: %v", arg, err)
        }
        coords[i] = val
    }
    
    return []painter.Operation{painter.BgRectOp{
        X1: coords[0], Y1: coords[1],
        X2: coords[2], Y2: coords[3],
    }}, nil
}

func (p *Parser) parseFigure(args []string) ([]painter.Operation, error) {
    if len(args) != 2 {
        return nil, fmt.Errorf("figure requires 2 arguments")
    }
    
    x, err := strconv.ParseFloat(args[0], 64)
    if err != nil {
        return nil, fmt.Errorf("invalid x coordinate: %v", err)
    }
    
    y, err := strconv.ParseFloat(args[1], 64)
    if err != nil {
        return nil, fmt.Errorf("invalid y coordinate: %v", err)
    }
    
    return []painter.Operation{painter.FigureOp{X: x, Y: y}}, nil
}

func (p *Parser) parseMove(args []string) ([]painter.Operation, error) {
    if len(args) != 2 {
        return nil, fmt.Errorf("move requires 2 arguments")
    }
    
    x, err := strconv.ParseFloat(args[0], 64)
    if err != nil {
        return nil, fmt.Errorf("invalid x coordinate: %v", err)
    }
    
    y, err := strconv.ParseFloat(args[1], 64)
    if err != nil {
        return nil, fmt.Errorf("invalid y coordinate: %v", err)
    }
    
    return []painter.Operation{painter.MoveOp{X: x, Y: y}}, nil
}
