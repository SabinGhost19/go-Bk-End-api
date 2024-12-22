package cart

import (
	mytypes "ecom_test/my_types"
	"net/http"

	"github.com/gorilla/mux"
)



type Handler struct{
	//store cart store
	store mytypes.OrderStore;
}

func GetCartHandler(store mytypes.OrderStore)*Handler{
	return &Handler{store: store};
}

func (h*Handler)RegisterRoutes(router *mux.Router)error{
	router.HandleFunc("/cart/checkout",h.handleCheckOut).Methods(http.MethodPost);
	return nil;
}
func (h*Handler)handleCheckOut(w http.ResponseWriter,req*http.Request){
	
}