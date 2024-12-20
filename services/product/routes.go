package product

import (
	mytypes "ecom_test/my_types"
	"ecom_test/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)


type Handler struct{
	store mytypes.ProductStore
}
func GetProductHandler(store mytypes.ProductStore)*Handler{
	return &Handler{store:store};
}

func (h*Handler)RegisterRoutes(router*mux.Router)error{
	//router.HandlerFunc("/products",h.handlerCreateProduct).Methods(http.MethodPost);
	router.HandleFunc("/addproduct",h.handlerCreateProduct).Methods(http.MethodPost);	
	router.HandleFunc("/products",h.handlerGetProduct).Methods(http.MethodGet);	
	return nil;
}
func (h*Handler)handlerCreateProduct(w http.ResponseWriter,req*http.Request){
	var payload mytypes.ProductPayload;
	err:=utils.ParseJson(req,&payload);
	if err!=nil{
		utils.WriteJsonError(w,http.StatusBadRequest,err);
		return ;
	}
	if err=utils.Validator.Struct(payload);err!=nil{
		errors:=err.(validator.ValidationErrors);
		utils.WriteJsonError(w,http.StatusBadRequest,fmt.Errorf("error at formating: %v",errors));
		return ; 
	}
	_,err=h.store.GetProductByName(payload.Name);
	if err==nil{
		utils.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("product with name: %v already exists", payload.Name))
	}
	err=h.store.CreateProduct(mytypes.Product{Name: payload.Name,
		Description: payload.Description,
		Image: payload.Image,
		Price: payload.Price,
		Quantity: payload.Quantity,
		});

	if err!=nil{
		utils.WriteJsonError(w,http.StatusInternalServerError,err);
		return ;
	}
	
	utils.WriteJson(w,http.StatusCreated,nil);
}
func (h*Handler)handlerGetProduct(w http.ResponseWriter,req*http.Request){
	products,err:=h.store.GetProducts();
	if err!=nil{
		utils.WriteJsonError(w,http.StatusInternalServerError,err);
		return ;
	}
	utils.WriteJson(w,http.StatusOK,products);
}