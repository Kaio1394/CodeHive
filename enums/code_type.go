package enums

type CodeType int

const (
	Sql CodeType = iota
	Csharp
	Golang
	Python
	Java
	Groovy
	Rust
	C
	Cplusplus
)

func (c CodeType) String() string {
	switch c {
	case Sql:
		return "Sql"
	case Csharp:
		return "C#"
	case Golang:
		return "Golang"
	case Java:
		return "Java"
	case Python:
		return "Python"
	case Groovy:
		return "Groovy"
	case Rust:
		return "Rust"
	case Cplusplus:
		return "Cplusplus"
	case C:
		return "C"
	default:
		return "Unknown"
	}
}
