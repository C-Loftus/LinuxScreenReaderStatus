package a11yStatus

import "encoding/xml"

// Node represents the root element of the introspection XML
// It contains multiple interfaces.
type Node struct {
	XMLName    xml.Name    `xml:"node"`
	Interfaces []Interface `xml:"interface"`
}

// Interface represents a DBus interface containing methods, signals, and properties.
type Interface struct {
	XMLName    xml.Name   `xml:"interface"`
	Name       string     `xml:"name,attr"`
	Methods    []Method   `xml:"method"`
	Signals    []Signal   `xml:"signal"`
	Properties []Property `xml:"property"`
}

// Method represents a method inside a DBus interface.
type Method struct {
	XMLName xml.Name `xml:"method"`
	Name    string   `xml:"name,attr"`
	Args    []Arg    `xml:"arg"`
}

// Signal represents a signal inside a DBus interface.
type Signal struct {
	XMLName xml.Name `xml:"signal"`
	Name    string   `xml:"name,attr"`
	Args    []Arg    `xml:"arg"`
}

// Property represents a property inside a DBus interface.
type Property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Access  string   `xml:"access,attr"`
}

// Arg represents an argument inside a method or signal.
type Arg struct {
	XMLName   xml.Name `xml:"arg"`
	Type      string   `xml:"type,attr"`
	Name      string   `xml:"name,attr,omitempty"`
	Direction string   `xml:"direction,attr,omitempty"`
}
