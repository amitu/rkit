/*
Package rkit implements a cross platform GUI toolkit.

rkit uses absolute bare minimum capability of underlying operating system, to
achieve maximum compatibility. rkit can be used to create apps in go that work
across different platforms.

Currently rkit supports gopherjs to build web apps. Plans include to support
desktop, android, ios and so on.

Text rendering, wrapping, image handling etc are all implemented in rkit without
support from native libraries, written in pure go. Based on these primitives a
widget library is built.

You can use the "github.com/amitu/rkit/widgets" to see all the widgets
available.

The basics of widgets: your GUI is composed of "widgets". There exists
rkit.Root, the root widget of main window. By default rkit.Root occupies the
whole Desktop in web/mobile platforms, and a configurable size in desktop
environments.

On desktop environments you may be able to launch other Window widgets.
rkit.Root is a rkit.Window, which happens to also be rkit.Widget.

You can create your own widgets by implementing rkit.Widget interface. You can
also "extend" an existing widget by embedding it in your widget. A widget must
only embed one other widget. Or none.

Your application will start by constructing your own main widget, and placing it
in rkit.Root widget using rkit.Root.Add(). In many cases the main widget would
be some form of layout widget.

LayoutWidgets are used to hold other widgets, and give those widgets a layout.
Some layouts could be horizontally stacked layout, or vertically stacked, or
generic layout, which has center, top, left, right, and bottom "portions".

Basic widgets for showing a constant text, Label, a line input, LineInput,
password input, PasswordInput, text area, TextArea, buttons, checkboxes, image
renderer, Image, drawing area, Canvas, exists in package rkit/widgets.

Internally there exists a "widget tree", rooted at each window. Each widget in
the root is drawn top to bottom, call to .Render() on any widget returns after
.Render() on each child has been carried out. Each .Render() returns a
image.Image, and root level image.Image is simply drawn on the underlying
canvas provided by operating system.
*/
package rkit
