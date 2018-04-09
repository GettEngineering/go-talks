package interfaces

// CODE1 BEGIN OMIT
type AnimalInterface interface {
	Say()
}

type Dog struct {
	AnimalInterface // <- Don't do like this // HL
	name string
}

func (d *Dog) Say() { fmt.Println("Woof!") }

// CODE1 END OMIT

// CODE2 BEGIN OMIT
...
func AnimalSay(animal AnimalInterface) { animal.Say() }
...
// CODE2 END OMIT

// CODE3 BEGIN OMIT
$ go build 3/interfaces3.go
$
// CODE3 END OMIT

// CODE4 BEGIN OMIT
$ ./interfaces3
Dog: {AnimalInterface:<nil> name:Snow}
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x10951a2]

goroutine 1 [running]:
main.main()
	/Users/kryabov/Documents/Work/mini_talks/3/interfaces3.go:21 +0xc2
$
// CODE4 END OMIT

// CODE5 BEGIN OMIT
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
// CODE5 END OMIT

// CODE5.1 BEGIN OMIT
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// CODE5.1 END OMIT

// CODE6 BEGIN OMIT
type reverse struct {
        // This embedded Interface permits Reverse to use the methods of
        // another Interface implementation.
        Interface
}

// Less returns the opposite of the embedded implementation's Less method.
func (r reverse) Less(i, j int) bool {
        return r.Interface.Less(j, i)
}
// CODE6 END OMIT

// CODE7 BEGIN OMIT
func Reverse(data Interface) Interface {
        return &reverse{data}
}
// CODE7 END OMIT

