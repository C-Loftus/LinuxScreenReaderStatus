# LinuxScreenReaderStatus

This simple library uses go D-Bus bindings to get the current screen reader status on Linux. 

It is useful for contextually enabling accessibility tools that may interop with screen readers like Orca.

## D-Bus Background

For a general overview of D-Bus that would be useful in extending this library to include other accessibility properties, see the links here:

- https://github.com/flexibeast/guides/blob/master/dbus.md
- https://dbus.freedesktop.org/doc/dbus-tutorial.html
- https://pythonhosted.org/txdbus/dbus_overview.html