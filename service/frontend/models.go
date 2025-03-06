package frontend

type Recommendation struct {
	Title       string `json:"title"`
	Product     string `json:"product"`
	Description string `json:"description"`
}

type Recommendations []Recommendation

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RecommendationsResponse struct {
	CustomerId      int              `json:"customerId"`
	CustomerName    string           `json:"customerName"`
	CustomerAge     int              `json:"customerAge"`
	Recommendations []Recommendation `json:"recommendations"`
}
