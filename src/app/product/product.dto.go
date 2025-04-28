package product

import "gin-boilerplate/src/database/dao"

// REQUEST
type CreateProductReq struct {
	Name string `json:"name" validate:"required,min=4" example:"item"`
	Price uint `json:"price" validate:"required,number" example:"3000"`
	Image string `json:"image" example:""`
}

func (v CreateProductReq) ToEntity(id uint) (r dao.ProductEntity) {
	r.Name = v.Name
	r.Price = v.Price
	if (v.Image != "") {
		r.Image = v.Image
	}
	r.StoreID = id
	return
}