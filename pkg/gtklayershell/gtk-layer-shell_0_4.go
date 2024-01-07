// Code generated by girgen. DO NOT EDIT.

package gtklayershell

// #include <stdlib.h>
// #include <gtk-layer-shell/gtk-layer-shell.h>
import "C"

// The function returns the following values:
//
//   - guint: major version number of the GTK Layer Shell library.
//
func GetMajorVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_layer_get_major_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// The function returns the following values:
//
//   - guint: micro/patch version number of the GTK Layer Shell library.
//
func GetMicroVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_layer_get_micro_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// The function returns the following values:
//
//   - guint: minor version number of the GTK Layer Shell library.
//
func GetMinorVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_layer_get_minor_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}