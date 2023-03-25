package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"time"
)

// CWebsite holds the schema definition for the CWebsite entity.
type CWebsite struct {
	ent.Schema
}

func (CWebsite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "c_website"},
	}
}

// Fields of the CWebsite.
func (CWebsite) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Comment("主键ID").Unique(),

		field.Int32("sort_id").SchemaType(map[string]string{
			dialect.MySQL: "int(8)", // Override MySQL.
		}).Optional().Comment("排序id"),

		field.Int32("category").SchemaType(map[string]string{
			dialect.MySQL: "int(2)", // Override MySQL.
		}).Optional().Comment("分类：1：工具和应用2：灵感和创作"),

		field.Int32("type").SchemaType(map[string]string{
			dialect.MySQL: "int(2)", // Override MySQL.
		}).Optional().Comment("分类：1：AI，2：工具"),

		field.String("website_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("网站的名称"),

		field.String("website_icon").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("网站icon地址"),

		field.String("website_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("网站的url"),

		field.String("summary").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)", // Override MySQL.
		}).Optional().Comment("简介"),

		field.String("description").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Optional().Comment("描述介绍"),

		field.Int64("create_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Optional().Comment("创建人"),

		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Default(time.Now).Comment("创建时间"),

		field.Int64("modify_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Optional().Comment("修改人"),

		field.Time("modify_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Default(time.Now).Comment("修改时间"),
	}

}

// Edges of the CWebsite.
func (CWebsite) Edges() []ent.Edge {
	return nil
}
