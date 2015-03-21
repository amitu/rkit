// Command rbndl can be used to generate a .rbndl file from a bunch of resource.
//
//     $ rbndl data/:* /usr/share/:dict/words output.rbndl
//
// folder structure inside is deduced by the colon. part after : will be
// preserved, and must be unique across files. eg in this case if we had
// data/foo/img.gif, it will be stored as foo/gif and words file would be stored
// as dict/words. also: data/:* is equivalent to data if data is a folder.
//
// rbndl is a thin wrapper on top of tar, but does not need tar installed on
// system.
//
// other helpful commands:
//
//
//    $ rbndl -l file.rbndl
//
// ls -lh like output. this command also shows the command that was used to
// generate the archive.
//
//
//     $ rbndl -d file.rbndl
//
// shows a git status like output. for this command, it stores a meta file in
// the archive which contains the command line used to generate the .rbndl file.
// it uses that information to see what files should be there, and what files
// are actually present, and it shows which ones have been modified.
//
//
//     $ rbndl -u file.rbndl
//
// this command can be used to update the content of file.rbndl, as if original
// rbndl command was executed again.

package main

func main() {
}
