# Notes

## Nil

In Go, an interface value is considered non-nil if it has both a type and a value. If we create a variable with the type of an interface and initialize it with a nil value, the variable will not be nil.

```
var ptrStruct *MyStruct = (*MyStruct)(nil)
var ptrInterface MyInterface = ptrStruct    
fmt.Println(ptrStruct)  
fmt.Println(ptrInterface)   
fmt.Println(ptrStruct == nil)    // true
fmt.Println(ptrInterface != nil) // true
```
Here we create a pointer to MyStruct and initializing it with nil. This makes `ptrStruct` a nil pointer.
However, when we assign `ptrStruct` to the interface variable `ptrInterface`, the interface variable `ptrInterface` still __contains type information__, which is *MyStruct, even though the underlying value (the pointer) is nil. This is why `ptrInterface` is not nil in terms of its type.

## Implicit Interface Implementation
Also called type-safe duck typing. In Go, we don't need to explicitly declare that a type implements an interface. If a type has all the methods of an interface, it automatically implements that interface.
