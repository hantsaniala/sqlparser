package query

// Query represents a parsed query
type Query struct {
	Type       Type
	TableName  string
	Conditions []Condition
	Updates    map[string]string
	Inserts    [][]string
	Fields     []string // Used for SELECT (i.e. SELECTed field names) and INSERT (INSERTEDed field names)
	Aliases    map[string]string
}

// Type is the type of SQL query, e.g. SELECT/UPDATE
type Type int

const (
	// UnknownType is the zero value for a Type
	UnknownType Type = iota
	// Select represents a SELECT query
	Select
	// Update represents an UPDATE query
	Update
	// Insert represents an INSERT query
	Insert
	// Delete represents a DELETE query
	Delete
)

// TypeString is a string slice with the names of all types in order
var TypeString = []string{
	"UnknownType",
	"Select",
	"Update",
	"Insert",
	"Delete",
}
