// Code generated by ent, DO NOT EDIT.

package publication

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Encedeus/pluginServer/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Publication {
	return predicate.Publication(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Publication {
	return predicate.Publication(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Publication {
	return predicate.Publication(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Publication {
	return predicate.Publication(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Publication {
	return predicate.Publication(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Publication {
	return predicate.Publication(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Publication {
	return predicate.Publication(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldCreatedAt, v))
}

// IsDeprecated applies equality check predicate on the "is_deprecated" field. It's identical to IsDeprecatedEQ.
func IsDeprecated(v bool) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldIsDeprecated, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldName, v))
}

// URIToFile applies equality check predicate on the "uri_to_file" field. It's identical to URIToFileEQ.
func URIToFile(v string) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldURIToFile, v))
}

// PluginID applies equality check predicate on the "plugin_id" field. It's identical to PluginIDEQ.
func PluginID(v uuid.UUID) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldPluginID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Publication {
	return predicate.Publication(sql.FieldLTE(FieldCreatedAt, v))
}

// IsDeprecatedEQ applies the EQ predicate on the "is_deprecated" field.
func IsDeprecatedEQ(v bool) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldIsDeprecated, v))
}

// IsDeprecatedNEQ applies the NEQ predicate on the "is_deprecated" field.
func IsDeprecatedNEQ(v bool) predicate.Publication {
	return predicate.Publication(sql.FieldNEQ(FieldIsDeprecated, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Publication {
	return predicate.Publication(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Publication {
	return predicate.Publication(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Publication {
	return predicate.Publication(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Publication {
	return predicate.Publication(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Publication {
	return predicate.Publication(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Publication {
	return predicate.Publication(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Publication {
	return predicate.Publication(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Publication {
	return predicate.Publication(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Publication {
	return predicate.Publication(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Publication {
	return predicate.Publication(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Publication {
	return predicate.Publication(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Publication {
	return predicate.Publication(sql.FieldContainsFold(FieldName, v))
}

// URIToFileEQ applies the EQ predicate on the "uri_to_file" field.
func URIToFileEQ(v string) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldURIToFile, v))
}

// URIToFileNEQ applies the NEQ predicate on the "uri_to_file" field.
func URIToFileNEQ(v string) predicate.Publication {
	return predicate.Publication(sql.FieldNEQ(FieldURIToFile, v))
}

// URIToFileIn applies the In predicate on the "uri_to_file" field.
func URIToFileIn(vs ...string) predicate.Publication {
	return predicate.Publication(sql.FieldIn(FieldURIToFile, vs...))
}

// URIToFileNotIn applies the NotIn predicate on the "uri_to_file" field.
func URIToFileNotIn(vs ...string) predicate.Publication {
	return predicate.Publication(sql.FieldNotIn(FieldURIToFile, vs...))
}

// URIToFileGT applies the GT predicate on the "uri_to_file" field.
func URIToFileGT(v string) predicate.Publication {
	return predicate.Publication(sql.FieldGT(FieldURIToFile, v))
}

// URIToFileGTE applies the GTE predicate on the "uri_to_file" field.
func URIToFileGTE(v string) predicate.Publication {
	return predicate.Publication(sql.FieldGTE(FieldURIToFile, v))
}

// URIToFileLT applies the LT predicate on the "uri_to_file" field.
func URIToFileLT(v string) predicate.Publication {
	return predicate.Publication(sql.FieldLT(FieldURIToFile, v))
}

// URIToFileLTE applies the LTE predicate on the "uri_to_file" field.
func URIToFileLTE(v string) predicate.Publication {
	return predicate.Publication(sql.FieldLTE(FieldURIToFile, v))
}

// URIToFileContains applies the Contains predicate on the "uri_to_file" field.
func URIToFileContains(v string) predicate.Publication {
	return predicate.Publication(sql.FieldContains(FieldURIToFile, v))
}

// URIToFileHasPrefix applies the HasPrefix predicate on the "uri_to_file" field.
func URIToFileHasPrefix(v string) predicate.Publication {
	return predicate.Publication(sql.FieldHasPrefix(FieldURIToFile, v))
}

// URIToFileHasSuffix applies the HasSuffix predicate on the "uri_to_file" field.
func URIToFileHasSuffix(v string) predicate.Publication {
	return predicate.Publication(sql.FieldHasSuffix(FieldURIToFile, v))
}

// URIToFileEqualFold applies the EqualFold predicate on the "uri_to_file" field.
func URIToFileEqualFold(v string) predicate.Publication {
	return predicate.Publication(sql.FieldEqualFold(FieldURIToFile, v))
}

// URIToFileContainsFold applies the ContainsFold predicate on the "uri_to_file" field.
func URIToFileContainsFold(v string) predicate.Publication {
	return predicate.Publication(sql.FieldContainsFold(FieldURIToFile, v))
}

// PluginIDEQ applies the EQ predicate on the "plugin_id" field.
func PluginIDEQ(v uuid.UUID) predicate.Publication {
	return predicate.Publication(sql.FieldEQ(FieldPluginID, v))
}

// PluginIDNEQ applies the NEQ predicate on the "plugin_id" field.
func PluginIDNEQ(v uuid.UUID) predicate.Publication {
	return predicate.Publication(sql.FieldNEQ(FieldPluginID, v))
}

// PluginIDIn applies the In predicate on the "plugin_id" field.
func PluginIDIn(vs ...uuid.UUID) predicate.Publication {
	return predicate.Publication(sql.FieldIn(FieldPluginID, vs...))
}

// PluginIDNotIn applies the NotIn predicate on the "plugin_id" field.
func PluginIDNotIn(vs ...uuid.UUID) predicate.Publication {
	return predicate.Publication(sql.FieldNotIn(FieldPluginID, vs...))
}

// HasPlugin applies the HasEdge predicate on the "plugin" edge.
func HasPlugin() predicate.Publication {
	return predicate.Publication(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PluginTable, PluginColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPluginWith applies the HasEdge predicate on the "plugin" edge with a given conditions (other predicates).
func HasPluginWith(preds ...predicate.Plugin) predicate.Publication {
	return predicate.Publication(func(s *sql.Selector) {
		step := newPluginStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Publication) predicate.Publication {
	return predicate.Publication(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Publication) predicate.Publication {
	return predicate.Publication(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Publication) predicate.Publication {
	return predicate.Publication(sql.NotPredicates(p))
}
