package cart

import (
	mytypes "ecom_test/my_types"
	"fmt"
)

func GetItemsId(items []mytypes.CartCheckoutItem)([]int,error){
	all_id:=make([]int,len(items));
	
	for index,item:=range items{
		if item.Quantity<=0{
			return nil,fmt.Errorf("invalid quantitiy of the product : %d",item.Quantity);
		}
		all_id[index]=item.ItemID;
	}
	return all_id,nil;
}

func createOrder(art mytypes.CartCheckoutPayload,products[]mytypes.Product){

}
func checkIfTheCartIsInStock(cart mytypes.CartCheckoutPayload){

}