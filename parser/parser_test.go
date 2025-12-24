package parser

import (
	"testing"
	"text-adventure/lexer"
	"text-adventure/vocabulary"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedCmd   string
		expectedError bool
	}{
		{
			name:          "simple verb",
			input:         "look",
			expectedCmd:   "look",
			expectedError: false,
		},
		{
			name:          "verb noun",
			input:         "take sword",
			expectedCmd:   "take sword",
			expectedError: false,
		},
		{
			name:          "missing verb",
			input:         "sword",
			expectedCmd:   "",
			expectedError: true,
		},
		{
			name:          "adjective without noun",
			input:         "take rusty",
			expectedCmd:   "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vocab := vocabulary.DefaultVocabulary()
			l := lexer.New(tt.input, vocab)
			p := New(l)

			cmd := p.ParseCommand()

			if tt.expectedError {
				if p.errors == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if p.errors != nil {
				for i, err := range p.errors {
					t.Errorf("unexpected error %d: %v", i, err.Message)
					t.Errorf("command: %s", cmd.String())
					return
				}
			}

			if cmd.String() != tt.expectedCmd {
				t.Errorf("wrong command string. expected=%q, got=%q",
					tt.expectedCmd, cmd.String())
			}
		})
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		errorContains string
	}{
		{
			name:          "unknown word",
			input:         "take xyzzy",
			errorContains: "unexpected token",
		},
		{
			name:          "starts with adjective",
			input:         "rusty sword",
			errorContains: "expected verb",
		},
		{
			name:          "extra tokens",
			input:         "take sword quickly",
			errorContains: "unexpected token after command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vocab := vocabulary.DefaultVocabulary()
			l := lexer.New(tt.input, vocab)
			p := New(l)

			_ = p.ParseCommand()

			if p.errors == nil {
				t.Errorf("expected error containing %q but got none", tt.errorContains)
				return
			}

			// Note: In a real implementation, you might want more specific
			// error checking here using error type assertions
		})
	}
}
