package jsonAPI

type ShowResponse struct {
	UserID int    `json:"userID"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string
}
