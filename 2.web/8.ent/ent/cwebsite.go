// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"golangStudy/2.web/8.ent/ent/cwebsite"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// CWebsite is the model entity for the CWebsite schema.
type CWebsite struct {
	config `json:"-"`
	// ID of the ent.
	// 主键ID
	ID int64 `json:"id,omitempty"`
	// 排序id
	SortID int32 `json:"sort_id,omitempty"`
	// 分类：1：工具和应用2：灵感和创作
	Category int32 `json:"category,omitempty"`
	// 分类：1：AI，2：工具
	Type int32 `json:"type,omitempty"`
	// 网站的名称
	WebsiteName string `json:"website_name,omitempty"`
	// 网站icon地址
	WebsiteIcon string `json:"website_icon,omitempty"`
	// 网站的url
	WebsiteURL string `json:"website_url,omitempty"`
	// 简介
	Summary string `json:"summary,omitempty"`
	// 描述介绍
	Description string `json:"description,omitempty"`
	// 创建人
	CreateID int64 `json:"create_id,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改人
	ModifyID int64 `json:"modify_id,omitempty"`
	// 修改时间
	ModifyTime time.Time `json:"modify_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CWebsite) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cwebsite.FieldID, cwebsite.FieldSortID, cwebsite.FieldCategory, cwebsite.FieldType, cwebsite.FieldCreateID, cwebsite.FieldModifyID:
			values[i] = new(sql.NullInt64)
		case cwebsite.FieldWebsiteName, cwebsite.FieldWebsiteIcon, cwebsite.FieldWebsiteURL, cwebsite.FieldSummary, cwebsite.FieldDescription:
			values[i] = new(sql.NullString)
		case cwebsite.FieldCreateTime, cwebsite.FieldModifyTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CWebsite", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CWebsite fields.
func (c *CWebsite) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cwebsite.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case cwebsite.FieldSortID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort_id", values[i])
			} else if value.Valid {
				c.SortID = int32(value.Int64)
			}
		case cwebsite.FieldCategory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				c.Category = int32(value.Int64)
			}
		case cwebsite.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = int32(value.Int64)
			}
		case cwebsite.FieldWebsiteName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_name", values[i])
			} else if value.Valid {
				c.WebsiteName = value.String
			}
		case cwebsite.FieldWebsiteIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_icon", values[i])
			} else if value.Valid {
				c.WebsiteIcon = value.String
			}
		case cwebsite.FieldWebsiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_url", values[i])
			} else if value.Valid {
				c.WebsiteURL = value.String
			}
		case cwebsite.FieldSummary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field summary", values[i])
			} else if value.Valid {
				c.Summary = value.String
			}
		case cwebsite.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case cwebsite.FieldCreateID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_id", values[i])
			} else if value.Valid {
				c.CreateID = value.Int64
			}
		case cwebsite.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case cwebsite.FieldModifyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field modify_id", values[i])
			} else if value.Valid {
				c.ModifyID = value.Int64
			}
		case cwebsite.FieldModifyTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modify_time", values[i])
			} else if value.Valid {
				c.ModifyTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CWebsite.
// Note that you need to call CWebsite.Unwrap() before calling this method if this CWebsite
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *CWebsite) Update() *CWebsiteUpdateOne {
	return NewCWebsiteClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the CWebsite entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *CWebsite) Unwrap() *CWebsite {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: CWebsite is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *CWebsite) String() string {
	var builder strings.Builder
	builder.WriteString("CWebsite(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("sort_id=")
	builder.WriteString(fmt.Sprintf("%v", c.SortID))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", c.Category))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", ")
	builder.WriteString("website_name=")
	builder.WriteString(c.WebsiteName)
	builder.WriteString(", ")
	builder.WriteString("website_icon=")
	builder.WriteString(c.WebsiteIcon)
	builder.WriteString(", ")
	builder.WriteString("website_url=")
	builder.WriteString(c.WebsiteURL)
	builder.WriteString(", ")
	builder.WriteString("summary=")
	builder.WriteString(c.Summary)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("create_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CreateID))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modify_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ModifyID))
	builder.WriteString(", ")
	builder.WriteString("modify_time=")
	builder.WriteString(c.ModifyTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CWebsites is a parsable slice of CWebsite.
type CWebsites []*CWebsite
