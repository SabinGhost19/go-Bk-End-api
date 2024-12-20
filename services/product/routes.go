package product

import (
	mytypes "ecom_test/my_types"
	"ecom_test/utils"
	"net/http"

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
	router.HandleFunc("/products",h.handlerCreateProduct).Methods(http.MethodGet);	
	return nil;
}

func (h*Handler)handlerCreateProduct(w http.ResponseWriter,req*http.Request){
	products,err:=h.store.GetProducts();
	if err!=nil{
		utils.WriteJsonError(w,http.StatusInternalServerError,err);
		return ;
	}
	utils.WriteJson(w,http.StatusOK,products);
}