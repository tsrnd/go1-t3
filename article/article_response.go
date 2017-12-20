package article

// CommonResponse responses common json data.
type CommonResponse struct {
	Result  int      `json:"result"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

// ArticleGetResponse responses from Get method.
// JSON responses payload structure
type ArticleGetResponse struct {
	CommonResponse
	ResponseArticle []ArticleResponse `json:"article"`
}

// ArticleGetTitleResponse responses from GetID method.
// JSON responses payload structure
type ArticleGetTitleResponse struct {
	CommonResponse
	ResponseArticle []ArticleResponse `json:"article"`
}

// ArticleGetCountResponse responses from GetCount method.
// JSON responses payload structure
type ArticleGetCountResponse struct {
	CommonResponse
	Count int `json:"count"`
}

// ArticlePostCountResponse responses from PostCount method.
// JSON responses payload structure
type ArticlePostCountResponse struct {
	CommonResponse
	Count int `json:"count"`
}

// ArticleResponse responses struct.
// JSON responsess payload structure
type ArticleResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ArticlePostVisenzeDiscoversearchResponse responses from GetVisenzeDiscoversearch method.
// JSON responses payload structure
type ArticlePostVisenzeDiscoversearchResponse struct {
	CommonResponse
	VisenzeDiscoversearchResponse `json:discoversearch`
}

// VisenzeDiscoversearchResponse struct.
type VisenzeDiscoversearchResponse struct {
	Status          string                    `json:"status"`
	Objects         []ObjectResponse          `json:"objects"`
	ObjectTypesList []ObjectTypesListResponse `json:"object_types_list"`
}

// ObjectResponse struct.
type ObjectResponse struct {
	Type       string             `json:"type"`
	Attributes AttributesResponse `json:"attributes"`
	Score      float64            `json:"score"`
	Box        []int64            `json:"box"`
	Total      int64              `json:"total"`
	Result     []ResultResponse   `json:"result"`
}

// ResultResponse struct.
type ResultResponse struct {
	IMName   string           `json:"im_name"`
	Score    float64          `json:"score"`
	ValueMap ValueMapResponse `json:"value_map"`
}

// ValueMapResponse struct.
type ValueMapResponse struct {
	IMURL string `json:"im_url"`
}

// ObjectTypesListResponse struct.
type ObjectTypesListResponse struct {
	Type           string             `json:"type"`
	AttributesList AttributesResponse `json:"attributes_list,omitempty"`
}

// AttributesResponse struct.
type AttributesResponse struct {
}
