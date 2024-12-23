package cart

import (
	mytypes "ecom_test/my_types"
	"ecom_test/services/auth"
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

	//get the user item for use in the order creation 
	userID:=auth.GetUserIdfromContext(req.Context());

	//parse and populate the struct payload 
	err:=utils.ParseJson(req,&cartpayload);
	if err!=nil{
		utils.WriteJsonError(w,http.StatusBadRequest,err);
		return ;
	}
	//validate the struct based on the mytypes declaration
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
	
	//create the oreder
	orderID,toatal,err:=h.createOrder(cartpayload.Items,products,userID);
	if err!=nil{
		utils.WriteJsonError(w, http.StatusBadRequest, err)
		return
	}

	//return the response 
	utils.WriteJson(w,http.StatusOK,map[string]interface{}{
		"total price":toatal,
		"order_id":orderID,
	});

}