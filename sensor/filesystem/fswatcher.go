package filesystem

import "github.com/fsnotify/fsnotify"

type FsWatcher interface {
	Start() (<-chan *FsEvent, error)
	Stop() error
}

func NewFsWatcher(fullPath string) FsWatcher {

	return &fsWatcher{path: fullPath}
}

type fsWatcher struct {
	path    string
	watcher *fsnotify.Watcher
	event   chan *FsEvent
	running bool
}

func (w *fsWatcher) Start() (<-chan *FsEvent, error) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	w.watcher = watcher
	w.watcher.Add(w.path)
	if err != nil {
		watcher.Close()
		return nil, err
	}

	w.event = make(chan *FsEvent)
	w.running = true

	go w.startWatch()

	return w.event, nil
}

func (w *fsWatcher) Stop() error {

	w.running = false
	return w.watcher.Close()
}

func (w *fsWatcher) startWatch() {
	defer close(w.event)

	for w.running {
		select {
		case event := <-w.watcher.Events:
			w.event <- &FsEvent{&event, nil}
		case err := <-w.watcher.Errors:
			w.event <- &FsEvent{nil, err}
		}
	}
}
