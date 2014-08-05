package main

import "fmt"
import "github.com/stcatz/api_parser"

func main() {
	s := `HOST: http://www.google.com/

--- Sample API v2 ---
---
Welcome to the our sample API documentation. All comments can be written in (support [Markdown](http://daringfireball.net/projects/markdown/syntax) syntax)
---

--
Shopping Cart Resources
The following is a section of resources related to the shopping cart
--
List products added into your shopping-cart. (comment block again in Markdown)
GET /shopping-cart
< 200
< Content-Type: application/json
{ "items": [
{ "url": "/shopping-cart/1", "product":"2ZY48XPZ", "quantity": 1, "name": "New socks", "price": 1.25 }
] }


Save new products in your shopping cart
POST /shopping-cart
> Content-Type: application/json
{ "product":"1AB23ORM", "quantity": 2 }
< 201
< Content-Type: application/json
{ "status": "created", "url": "/shopping-cart/2" }


-- Payment Resources --
This resource allows you to submit payment information to process your *shopping cart* items
POST /payment
{ "cc": "12345678900", "cvc": "123", "expiry": "0112" }
< 200
{ "receipt": "/payment/receipt/1" }`
	a, _ := api_parser.Parse(s)

	fmt.Printf(string(a))
}