package snippet

type Snippet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Complected  bool   `json:"complected"`
}
