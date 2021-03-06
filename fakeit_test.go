package fakeit

import (
	"fmt"

	"github.com/go-openapi/strfmt"
)

type StatusVal int

type BigLebowski struct {
	Name  string
	Age   int
	Derps strfmt.CreditCard
	Slice []strfmt.Password
}

type testStruct struct {
	Date       strfmt.Date
	DateTime   strfmt.DateTime
	Duration   strfmt.Duration
	Email      strfmt.Email
	Base64     strfmt.Base64
	URI        strfmt.URI
	Hostname   strfmt.Hostname
	IPv4       strfmt.IPv4
	IPv6       strfmt.IPv6
	MAC        strfmt.MAC
	UUID       strfmt.UUID
	UUID3      strfmt.UUID3
	UUID4      strfmt.UUID4
	UUID5      strfmt.UUID5
	ISBN       strfmt.ISBN
	ISBN10     strfmt.ISBN10
	ISBN13     strfmt.ISBN13
	CreditCard strfmt.CreditCard
	SSN        strfmt.SSN
	HexColor   strfmt.HexColor
	RGBColor   strfmt.RGBColor
	Password   strfmt.Password

	Bool   bool
	String string

	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Uintptr    uintptr
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128

	Array    [2]strfmt.Password
	PtrArray [2]*strfmt.Password
	PtrSlice []*strfmt.Password
	Slice    []BigLebowski

	PtrMap map[string]*BigLebowski
	Map    map[string]BigLebowski

	// Chan chan int

	Interface interface{}

	Ptr       *int
	PtrStruct *BigLebowski
	Struct    BigLebowski

	StatusVal StatusVal
}

type DunWork struct {
	Ptr       *int
	PtrStruct *BigLebowski
	Array     [2]strfmt.Password
	Slice     []BigLebowski
	PtrSlice  []*strfmt.Password
	PtrArray  [2]*strfmt.Password

	Map    map[string]BigLebowski
	MapPtr map[string]*BigLebowski
}

func main() {
	// test := DunWork{}
	test := testStruct{}
	FillStruct(&test)
	fmt.Println("main")
	fmt.Printf("%+v\n", test)
}
