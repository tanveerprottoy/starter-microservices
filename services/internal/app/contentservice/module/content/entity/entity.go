package entity

type Content struct {
	Id        string `db:"_id" json:"id"`
	Name      string `db:"name" json:"name"`
	CreatedAt int64  `db:"createdAt" json:"createdAt"`
	UpdatedAt int64  `db:"updatedAt" json:"updatedAt"`
	UserId    string `db:"userId" json:"userId"`
}
