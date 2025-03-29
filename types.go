package a11yStatus

import "encoding/xml"

// node represents the root element of the introspection XML
// It contains multiple interfaces.
type node struct {
	XMLName    xml.Name     `xml:"node"`
	Interfaces []_interface `xml:"interface"`
}

// _interface represents a DBus interface containing methods, signals, and properties.
// it is prefixed with an underscore to avoid clashing with the golang interface keyword
type _interface struct {
	XMLName    xml.Name   `xml:"interface"`
	Name       string     `xml:"name,attr"`
	Methods    []method   `xml:"method"`
	Signals    []signal   `xml:"signal"`
	Properties []property `xml:"property"`
}

// method represents a method inside a DBus interface.
type method struct {
	XMLName xml.Name `xml:"method"`
	Name    string   `xml:"name,attr"`
	Args    []arg    `xml:"arg"`
}

// signal represents a signal inside a DBus interface.
type signal struct {
	XMLName xml.Name `xml:"signal"`
	Name    string   `xml:"name,attr"`
	Args    []arg    `xml:"arg"`
}

// property represents a property inside a DBus interface.
type property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Access  string   `xml:"access,attr"`
}

// arg represents an argument inside a method or signal.
type arg struct {
	XMLName   xml.Name `xml:"arg"`
	Type      string   `xml:"type,attr"`
	Name      string   `xml:"name,attr,omitempty"`
	Direction string   `xml:"direction,attr,omitempty"`
}
