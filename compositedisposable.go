package atom

import "github.com/gopherjs/gopherjs/js"

// CompositeDisposable doc at: https://atom.io/docs/api/v0.165.0/CompositeDisposable
type CompositeDisposable struct {
	o js.Object
}

// NewCompositeDisposable doc at: https://atom.io/docs/api/v0.165.0/CompositeDisposable#instance-constructor
// This is a goified constructor, it dose the same.
func NewCompositeDisposable(disposables ...Disposable) *CompositeDisposable {
	cd := &CompositeDisposable{}
	cd.o = atomApi.Get("CompositeDisposable").New()
	print(cd.o)
	for _, disposable := range disposables {
		cd.Add(disposable)
	}
	print(cd.o)
	return cd
}

// Dispose doc at: https://atom.io/docs/api/v0.165.0/CompositeDisposable#instance-dispose
func (c *CompositeDisposable) Dispose() {
	c.o.Call("dispose")
}

// Add doc at: https://atom.io/docs/api/v0.165.0/CompositeDisposable#instance-add
func (c *CompositeDisposable) Add(disposable Disposable) {
	o := disposable.getJsObject()
	c.o.Call("add", o)
}

// Remove doc at: https://atom.io/docs/api/v0.165.0/CompositeDisposable#instance-remove
func (c *CompositeDisposable) Remove(dispoable Disposable) {
	o := dispoable.getJsObject()
	c.o.Call("remove", o)
}

// Clear doc at: https://atom.io/docs/api/v0.165.0/CompositeDisposable#instance-clear
func (c *CompositeDisposable) Clear() {
	c.o.Call("clear")
}
