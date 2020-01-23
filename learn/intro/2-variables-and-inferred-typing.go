// http://www.golangbootcamp.com/book/_single-page#sec-variables

package main

import "fmt"

func main(){
  var name = "Hitesh";
  // Println for line break after print
  fmt.Println(name);
  
  age := 21;
  // fmt.Printf(age) would not work
  fmt.Printf("%d", age);
  
  var(
    details string = "None";
  );
  fmt.Printf(details);

  action := func() {
    fmt.Printf("In Function")
  }
  action();
}

// Output
// Hitesh
// 21NoneIn Function
