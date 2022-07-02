package jsonAPI

type Post struct {
	UserID   int    `json:"userID"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string
	Comments []Comment
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Comment struct {
	PostID int `json:"postId"`
	ID     int `json:"id"`
	Name   string
	Email  string
	Body   string
}

type UserData struct {
	UserData User
	PostData Post
}
