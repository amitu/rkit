// Package cursor represents the UI cursor by the pointing device.
package cursor

// Cursor represents the mouse cursor. Not every device will support cursors,
// but those that will will have the following ones.
//
// In future we may support arbitrary bitmaps as cursors, but its not really
// priority, and they are so 2000s.
type Cursor int

// These are the various cursors available. Not all may be supported. As of now
// there is no way to programatically detect if one is supported or not, only
// way to find out is by testing it across different drivers
const (
	Default Cursor = iota
	Alias
	AllScroll
	Auto
	Cell
	ContextMenu
	ColResize
	Copy
	CrossHair
	EResize
	EWResize
	Grab
	Grabbing
	Help
	Move
	NResize
	NEResize
	NESWResize
	NSResize
	NWResize
	NWSEResize
	NoDrop
	None
	NotAllowed
	Pointer
	Progress
	RowResize
	SResize
	SEResize
	SWResize
	Text
	VerticalText
	WResize
	Wait
	ZoomIn
	ZoomOut
	Initial
	Inherit
)
