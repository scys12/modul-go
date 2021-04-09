package product

import (
	"github.com/scys12/modul-go/internal/model"
	"github.com/scys12/modul-go/pkg/payload"
)

func (ps *ProductService) BuyProduct(productID int) (err error) {
	err = ps.productRepo.BuyProduct(productID)
	return err
}

func (ps *ProductService) DeleteProduct(productID, userID int) (err error) {
	err = ps.productRepo.DeleteProduct(productID, userID)
	return err
}

func (ps *ProductService) GetAvailableProducts() (products []model.Product, err error) {
	products, err = ps.productRepo.GetAvailableProducts()
	return
}

func (ps *ProductService) GetProduct(productID int) (product model.Product, err error) {
	product, err = ps.productRepo.GetProduct(productID)
	return
}

func (ps *ProductService) GetSellerProducts(sellerID int) (products []model.Product, err error) {
	products, err = ps.productRepo.GetSellerProducts(sellerID)
	return
}

func (ps *ProductService) InsertProduct(productReq payload.ProductRequest, userID int) (err error) {
	product := model.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
		SellerID:    userID,
	}
	err = ps.productRepo.InsertProduct(product)
	return
}

func (ps *ProductService) UpdateProduct(productReq payload.ProductRequest, productID, userID int) (err error) {
	updatedProduct, err := ps.productRepo.GetCurrentUserProduct(productID, userID)
	if err != nil {
		return
	}
	updatedProduct.Name = productReq.Name
	updatedProduct.Description = productReq.Description
	updatedProduct.Price = productReq.Price

	err = ps.productRepo.UpdateProduct(updatedProduct)
	return
}
