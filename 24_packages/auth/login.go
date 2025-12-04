package auth

//We have to keep the function or struct name "capital" if we want to export it...if we don't want to export the function then keep it small
import (
	"fmt"
	product "myApp/products"
	"github.com/fatih/color"
)

func LoginUser() {
	fmt.Println("User logged in")
	product.CreateProduct()
	color.Blue("Hey")
}
