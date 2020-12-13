package main

import (
	fiken "./fikenrq"
	shopify "./shopifyrq"
)

func main() {
	//testing bool, when true: test fiken package
	//When false: test shopify package
	var isFiken bool = false

	if isFiken {
		fiken.Test()
	} else {
		// shopify.TestOrders()
		shopify.TestOrdersWithDocker()
	}

}
