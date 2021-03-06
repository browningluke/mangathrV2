// Code generated by entc, DO NOT EDIT.

package chapter

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/browningluke/mangathrV2/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ChapterID applies equality check predicate on the "ChapterID" field. It's identical to ChapterIDEQ.
func ChapterID(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChapterID), v))
	})
}

// Num applies equality check predicate on the "Num" field. It's identical to NumEQ.
func Num(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNum), v))
	})
}

// Title applies equality check predicate on the "Title" field. It's identical to TitleEQ.
func Title(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// CreatedOn applies equality check predicate on the "CreatedOn" field. It's identical to CreatedOnEQ.
func CreatedOn(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedOn), v))
	})
}

// RegisteredOn applies equality check predicate on the "RegisteredOn" field. It's identical to RegisteredOnEQ.
func RegisteredOn(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisteredOn), v))
	})
}

// ChapterIDEQ applies the EQ predicate on the "ChapterID" field.
func ChapterIDEQ(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChapterID), v))
	})
}

// ChapterIDNEQ applies the NEQ predicate on the "ChapterID" field.
func ChapterIDNEQ(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChapterID), v))
	})
}

// ChapterIDIn applies the In predicate on the "ChapterID" field.
func ChapterIDIn(vs ...string) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChapterID), v...))
	})
}

// ChapterIDNotIn applies the NotIn predicate on the "ChapterID" field.
func ChapterIDNotIn(vs ...string) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChapterID), v...))
	})
}

// ChapterIDGT applies the GT predicate on the "ChapterID" field.
func ChapterIDGT(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChapterID), v))
	})
}

// ChapterIDGTE applies the GTE predicate on the "ChapterID" field.
func ChapterIDGTE(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChapterID), v))
	})
}

// ChapterIDLT applies the LT predicate on the "ChapterID" field.
func ChapterIDLT(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChapterID), v))
	})
}

// ChapterIDLTE applies the LTE predicate on the "ChapterID" field.
func ChapterIDLTE(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChapterID), v))
	})
}

// ChapterIDContains applies the Contains predicate on the "ChapterID" field.
func ChapterIDContains(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChapterID), v))
	})
}

// ChapterIDHasPrefix applies the HasPrefix predicate on the "ChapterID" field.
func ChapterIDHasPrefix(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChapterID), v))
	})
}

// ChapterIDHasSuffix applies the HasSuffix predicate on the "ChapterID" field.
func ChapterIDHasSuffix(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChapterID), v))
	})
}

// ChapterIDEqualFold applies the EqualFold predicate on the "ChapterID" field.
func ChapterIDEqualFold(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChapterID), v))
	})
}

// ChapterIDContainsFold applies the ContainsFold predicate on the "ChapterID" field.
func ChapterIDContainsFold(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChapterID), v))
	})
}

// NumEQ applies the EQ predicate on the "Num" field.
func NumEQ(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNum), v))
	})
}

// NumNEQ applies the NEQ predicate on the "Num" field.
func NumNEQ(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNum), v))
	})
}

// NumIn applies the In predicate on the "Num" field.
func NumIn(vs ...string) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNum), v...))
	})
}

// NumNotIn applies the NotIn predicate on the "Num" field.
func NumNotIn(vs ...string) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNum), v...))
	})
}

// NumGT applies the GT predicate on the "Num" field.
func NumGT(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNum), v))
	})
}

// NumGTE applies the GTE predicate on the "Num" field.
func NumGTE(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNum), v))
	})
}

// NumLT applies the LT predicate on the "Num" field.
func NumLT(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNum), v))
	})
}

// NumLTE applies the LTE predicate on the "Num" field.
func NumLTE(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNum), v))
	})
}

// NumContains applies the Contains predicate on the "Num" field.
func NumContains(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNum), v))
	})
}

// NumHasPrefix applies the HasPrefix predicate on the "Num" field.
func NumHasPrefix(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNum), v))
	})
}

// NumHasSuffix applies the HasSuffix predicate on the "Num" field.
func NumHasSuffix(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNum), v))
	})
}

// NumEqualFold applies the EqualFold predicate on the "Num" field.
func NumEqualFold(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNum), v))
	})
}

// NumContainsFold applies the ContainsFold predicate on the "Num" field.
func NumContainsFold(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNum), v))
	})
}

// TitleEQ applies the EQ predicate on the "Title" field.
func TitleEQ(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "Title" field.
func TitleNEQ(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "Title" field.
func TitleIn(vs ...string) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "Title" field.
func TitleNotIn(vs ...string) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "Title" field.
func TitleGT(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "Title" field.
func TitleGTE(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "Title" field.
func TitleLT(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "Title" field.
func TitleLTE(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "Title" field.
func TitleContains(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "Title" field.
func TitleHasPrefix(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "Title" field.
func TitleHasSuffix(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleIsNil applies the IsNil predicate on the "Title" field.
func TitleIsNil() predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTitle)))
	})
}

// TitleNotNil applies the NotNil predicate on the "Title" field.
func TitleNotNil() predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTitle)))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "Title" field.
func TitleEqualFold(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "Title" field.
func TitleContainsFold(v string) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// CreatedOnEQ applies the EQ predicate on the "CreatedOn" field.
func CreatedOnEQ(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedOn), v))
	})
}

// CreatedOnNEQ applies the NEQ predicate on the "CreatedOn" field.
func CreatedOnNEQ(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedOn), v))
	})
}

// CreatedOnIn applies the In predicate on the "CreatedOn" field.
func CreatedOnIn(vs ...time.Time) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedOn), v...))
	})
}

// CreatedOnNotIn applies the NotIn predicate on the "CreatedOn" field.
func CreatedOnNotIn(vs ...time.Time) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedOn), v...))
	})
}

// CreatedOnGT applies the GT predicate on the "CreatedOn" field.
func CreatedOnGT(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedOn), v))
	})
}

// CreatedOnGTE applies the GTE predicate on the "CreatedOn" field.
func CreatedOnGTE(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedOn), v))
	})
}

// CreatedOnLT applies the LT predicate on the "CreatedOn" field.
func CreatedOnLT(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedOn), v))
	})
}

// CreatedOnLTE applies the LTE predicate on the "CreatedOn" field.
func CreatedOnLTE(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedOn), v))
	})
}

// CreatedOnIsNil applies the IsNil predicate on the "CreatedOn" field.
func CreatedOnIsNil() predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedOn)))
	})
}

// CreatedOnNotNil applies the NotNil predicate on the "CreatedOn" field.
func CreatedOnNotNil() predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedOn)))
	})
}

// RegisteredOnEQ applies the EQ predicate on the "RegisteredOn" field.
func RegisteredOnEQ(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisteredOn), v))
	})
}

// RegisteredOnNEQ applies the NEQ predicate on the "RegisteredOn" field.
func RegisteredOnNEQ(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisteredOn), v))
	})
}

// RegisteredOnIn applies the In predicate on the "RegisteredOn" field.
func RegisteredOnIn(vs ...time.Time) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisteredOn), v...))
	})
}

// RegisteredOnNotIn applies the NotIn predicate on the "RegisteredOn" field.
func RegisteredOnNotIn(vs ...time.Time) predicate.Chapter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chapter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisteredOn), v...))
	})
}

// RegisteredOnGT applies the GT predicate on the "RegisteredOn" field.
func RegisteredOnGT(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisteredOn), v))
	})
}

// RegisteredOnGTE applies the GTE predicate on the "RegisteredOn" field.
func RegisteredOnGTE(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisteredOn), v))
	})
}

// RegisteredOnLT applies the LT predicate on the "RegisteredOn" field.
func RegisteredOnLT(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisteredOn), v))
	})
}

// RegisteredOnLTE applies the LTE predicate on the "RegisteredOn" field.
func RegisteredOnLTE(v time.Time) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisteredOn), v))
	})
}

// HasManga applies the HasEdge predicate on the "Manga" edge.
func HasManga() predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MangaTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MangaTable, MangaColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMangaWith applies the HasEdge predicate on the "Manga" edge with a given conditions (other predicates).
func HasMangaWith(preds ...predicate.Manga) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MangaInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MangaTable, MangaColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Chapter) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Chapter) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Chapter) predicate.Chapter {
	return predicate.Chapter(func(s *sql.Selector) {
		p(s.Not())
	})
}
