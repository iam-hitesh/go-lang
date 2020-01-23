// http://www.golangbootcamp.com/book/_single-page#sec-variables

package main

import "fmt"

func main(){
  var name = "Hitesh";
  fmt.Printf(name);
  
  age := 21;
  // fmt.Printf(age) would not work
  fmt.Printf("%d", age);
  
  var(
    details string = "None";
  );
  fmt.Printf(details);
}
  
