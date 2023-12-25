// Code generated by ent, DO NOT EDIT.

package verificationsession

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Encedeus/pluginServer/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldContainsFold(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldEQ(FieldUserID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.VerificationSession {
	return predicate.VerificationSession(sql.FieldNotIn(FieldUserID, vs...))
}

// HasSession applies the HasEdge predicate on the "session" edge.
func HasSession() predicate.VerificationSession {
	return predicate.VerificationSession(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SessionTable, SessionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSessionWith applies the HasEdge predicate on the "session" edge with a given conditions (other predicates).
func HasSessionWith(preds ...predicate.User) predicate.VerificationSession {
	return predicate.VerificationSession(func(s *sql.Selector) {
		step := newSessionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VerificationSession) predicate.VerificationSession {
	return predicate.VerificationSession(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VerificationSession) predicate.VerificationSession {
	return predicate.VerificationSession(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VerificationSession) predicate.VerificationSession {
	return predicate.VerificationSession(sql.NotPredicates(p))
}
