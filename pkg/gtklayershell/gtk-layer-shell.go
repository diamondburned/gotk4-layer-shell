// Code generated by girgen. DO NOT EDIT.

package gtklayershell

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #cgo pkg-config: gtk-layer-shell-0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk-layer-shell/gtk-layer-shell.h>
import "C"

type Edge int

const (
	// LayerShellEdgeLeft: left edge of the screen.
	LayerShellEdgeLeft Edge = iota
	// LayerShellEdgeRight: right edge of the screen.
	LayerShellEdgeRight
	// LayerShellEdgeTop: top edge of the screen.
	LayerShellEdgeTop
	// LayerShellEdgeBottom: bottom edge of the screen.
	LayerShellEdgeBottom
	// LayerShellEdgeEntryNumber: should not be used except to get the number of
	// entries. (NOTE: may change in future releases as more entries are added).
	LayerShellEdgeEntryNumber
)

// String returns the name in string for Edge.
func (e Edge) String() string {
	switch e {
	case LayerShellEdgeLeft:
		return "Left"
	case LayerShellEdgeRight:
		return "Right"
	case LayerShellEdgeTop:
		return "Top"
	case LayerShellEdgeBottom:
		return "Bottom"
	case LayerShellEdgeEntryNumber:
		return "EntryNumber"
	default:
		return fmt.Sprintf("Edge(%d)", e)
	}
}

// KeyboardMode: GTK_LAYER_SHELL_KEYBOARD_MODE_NONE: This window should not
// receive keyboard events. GTK_LAYER_SHELL_KEYBOARD_MODE_EXCLUSIVE: This window
// should have exclusive focus if it is on the top or overlay layer.
// GTK_LAYER_SHELL_KEYBOARD_MODE_ON_DEMAND: The user should be able to focus and
// unfocues this window in an implementation defined way. Not supported for
// protocol version < 4. GTK_LAYER_SHELL_KEYBOARD_MODE_ENTRY_NUMBER: Should not
// be used except to get the number of entries. (NOTE: may change in future
// releases as more entries are added).
type KeyboardMode int

const (
	LayerShellKeyboardModeNone KeyboardMode = iota
	LayerShellKeyboardModeExclusive
	LayerShellKeyboardModeOnDemand
	LayerShellKeyboardModeEntryNumber
)

// String returns the name in string for KeyboardMode.
func (k KeyboardMode) String() string {
	switch k {
	case LayerShellKeyboardModeNone:
		return "None"
	case LayerShellKeyboardModeExclusive:
		return "Exclusive"
	case LayerShellKeyboardModeOnDemand:
		return "OnDemand"
	case LayerShellKeyboardModeEntryNumber:
		return "EntryNumber"
	default:
		return fmt.Sprintf("KeyboardMode(%d)", k)
	}
}

type Layer int

const (
	// LayerShellLayerBackground: background layer.
	LayerShellLayerBackground Layer = iota
	// LayerShellLayerBottom: bottom layer.
	LayerShellLayerBottom
	// LayerShellLayerTop: top layer.
	LayerShellLayerTop
	// LayerShellLayerOverlay: overlay layer.
	LayerShellLayerOverlay
	// LayerShellLayerEntryNumber: should not be used except to get the number
	// of entries. (NOTE: may change in future releases as more entries are
	// added).
	LayerShellLayerEntryNumber
)

// String returns the name in string for Layer.
func (l Layer) String() string {
	switch l {
	case LayerShellLayerBackground:
		return "Background"
	case LayerShellLayerBottom:
		return "Bottom"
	case LayerShellLayerTop:
		return "Top"
	case LayerShellLayerOverlay:
		return "Overlay"
	case LayerShellLayerEntryNumber:
		return "EntryNumber"
	default:
		return fmt.Sprintf("Layer(%d)", l)
	}
}

// AutoExclusiveZoneEnable: when auto exclusive zone is enabled, exclusive zone
// is automatically set to the size of the window + relevant margin. To disable
// auto exclusive zone, just set the exclusive zone to 0 or any other fixed
// value.
//
// NOTE: you can control the auto exclusive zone by changing the margin on the
// non-anchored edge. This behavior is specific to gtk-layer-shell and not part
// of the underlying protocol.
func AutoExclusiveZoneEnable(window *gtk.Window) {
	var _arg1 *C.GtkWindow // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_layer_auto_exclusive_zone_enable(_arg1)
	runtime.KeepAlive(window)
}

func AutoExclusiveZoneIsEnabled(window *gtk.Window) bool {
	var _arg1 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_layer_auto_exclusive_zone_is_enabled(_arg1)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func GetAnchor(window *gtk.Window, edge Edge) bool {
	var _arg1 *C.GtkWindow        // out
	var _arg2 C.GtkLayerShellEdge // out
	var _cret C.gboolean          // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkLayerShellEdge(edge)

	_cret = C.gtk_layer_get_anchor(_arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(edge)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func GetExclusiveZone(window *gtk.Window) int {
	var _arg1 *C.GtkWindow // out
	var _cret C.int        // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_layer_get_exclusive_zone(_arg1)
	runtime.KeepAlive(window)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GetKeyboardInteractivity: deprecated: Use gtk_layer_get_keyboard_mode ()
// instead.
func GetKeyboardInteractivity(window *gtk.Window) bool {
	var _arg1 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_layer_get_keyboard_interactivity(_arg1)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func GetKeyboardMode(window *gtk.Window) KeyboardMode {
	var _arg1 *C.GtkWindow                // out
	var _cret C.GtkLayerShellKeyboardMode // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_layer_get_keyboard_mode(_arg1)
	runtime.KeepAlive(window)

	var _keyboardMode KeyboardMode // out

	_keyboardMode = KeyboardMode(_cret)

	return _keyboardMode
}

func GetLayer(window *gtk.Window) Layer {
	var _arg1 *C.GtkWindow         // out
	var _cret C.GtkLayerShellLayer // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_layer_get_layer(_arg1)
	runtime.KeepAlive(window)

	var _layer Layer // out

	_layer = Layer(_cret)

	return _layer
}

func GetMajorVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_layer_get_major_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func GetMargin(window *gtk.Window, edge Edge) int {
	var _arg1 *C.GtkWindow        // out
	var _arg2 C.GtkLayerShellEdge // out
	var _cret C.int               // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkLayerShellEdge(edge)

	_cret = C.gtk_layer_get_margin(_arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(edge)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func GetMicroVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_layer_get_micro_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func GetMinorVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_layer_get_minor_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// GetNamespace: NOTE: this function does not return ownership of the string. Do
// not free the returned string. Future calls into the library may invalidate
// the returned string.
func GetNamespace(window *gtk.Window) string {
	var _arg1 *C.GtkWindow // out
	var _cret *C.char      // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_layer_get_namespace(_arg1)
	runtime.KeepAlive(window)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetProtocolVersion: may block for a Wayland roundtrip the first time it's
// called.
func GetProtocolVersion() uint {
	var _cret C.guint // in

	_cret = C.gtk_layer_get_protocol_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// InitForWindow: set the window up to be a layer surface once it is mapped.
// this must be called before the window is realized.
func InitForWindow(window *gtk.Window) {
	var _arg1 *C.GtkWindow // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_layer_init_for_window(_arg1)
	runtime.KeepAlive(window)
}

func IsLayerWindow(window *gtk.Window) bool {
	var _arg1 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_layer_is_layer_window(_arg1)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSupported: may block for a Wayland roundtrip the first time it's called.
func IsSupported() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_layer_is_supported()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAnchor: set whether window should be anchored to edge.
//
// - If two perpendicular edges are anchored, the surface with be anchored to
// that corner
//
// - If two opposite edges are anchored, the window will be stretched across the
// screen in that direction
//
// Default is FALSE for each LayerShellEdge.
func SetAnchor(window *gtk.Window, edge Edge, anchorToEdge bool) {
	var _arg1 *C.GtkWindow        // out
	var _arg2 C.GtkLayerShellEdge // out
	var _arg3 C.gboolean          // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkLayerShellEdge(edge)
	if anchorToEdge {
		_arg3 = C.TRUE
	}

	C.gtk_layer_set_anchor(_arg1, _arg2, _arg3)
	runtime.KeepAlive(window)
	runtime.KeepAlive(edge)
	runtime.KeepAlive(anchorToEdge)
}

// SetExclusiveZone has no effect unless the surface is anchored to an edge.
// Requests that the compositor does not place other surfaces within the given
// exclusive zone of the anchored edge. For example, a panel can request to not
// be covered by maximized windows. See wlr-layer-shell-unstable-v1.xml for
// details.
//
// Default is 0.
func SetExclusiveZone(window *gtk.Window, exclusiveZone int) {
	var _arg1 *C.GtkWindow // out
	var _arg2 C.int        // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.int(exclusiveZone)

	C.gtk_layer_set_exclusive_zone(_arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(exclusiveZone)
}

// SetKeyboardInteractivity: whether the window should receive keyboard events
// from the compositor.
//
//
// Default is FALSE
//
// Deprecated: Use gtk_layer_set_keyboard_mode () instead.
func SetKeyboardInteractivity(window *gtk.Window, interactivity bool) {
	var _arg1 *C.GtkWindow // out
	var _arg2 C.gboolean   // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	if interactivity {
		_arg2 = C.TRUE
	}

	C.gtk_layer_set_keyboard_interactivity(_arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(interactivity)
}

// SetKeyboardMode sets if/when window should receive keyboard events from the
// compositor, see GtkLayerShellKeyboardMode for details.
//
// Default is K_LAYER_SHELL_KEYBOARD_MODE_NONE.
func SetKeyboardMode(window *gtk.Window, mode KeyboardMode) {
	var _arg1 *C.GtkWindow                // out
	var _arg2 C.GtkLayerShellKeyboardMode // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkLayerShellKeyboardMode(mode)

	C.gtk_layer_set_keyboard_mode(_arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(mode)
}

// SetLayer: set the "layer" on which the surface appears (controls if it is
// over top of or below other surfaces). The layer may be changed on-the-fly in
// the current version of the layer shell protocol, but on compositors that only
// support an older version the window is remapped so the change can take
// effect.
//
// Default is K_LAYER_SHELL_LAYER_TOP.
func SetLayer(window *gtk.Window, layer Layer) {
	var _arg1 *C.GtkWindow         // out
	var _arg2 C.GtkLayerShellLayer // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkLayerShellLayer(layer)

	C.gtk_layer_set_layer(_arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(layer)
}

// SetMargin: set the margin for a specific edge of a window. Effects both
// surface's distance from the edge and its exclusive zone size (if auto
// exclusive zone enabled).
//
// Default is 0 for each LayerShellEdge.
func SetMargin(window *gtk.Window, edge Edge, marginSize int) {
	var _arg1 *C.GtkWindow        // out
	var _arg2 C.GtkLayerShellEdge // out
	var _arg3 C.int               // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkLayerShellEdge(edge)
	_arg3 = C.int(marginSize)

	C.gtk_layer_set_margin(_arg1, _arg2, _arg3)
	runtime.KeepAlive(window)
	runtime.KeepAlive(edge)
	runtime.KeepAlive(marginSize)
}

// SetMonitor: set the output for the window to be placed on, or NULL to let the
// compositor choose. If the window is currently mapped, it will get remapped so
// the change can take effect.
//
// Default is NULL.
func SetMonitor(window *gtk.Window, monitor *gdk.Monitor) {
	var _arg1 *C.GtkWindow  // out
	var _arg2 *C.GdkMonitor // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	C.gtk_layer_set_monitor(_arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(monitor)
}

// SetNamespace: set the "namespace" of the surface.
//
// No one is quite sure what this is for, but it probably should be something
// generic ("panel", "osk", etc). The name_space string is copied, and caller
// maintians ownership of original. If the window is currently mapped, it will
// get remapped so the change can take effect.
//
// Default is "gtk-layer-shell" (which will be used if set to NULL).
func SetNamespace(window *gtk.Window, nameSpace string) {
	var _arg1 *C.GtkWindow // out
	var _arg2 *C.char      // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(nameSpace)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_layer_set_namespace(_arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(nameSpace)
}