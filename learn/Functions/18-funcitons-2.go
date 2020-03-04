// http://www.golangbootcamp.com/book/_single-page#sec-functions

package main

import (
    "fmt"
)

func location(city string) (string, string) {
  var region string
	var continent string

	switch city {
	  case "Jaipur", "Jodhpur", "Udaipur":
		  region, continent = "Rajasthan", "Asia"
	  case "Bengaluru", "Mangalore":
		  region, continent = "Karnataka", "Asia"
	  default:
		  region, continent = "Unknown", "Unknown"
	}
	
  return region, continent
}

func main(){
  fmt.Println(location("Jaipur"))
}

/**
// Output

Rajasthan Asia

**/
