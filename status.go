package a11yStatus

import (
	"encoding/xml"
	"fmt"

	"github.com/godbus/dbus/v5"
)

// OrgA11yStatus is the struct for the org.a11y.Status dbus interface
// It contains the IsEnabled and ScreenReaderEnabled properties
// which both represent the current screen reader status but in different ways
type OrgA11yStatus struct {
	// A screen reader is running right now
	// Enabled by Orca on startup
	IsEnabled bool
	// A screen reader should be running, e. g. the preference is turned on in the system settings
	// Checked but not set by Orca
	ScreenReaderEnabled bool
}

// Get the current screen reader status as defined by the org.a11y.Status interface
func ScreenReaderStatus() (OrgA11yStatus, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return OrgA11yStatus{}, fmt.Errorf("failed to connect to session bus: %v", err)
	}
	defer conn.Close()

	obj := conn.Object("org.a11y.Bus", "/org/a11y/bus")

	// Call the standard DBus "Introspect" method
	var introspectionXML string
	err = obj.Call("org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&introspectionXML)
	if err != nil {
		return OrgA11yStatus{}, fmt.Errorf("failed to call Introspect method: %v", err)
	}

	// Parse the XML using the defined structs
	var node node
	err = xml.Unmarshal([]byte(introspectionXML), &node)
	if err != nil {
		return OrgA11yStatus{}, fmt.Errorf("failed to unmarshal XML: %v", err)
	}

	var isEnabled, screenReaderEnabled bool

	for _, iface := range node.Interfaces {
		if iface.Name == "org.a11y.Status" {
			if variant, err := obj.GetProperty("org.a11y.Status.IsEnabled"); err == nil {
				_ = variant.Store(&isEnabled)
			} else {
				return OrgA11yStatus{}, fmt.Errorf("failed to get IsEnabled property: %v", err)
			}
			if variant, err := obj.GetProperty("org.a11y.Status.ScreenReaderEnabled"); err == nil {
				_ = variant.Store(&screenReaderEnabled)
			} else {
				return OrgA11yStatus{}, fmt.Errorf("failed to get ScreenReaderEnabled property: %v", err)
			}
			return OrgA11yStatus{IsEnabled: isEnabled, ScreenReaderEnabled: screenReaderEnabled}, nil
		}
	}

	return OrgA11yStatus{}, nil
}
