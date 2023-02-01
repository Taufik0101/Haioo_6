package dto

type CreateCart struct {
	KodeProduk string `json:"kodeProduk" form:"kodeProduk" binding:"required"`
	NamaProduk string `json:"namaProduk" form:"namaProduk" binding:"required"`
	Kuantitas  int64  `json:"kuantitas" form:"kuantitas" binding:"required"`
}

type ListCartInput struct {
	Filter *FilterCartInput `json:"filter,omitempty" form:"filter"`
}

type FilterCartInput struct {
	NamaProduk *string `json:"namaProduk,omitempty" form:"namaProduk"`
	Kuantitas  *int64  `json:"kuantitas,omitempty" form:"kuantitas"`
}

type ListCartResponse struct {
	Value string `json:"value" form:"value"`
}
