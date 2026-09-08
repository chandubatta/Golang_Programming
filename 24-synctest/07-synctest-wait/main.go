package wait

import "strings"

// ============================================================
// synctest.Wait
// ============================================================

type Uploader struct {
	jobs     chan string
	uploaded []string
}

func NewUploader() *Uploader {
	u := &Uploader{jobs: make(chan string)}
	go u.run()
	return u
}

func (u *Uploader) run() {
	for name := range u.jobs {
		u.uploaded = append(u.uploaded, strings.ToUpper(name))
	}
}

func (u *Uploader) Upload(name string) { u.jobs <- name }
func (u *Uploader) Close()             { close(u.jobs) }
func (u *Uploader) Uploaded() []string { return u.uploaded }
