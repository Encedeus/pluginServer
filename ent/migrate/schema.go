// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PluginsColumns holds the columns for the "plugins" table.
	PluginsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "owner_id", Type: field.TypeUUID},
		{Name: "source_id", Type: field.TypeInt},
	}
	// PluginsTable holds the schema information for the "plugins" table.
	PluginsTable = &schema.Table{
		Name:       "plugins",
		Columns:    PluginsColumns,
		PrimaryKey: []*schema.Column{PluginsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "plugins_users_owner",
				Columns:    []*schema.Column{PluginsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "plugins_sources_source",
				Columns:    []*schema.Column{PluginsColumns[3]},
				RefColumns: []*schema.Column{SourcesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PublicationsColumns holds the columns for the "publications" table.
	PublicationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "is_deprecated", Type: field.TypeBool, Default: false},
		{Name: "name", Type: field.TypeString, Size: 32},
		{Name: "uri_to_file", Type: field.TypeString},
		{Name: "plugin_id", Type: field.TypeUUID},
	}
	// PublicationsTable holds the schema information for the "publications" table.
	PublicationsTable = &schema.Table{
		Name:       "publications",
		Columns:    PublicationsColumns,
		PrimaryKey: []*schema.Column{PublicationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "publications_plugins_publications",
				Columns:    []*schema.Column{PublicationsColumns[5]},
				RefColumns: []*schema.Column{PluginsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SourcesColumns holds the columns for the "sources" table.
	SourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "repository", Type: field.TypeString},
	}
	// SourcesTable holds the schema information for the "sources" table.
	SourcesTable = &schema.Table{
		Name:       "sources",
		Columns:    SourcesColumns,
		PrimaryKey: []*schema.Column{SourcesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "auth_updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "email_verified", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PluginsTable,
		PublicationsTable,
		SourcesTable,
		UsersTable,
	}
)

func init() {
	PluginsTable.ForeignKeys[0].RefTable = UsersTable
	PluginsTable.ForeignKeys[1].RefTable = SourcesTable
	PublicationsTable.ForeignKeys[0].RefTable = PluginsTable
}
