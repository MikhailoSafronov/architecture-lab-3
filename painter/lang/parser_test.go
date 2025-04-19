package lang

import (
    "strings"
    "testing"
)

func TestParser_Parse(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantOps int
        wantErr bool
    }{
        {"valid bgrect", "bgrect 0.1 0.2 0.3 0.4", 1, false},
        {"invalid bgrect", "bgrect 0.1", 0, true},
        {"move command", "move 0.5 0.5", 1, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            p := &Parser{}
            ops, err := p.Parse(strings.NewReader(tt.input))
            if (err != nil) != tt.wantErr {
                t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if len(ops) != tt.wantOps {
                t.Errorf("Parse() got %d ops, want %d", len(ops), tt.wantOps)
            }
        })
    }
}
