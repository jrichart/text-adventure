package ast

import "fmt"

type Node interface {
	String() string
}

type Command struct {
	VerbGroup VerbGroup
	NounGroup NounGroup
}

type NounGroup struct {
	Adjectives []*Adjective
	Noun       Noun
}

type VerbGroup struct {
	Verb Verb
}

type Noun struct {
	Value string
}

type Adjective struct {
	Value string
}

type Verb struct {
	Value string
}

func (c *Command) String() string {
	result := ""
	result += c.VerbGroup.String()
	result += c.NounGroup.String()
	return result
}

func (ng *NounGroup) String() string {
	result := ""
	if len(ng.Adjectives) > 0 {
		for _, adj := range ng.Adjectives {
			result += adj.String() + " "
		}
	}
	if result == "" && ng.Noun.Value != "" {
		result += " "
	}
	result += ng.Noun.String()
	return result
}

func (n *Noun) String() string {
	return n.Value
}

func (a *Adjective) String() string {
	return a.Value
}

func (vb *VerbGroup) String() string {
	return vb.Verb.String()
}

func (v *Verb) String() string {
	return v.Value
}

type ParseError struct {
	Message string
	Token   string
}

func (pe *ParseError) Error() string {
	return fmt.Sprintf("parse error: %s (at token: %q)", pe.Message, pe.Token)
}
