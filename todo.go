package todo

type TodoList struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Desciption string `json:"desciption"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Desciption string `json:"desciption"`
	Done       bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	itemId int
}
