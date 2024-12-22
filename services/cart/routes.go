package cart

import (
	mytypes "ecom_test/my_types"
	"ecom_test/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)



type Handler struct{

	store mytypes.OrderStore;
	productStore mytypes.ProductStore;
}

func GetCartHandler(store mytypes.OrderStore,productStore mytypes.ProductStore)*Handler{
	return &Handler{store: store,productStore: productStore};
}

func (h*Handler)RegisterRoutes(router *mux.Router)error{
	router.HandleFunc("/cart/checkout",h.handleCheckOut).Methods(http.MethodPost);
	return nil;
}

func (h*Handler)handleCheckOut(w http.ResponseWriter,req*http.Request){
	var cartpayload mytypes.CartCheckoutPayload;
	err:=utils.ParseJson(req,&cartpayload);
	if err!=nil{
		utils.WriteJsonError(w,http.StatusBadRequest,err);
		return ;
	}
	if err=utils.Validator.Struct(cartpayload);err!=nil{
		errors:=err.(validator.ValidationErrors);
		utils.WriteJsonError(w,http.StatusBadRequest,fmt.Errorf("error at formating: %v",errors));
		return ; 
	}
	
	all_id,err:=GetItemsId(cartpayload.Items);
	if err!=nil{
		utils.WriteJsonError(w,http.StatusBadRequest,err);
	}

	//get the products id 
	products,err:=h.productStore.GetProductsByIds(all_id);
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)
		return
	}
	
	createOrder(cartpayload,products);
	

}