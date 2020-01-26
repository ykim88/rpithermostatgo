package filesystem

import "github.com/fsnotify/fsnotify"

type FsEvent struct {
	Event *fsnotify.Event
	Error error
}
