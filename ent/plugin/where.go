// Code generated by ent, DO NOT EDIT.

package plugin

import (
	"PluginServer/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Plugin {
	return predicate.Plugin(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Plugin {
	return predicate.Plugin(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Plugin {
	return predicate.Plugin(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Plugin {
	return predicate.Plugin(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Plugin {
	return predicate.Plugin(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Plugin {
	return predicate.Plugin(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Plugin {
	return predicate.Plugin(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldDescription, v))
}

// Repo applies equality check predicate on the "repo" field. It's identical to RepoEQ.
func Repo(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldRepo, v))
}

// Homepage applies equality check predicate on the "homepage" field. It's identical to HomepageEQ.
func Homepage(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldHomepage, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldOwnerID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Plugin {
	return predicate.Plugin(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Plugin {
	return predicate.Plugin(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Plugin {
	return predicate.Plugin(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Plugin {
	return predicate.Plugin(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Plugin {
	return predicate.Plugin(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Plugin {
	return predicate.Plugin(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Plugin {
	return predicate.Plugin(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldContainsFold(FieldDescription, v))
}

// RepoEQ applies the EQ predicate on the "repo" field.
func RepoEQ(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldRepo, v))
}

// RepoNEQ applies the NEQ predicate on the "repo" field.
func RepoNEQ(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldNEQ(FieldRepo, v))
}

// RepoIn applies the In predicate on the "repo" field.
func RepoIn(vs ...string) predicate.Plugin {
	return predicate.Plugin(sql.FieldIn(FieldRepo, vs...))
}

// RepoNotIn applies the NotIn predicate on the "repo" field.
func RepoNotIn(vs ...string) predicate.Plugin {
	return predicate.Plugin(sql.FieldNotIn(FieldRepo, vs...))
}

// RepoGT applies the GT predicate on the "repo" field.
func RepoGT(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldGT(FieldRepo, v))
}

// RepoGTE applies the GTE predicate on the "repo" field.
func RepoGTE(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldGTE(FieldRepo, v))
}

// RepoLT applies the LT predicate on the "repo" field.
func RepoLT(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldLT(FieldRepo, v))
}

// RepoLTE applies the LTE predicate on the "repo" field.
func RepoLTE(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldLTE(FieldRepo, v))
}

// RepoContains applies the Contains predicate on the "repo" field.
func RepoContains(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldContains(FieldRepo, v))
}

// RepoHasPrefix applies the HasPrefix predicate on the "repo" field.
func RepoHasPrefix(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldHasPrefix(FieldRepo, v))
}

// RepoHasSuffix applies the HasSuffix predicate on the "repo" field.
func RepoHasSuffix(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldHasSuffix(FieldRepo, v))
}

// RepoIsNil applies the IsNil predicate on the "repo" field.
func RepoIsNil() predicate.Plugin {
	return predicate.Plugin(sql.FieldIsNull(FieldRepo))
}

// RepoNotNil applies the NotNil predicate on the "repo" field.
func RepoNotNil() predicate.Plugin {
	return predicate.Plugin(sql.FieldNotNull(FieldRepo))
}

// RepoEqualFold applies the EqualFold predicate on the "repo" field.
func RepoEqualFold(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEqualFold(FieldRepo, v))
}

// RepoContainsFold applies the ContainsFold predicate on the "repo" field.
func RepoContainsFold(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldContainsFold(FieldRepo, v))
}

// HomepageEQ applies the EQ predicate on the "homepage" field.
func HomepageEQ(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldHomepage, v))
}

// HomepageNEQ applies the NEQ predicate on the "homepage" field.
func HomepageNEQ(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldNEQ(FieldHomepage, v))
}

// HomepageIn applies the In predicate on the "homepage" field.
func HomepageIn(vs ...string) predicate.Plugin {
	return predicate.Plugin(sql.FieldIn(FieldHomepage, vs...))
}

// HomepageNotIn applies the NotIn predicate on the "homepage" field.
func HomepageNotIn(vs ...string) predicate.Plugin {
	return predicate.Plugin(sql.FieldNotIn(FieldHomepage, vs...))
}

// HomepageGT applies the GT predicate on the "homepage" field.
func HomepageGT(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldGT(FieldHomepage, v))
}

// HomepageGTE applies the GTE predicate on the "homepage" field.
func HomepageGTE(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldGTE(FieldHomepage, v))
}

// HomepageLT applies the LT predicate on the "homepage" field.
func HomepageLT(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldLT(FieldHomepage, v))
}

// HomepageLTE applies the LTE predicate on the "homepage" field.
func HomepageLTE(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldLTE(FieldHomepage, v))
}

// HomepageContains applies the Contains predicate on the "homepage" field.
func HomepageContains(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldContains(FieldHomepage, v))
}

// HomepageHasPrefix applies the HasPrefix predicate on the "homepage" field.
func HomepageHasPrefix(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldHasPrefix(FieldHomepage, v))
}

// HomepageHasSuffix applies the HasSuffix predicate on the "homepage" field.
func HomepageHasSuffix(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldHasSuffix(FieldHomepage, v))
}

// HomepageIsNil applies the IsNil predicate on the "homepage" field.
func HomepageIsNil() predicate.Plugin {
	return predicate.Plugin(sql.FieldIsNull(FieldHomepage))
}

// HomepageNotNil applies the NotNil predicate on the "homepage" field.
func HomepageNotNil() predicate.Plugin {
	return predicate.Plugin(sql.FieldNotNull(FieldHomepage))
}

// HomepageEqualFold applies the EqualFold predicate on the "homepage" field.
func HomepageEqualFold(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldEqualFold(FieldHomepage, v))
}

// HomepageContainsFold applies the ContainsFold predicate on the "homepage" field.
func HomepageContainsFold(v string) predicate.Plugin {
	return predicate.Plugin(sql.FieldContainsFold(FieldHomepage, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v uuid.UUID) predicate.Plugin {
	return predicate.Plugin(sql.FieldLTE(FieldOwnerID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Plugin) predicate.Plugin {
	return predicate.Plugin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Plugin) predicate.Plugin {
	return predicate.Plugin(func(s *sql.Selector) {
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
func Not(p predicate.Plugin) predicate.Plugin {
	return predicate.Plugin(func(s *sql.Selector) {
		p(s.Not())
	})
}
