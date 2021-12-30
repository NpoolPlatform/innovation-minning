// Code generated by entc, DO NOT EDIT.

package capital

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Introduction applies equality check predicate on the "introduction" field. It's identical to IntroductionEQ.
func Introduction(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntroduction), v))
	})
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// IntroductionEQ applies the EQ predicate on the "introduction" field.
func IntroductionEQ(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntroduction), v))
	})
}

// IntroductionNEQ applies the NEQ predicate on the "introduction" field.
func IntroductionNEQ(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIntroduction), v))
	})
}

// IntroductionIn applies the In predicate on the "introduction" field.
func IntroductionIn(vs ...string) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIntroduction), v...))
	})
}

// IntroductionNotIn applies the NotIn predicate on the "introduction" field.
func IntroductionNotIn(vs ...string) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIntroduction), v...))
	})
}

// IntroductionGT applies the GT predicate on the "introduction" field.
func IntroductionGT(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIntroduction), v))
	})
}

// IntroductionGTE applies the GTE predicate on the "introduction" field.
func IntroductionGTE(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIntroduction), v))
	})
}

// IntroductionLT applies the LT predicate on the "introduction" field.
func IntroductionLT(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIntroduction), v))
	})
}

// IntroductionLTE applies the LTE predicate on the "introduction" field.
func IntroductionLTE(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIntroduction), v))
	})
}

// IntroductionContains applies the Contains predicate on the "introduction" field.
func IntroductionContains(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIntroduction), v))
	})
}

// IntroductionHasPrefix applies the HasPrefix predicate on the "introduction" field.
func IntroductionHasPrefix(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIntroduction), v))
	})
}

// IntroductionHasSuffix applies the HasSuffix predicate on the "introduction" field.
func IntroductionHasSuffix(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIntroduction), v))
	})
}

// IntroductionEqualFold applies the EqualFold predicate on the "introduction" field.
func IntroductionEqualFold(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIntroduction), v))
	})
}

// IntroductionContainsFold applies the ContainsFold predicate on the "introduction" field.
func IntroductionContainsFold(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIntroduction), v))
	})
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogo), v))
	})
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogo), v...))
	})
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogo), v...))
	})
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogo), v))
	})
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogo), v))
	})
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogo), v))
	})
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogo), v))
	})
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogo), v))
	})
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogo), v))
	})
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogo), v))
	})
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogo), v))
	})
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogo), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.Capital {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Capital(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Capital) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Capital) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
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
func Not(p predicate.Capital) predicate.Capital {
	return predicate.Capital(func(s *sql.Selector) {
		p(s.Not())
	})
}