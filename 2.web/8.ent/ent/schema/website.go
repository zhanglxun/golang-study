package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)
import "entgo.io/ent/schema/field"

// Website holds the schema definition for the Website entity.
type Website struct {
	ent.Schema
}

func (Website) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "website"},
	}
}

// Fields of the Website.
func (Website) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Comment("主键ID"),
		field.Int8("sort_id").Comment("主键ID"),
		field.Int8("category").Comment("分类的ID"),
		field.String("website_name").Comment("网站名称"),
		field.String("website_icon").Comment("网站icon地址"),
		field.String("website_url").Comment("网站的url"),
		field.String("summary").Comment("简介"),
		field.String("description").Comment("描述介绍"),

		field.String("create_id").Comment("创建人"),
		field.Time("create_time").Comment("创建时间"),
		field.String("modify_id").Comment("修改人"),
		field.Time("modify_time").Comment("修改时间"),
	}
}

// Edges of the Website.
func (Website) Edges() []ent.Edge {
	return nil
}
