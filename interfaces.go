// Generic interfaces
package main

type T interface {
	Eq(T) bool
}
type Array interface {
	ToS() S
}
