// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"mangathrV2/ent/manga"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Manga is the model entity for the Manga schema.
type Manga struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MangaID holds the value of the "MangaID" field.
	MangaID string `json:"MangaID,omitempty"`
	// Source holds the value of the "Source" field.
	Source string `json:"Source,omitempty"`
	// Title holds the value of the "Title" field.
	Title string `json:"Title,omitempty"`
	// Mapping holds the value of the "Mapping" field.
	Mapping string `json:"Mapping,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MangaQuery when eager-loading is set.
	Edges MangaEdges `json:"edges"`
}

// MangaEdges holds the relations/edges for other nodes in the graph.
type MangaEdges struct {
	// Chapters holds the value of the Chapters edge.
	Chapters []*Chapter `json:"Chapters,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ChaptersOrErr returns the Chapters value or an error if the edge
// was not loaded in eager-loading.
func (e MangaEdges) ChaptersOrErr() ([]*Chapter, error) {
	if e.loadedTypes[0] {
		return e.Chapters, nil
	}
	return nil, &NotLoadedError{edge: "Chapters"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Manga) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case manga.FieldID:
			values[i] = new(sql.NullInt64)
		case manga.FieldMangaID, manga.FieldSource, manga.FieldTitle, manga.FieldMapping:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Manga", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Manga fields.
func (m *Manga) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case manga.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case manga.FieldMangaID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MangaID", values[i])
			} else if value.Valid {
				m.MangaID = value.String
			}
		case manga.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Source", values[i])
			} else if value.Valid {
				m.Source = value.String
			}
		case manga.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Title", values[i])
			} else if value.Valid {
				m.Title = value.String
			}
		case manga.FieldMapping:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Mapping", values[i])
			} else if value.Valid {
				m.Mapping = value.String
			}
		}
	}
	return nil
}

// QueryChapters queries the "Chapters" edge of the Manga entity.
func (m *Manga) QueryChapters() *ChapterQuery {
	return (&MangaClient{config: m.config}).QueryChapters(m)
}

// Update returns a builder for updating this Manga.
// Note that you need to call Manga.Unwrap() before calling this method if this Manga
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Manga) Update() *MangaUpdateOne {
	return (&MangaClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Manga entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Manga) Unwrap() *Manga {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Manga is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Manga) String() string {
	var builder strings.Builder
	builder.WriteString("Manga(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", MangaID=")
	builder.WriteString(m.MangaID)
	builder.WriteString(", Source=")
	builder.WriteString(m.Source)
	builder.WriteString(", Title=")
	builder.WriteString(m.Title)
	builder.WriteString(", Mapping=")
	builder.WriteString(m.Mapping)
	builder.WriteByte(')')
	return builder.String()
}

// Mangas is a parsable slice of Manga.
type Mangas []*Manga

func (m Mangas) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}