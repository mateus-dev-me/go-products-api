package controllers

import (
	"net/http"
	"strconv"

	"go-products/internal/domain"
	"go-products/internal/infrastructure/db"
	"go-products/internal/use_cases"

	"github.com/gin-gonic/gin"
)

type productController struct {
	GetAll  *use_cases.GetAllProductsUseCase
	GetByID *use_cases.GetProductByIDUseCase
	Save    *use_cases.SaveProductUseCase
	Update  *use_cases.UpdateProductUseCase
	Delete  *use_cases.DeleteProductUseCase
}

func NewProductController(repo db.ProductRepositoryDB) *productController {
	return &productController{
		GetAll:  use_cases.NewGetAllProductsUseCase(repo),
		GetByID: use_cases.NewGetByIDProductUseCase(repo),
		Save:    use_cases.NewSaveProductUseCase(repo),
		Update:  use_cases.NewUpdateProductUseCase(repo),
		Delete:  use_cases.NewDeleteProductUseCase(repo),
	}
}

func (c *productController) GetAllHandler(ctx *gin.Context) {
	products, err := c.GetAll.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went wrong")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  products,
		"count": len(products),
	})
}

func (c *productController) SaveHandler(ctx *gin.Context) {
	var product domain.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := c.Save.Execute(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (c *productController) GetByIDHandler(ctx *gin.Context) {
	productId := ctx.Param("id")
	if productId == "" {
		response := domain.Response{
			Message: "invalid product_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
	}
	id, err := strconv.Atoi(productId)
	if err != nil {
		response := domain.Response{
			Message: "invalid product_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := c.GetByID.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *productController) UpdateHandler(ctx *gin.Context) {
	productId := ctx.Param("id")
	if productId == "" {
		response := domain.Response{
			Message: "invalid product_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(productId)
	if err != nil {
		response := domain.Response{
			Message: "invalid product_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product domain.Product
	product.ID = id

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid payload")
		return
	}

	if err := c.Update.Execute(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}
	ctx.JSON(http.StatusOK, "Product updated successfully")
}

func (c *productController) DeleteHandler(ctx *gin.Context) {
	productId := ctx.Param("id")
	if productId == "" {
		response := domain.Response{
			Message: "invalid product_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(productId)
	if err != nil {
		response := domain.Response{
			Message: "invalid product_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := c.Delete.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}
	ctx.JSON(http.StatusOK, "Product deleted successfully")
}
