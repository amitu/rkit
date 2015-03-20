package cursor

/*
	The mouse cursor. Not every device will support cursors, but those that will
	will have the following ones.

	In future we may support arbitrary bitmaps as cursors, but its not really
	priority, and they are so 2000.
*/
type Cursor int

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
