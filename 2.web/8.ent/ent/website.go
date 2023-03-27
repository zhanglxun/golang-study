// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"golangStudy/2.web/8.ent/ent/website"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Website is the model entity for the Website schema.
type Website struct {
	config `json:"-"`
	// ID of the ent.
	// 主键ID
	ID int32 `json:"id,omitempty"`
	// 主键ID
	SortID int8 `json:"sort_id,omitempty"`
	// 分类的ID
	Category int8 `json:"category,omitempty"`
	// 网站名称
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
	CreateID string `json:"create_id,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 修改人
	ModifyID string `json:"modify_id,omitempty"`
	// 修改时间
	ModifyTime time.Time `json:"modify_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Website) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case website.FieldID, website.FieldSortID, website.FieldCategory:
			values[i] = new(sql.NullInt64)
		case website.FieldWebsiteName, website.FieldWebsiteIcon, website.FieldWebsiteURL, website.FieldSummary, website.FieldDescription, website.FieldCreateID, website.FieldModifyID:
			values[i] = new(sql.NullString)
		case website.FieldCreateTime, website.FieldModifyTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Website", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Website fields.
func (w *Website) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case website.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int32(value.Int64)
		case website.FieldSortID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort_id", values[i])
			} else if value.Valid {
				w.SortID = int8(value.Int64)
			}
		case website.FieldCategory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				w.Category = int8(value.Int64)
			}
		case website.FieldWebsiteName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_name", values[i])
			} else if value.Valid {
				w.WebsiteName = value.String
			}
		case website.FieldWebsiteIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_icon", values[i])
			} else if value.Valid {
				w.WebsiteIcon = value.String
			}
		case website.FieldWebsiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_url", values[i])
			} else if value.Valid {
				w.WebsiteURL = value.String
			}
		case website.FieldSummary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field summary", values[i])
			} else if value.Valid {
				w.Summary = value.String
			}
		case website.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				w.Description = value.String
			}
		case website.FieldCreateID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field create_id", values[i])
			} else if value.Valid {
				w.CreateID = value.String
			}
		case website.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				w.CreateTime = value.Time
			}
		case website.FieldModifyID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field modify_id", values[i])
			} else if value.Valid {
				w.ModifyID = value.String
			}
		case website.FieldModifyTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modify_time", values[i])
			} else if value.Valid {
				w.ModifyTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Website.
// Note that you need to call Website.Unwrap() before calling this method if this Website
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Website) Update() *WebsiteUpdateOne {
	return NewWebsiteClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Website entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Website) Unwrap() *Website {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Website is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Website) String() string {
	var builder strings.Builder
	builder.WriteString("Website(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("sort_id=")
	builder.WriteString(fmt.Sprintf("%v", w.SortID))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", w.Category))
	builder.WriteString(", ")
	builder.WriteString("website_name=")
	builder.WriteString(w.WebsiteName)
	builder.WriteString(", ")
	builder.WriteString("website_icon=")
	builder.WriteString(w.WebsiteIcon)
	builder.WriteString(", ")
	builder.WriteString("website_url=")
	builder.WriteString(w.WebsiteURL)
	builder.WriteString(", ")
	builder.WriteString("summary=")
	builder.WriteString(w.Summary)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(w.Description)
	builder.WriteString(", ")
	builder.WriteString("create_id=")
	builder.WriteString(w.CreateID)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(w.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modify_id=")
	builder.WriteString(w.ModifyID)
	builder.WriteString(", ")
	builder.WriteString("modify_time=")
	builder.WriteString(w.ModifyTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Websites is a parsable slice of Website.
type Websites []*Website
