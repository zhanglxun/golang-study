package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Accounts holds the schema definition for the Accounts entity.
type Accounts struct {
	ent.Schema
}

// Fields of the Accounts.
func (Accounts) Fields() []ent.Field {

	return []ent.Field{

		field.Int32("id").SchemaType(map[string]string{
			dialect.MySQL: "int(20)", // Override MySQL.
		}).Unique(),

		field.String("account").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("pwd").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("nickname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("mobile").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),
	}

}

// Edges of the Accounts.
func (Accounts) Edges() []ent.Edge {
	return nil
}
