package query

type FieldType int

const (
	UnknownFieldType FieldType = iota
	Char
	Vachar
	Binary
	Varbinary
	Tinyblob
	Tinytext
	Text
	Blob
	Mediumtext
	Mediumblob
	Longtext
	Longblob
	// TODO: Add ENUM() and SET()

	Bit
	Tinyint
	Bool
	Boolean
	Smallint
	Mediumint
	Int
	Integer
	Bigint
	Float
	Double
	DoublePrecision
	Decimal
	Dec

	Date
	DateTime
	Timestamp
	Time
	Year
)

type Field struct {
	Name    string
	Type    FieldType
	Length  int
	Digit   int // For FLOAT, DOUBLE, DECIMAL, DEC type
	FSP     int // For DATETIME, TIMESTAMP, TIME
	NotNull bool
	Unique  bool
}
