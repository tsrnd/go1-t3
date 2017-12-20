package article

import (
	"github.com/jinzhu/gorm"
)

// ----------------------------------------------------------
// Database
// ----------------------------------------------------------

// Article table struct.
// http://jinzhu.me/gorm/models.html#model-definition
type Article struct {
	gorm.Model
	Title   string `gorm:"column:title;type:varchar(45);not null;"`
	Content string `gorm:"column:content;type:text;not null;column:content"`
}

// ----------------------------------------------------------
// API:Visenze
// struct aute generate: https://app.quicktype.io/#l=go
// ----------------------------------------------------------

// VisenzeDiscoversearch struct.
type VisenzeDiscoversearch struct {
	Status          string            `json:"status"`
	Method          string            `json:"method"`
	Error           []interface{}     `json:"error,omitempty"`
	ResultLimit     int64             `json:"result_limit"`
	DetectionLimit  int64             `json:"detection_limit"`
	Page            int64             `json:"page"`
	Objects         []Object          `json:"objects"`
	ObjectTypesList []ObjectTypesList `json:"object_types_list"`
	IMID            string            `json:"im_id"`
	Reqid           string            `json:"reqid"`
}

// Object struct.
type Object struct {
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
	Score      float64    `json:"score"`
	Box        []int64    `json:"box"`
	Total      int64      `json:"total"`
	Result     []Result   `json:"result"`
}

// Result struct.
type Result struct {
	IMName   string   `json:"im_name"`
	Score    float64  `json:"score"`
	ValueMap ValueMap `json:"value_map"`
}

// ValueMap struct.
type ValueMap struct {
	IMURL string `json:"im_url"`
}

// ObjectTypesList struct.
type ObjectTypesList struct {
	Type           string     `json:"type"`
	AttributesList Attributes `json:"attributes_list,omitempty"`
}

// Attributes struct.
type Attributes struct {
}
