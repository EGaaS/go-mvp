package tx

import "fmt"

type AppendPage struct {
	Header
	Global string
	Name   string
	Value  string
}

func (a AppendPage) ForSign() string {
	return fmt.Sprintf("%s,%s,%d,%d,%s,%s,%s", a.Header.Type, a.Header.Time, a.Header.UserID, a.Header.StateID, a.Global, a.Name, a.Value)
}

type EditPage struct {
	Header
	Global     string
	Name       string
	Value      string
	Menu       string
	Conditions string
}

func (a EditPage) ForSign() string {
	return fmt.Sprintf("%s,%s,%d,%d,%s,%s,%s,%s,%s", a.Header.Type, a.Header.Time, a.Header.UserID, a.Header.StateID, a.Global, a.Name, a.Value, a.Menu, a.Conditions)
}

type NewPage struct {
	Header
	Global     string
	Name       string
	Value      string
	Menu       string
	Conditions string
}

func (a NewPage) ForSign() string {
	return fmt.Sprintf("%s,%s,%d,%d,%s,%s,%s,%s,%s", a.Header.Type, a.Header.Time, a.Header.UserID, a.Header.StateID, a.Global, a.Name, a.Value, a.Menu, a.Conditions)
}