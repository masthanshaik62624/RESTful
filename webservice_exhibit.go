package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

//Nested address struct
type Address struct {
	StreetInfo string `json:"StreetInfo"`
	City string `json:"City"`
	State string `json:"State"`
	ZipCode int `json:"ZipCode"`
}

//outer struct
type CustomerInfo struct {
	Name string `json:"Name"`
	Phone int `json:"Phone"`
	Location Address `json:"Address"`
}

func webAPI (wg *sync.WaitGroup) {
	defer wg.Done()

	r := gin.Default()

	var newStruct CustomerInfo

	//the default listening port is 8080
	r.POST("customer-info-service", func(c *gin.Context) {

		//var newStruct CustomerInfo

		if err := c.ShouldBindJSON(&newStruct); err != nil {
			c.AbortWithStatus(500)
			return
		}

		fmt.Println("Customer Name and City: ")
		fmt.Println(newStruct.Name)
		fmt.Println(newStruct.Location.City)

	})

	r.GET("/customer", func (c *gin.Context) {
		c.JSON(200, newStruct)
	})

	r.Run()

}

func main () {
	var wg sync.WaitGroup

	wg.Add(1)

	go webAPI(&wg)
	wg.Wait()
}