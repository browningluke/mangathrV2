// Code generated by entc, DO NOT EDIT.

package chapter

const (
	// Label holds the string label denoting the chapter type in the database.
	Label = "chapter"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChapterID holds the string denoting the chapterid field in the database.
	FieldChapterID = "chapter_id"
	// FieldNum holds the string denoting the num field in the database.
	FieldNum = "num"
	// EdgeManga holds the string denoting the manga edge name in mutations.
	EdgeManga = "Manga"
	// Table holds the table name of the chapter in the database.
	Table = "chapters"
	// MangaTable is the table that holds the Manga relation/edge.
	MangaTable = "chapters"
	// MangaInverseTable is the table name for the Manga entity.
	// It exists in this package in order to avoid circular dependency with the "manga" package.
	MangaInverseTable = "mangas"
	// MangaColumn is the table column denoting the Manga relation/edge.
	MangaColumn = "manga_chapters"
)

// Columns holds all SQL columns for chapter fields.
var Columns = []string{
	FieldID,
	FieldChapterID,
	FieldNum,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "chapters"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"manga_chapters",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}