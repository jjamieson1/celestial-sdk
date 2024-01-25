package models

import "time"

type Comment struct {
	CommentId	int 	`json:"comment_id"`
	UserId 		string	`json:"user_id"`
	ItemId		string	`json:"item_id"`
	Title		string	`json:"title"`
	Description	string	`json:"description"`
	Rating		int		`json:"rating"`
	TypeId		int		`json:"type_id"`
	Approved	bool	`json:"approved"`
	CreatedDate	time.Time	`json:"createdDate"`
}
