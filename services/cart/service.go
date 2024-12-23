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

func (h*Handler)createOrder(cartItems []mytypes.CartCheckoutItem,products []mytypes.Product)(float64,error){
	product_map:=make(map[int]mytypes.Product);
	
	//create the map
	for _,product:=range products{
		product_map[int(product.ID)]=product;
	}

	err:=checkIfTheCartIsInStock(cartItems,product_map);
	if err!=nil{
		return 0,err;
	}

	total_price:=calculateTotalPrice(product_map,cartItems);

	//reduce the quantity of the products in the store
	for _,item:=range cartItems{
		product:=product_map[item.ItemID];
		product.Quantity=item.Quantity;
		h.productStore.UpdateProduct(product);
	}
	

}
func calculateTotalPrice(product_map map[int]mytypes.Product,items []mytypes.CartCheckoutItem)float64{
	var total float64;

	for _,item:=range items{
		product:=product_map[item.ItemID];
		total+=product.Price*float64(item.Quantity);
	}
	return total;
}
func checkIfTheCartIsInStock(cart []mytypes.CartCheckoutItem,product_map map[int]mytypes.Product)error{
	
	if len(cart)<=0{
		return fmt.Errorf("cart is empty")
	}

	for _,item:=range cart{
		product,ok:=product_map[item.ItemID];
		if !ok{
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity<item.Quantity{
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}
	return nil;
}