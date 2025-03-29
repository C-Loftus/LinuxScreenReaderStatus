package a11yStatus

import (
	"encoding/xml"
	"testing"
)

// Test serialization and deserialization of dbus xml
func TestDBusIntrospectionXML(t *testing.T) {
	inputXML := `<?xml version='1.0'?>
<node>
  <interface name="org.freedesktop.DBus.Introspectable">
    <method name="Introspect">
      <arg type="s" name="xml_data" direction="out"/>
    </method>
  </interface>
</node>`

	var node node
	err := xml.Unmarshal([]byte(inputXML), &node)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	if len(node.Interfaces) != 1 {
		t.Errorf("Expected 1 interface, got %d", len(node.Interfaces))
	}

	if node.Interfaces[0].Name != "org.freedesktop.DBus.Introspectable" {
		t.Errorf("Unexpected interface name: %s", node.Interfaces[0].Name)
	}

	if len(node.Interfaces[0].Methods) != 1 || node.Interfaces[0].Methods[0].Name != "Introspect" {
		t.Errorf("Method 'Introspect' not found in parsed XML")
	}

	// Test serialization
	outputXML, err := xml.MarshalIndent(node, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	if string(outputXML) == "" {
		t.Error("Serialized XML is empty")
	}
}

// Test the screen reader status.
// To test this you should first start Orca or set the dbus property
// yourself. We do this instead of automatically setting it ourselves
// so we don't potentially disrupt the user's screen reader
func TestScreenReaderStatus(t *testing.T) {
	status, err := ScreenReaderStatus()
	if err != nil {
		t.Fatalf("Failed to get A11yStatus: %v", err)
	}

	if !status.IsEnabled {
		t.Error("A11y is not enabled; This is a sign either Orca is not running or the dbus property was not retrieved properly")
	}
}
