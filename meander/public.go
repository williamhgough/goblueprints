package meander

// Facade exposes Public method
type Facade interface {
	Public() interface{}
}

// Public allows type to expose itself differently
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
