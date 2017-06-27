package main

import "fmt"

func akkerman(m,n int64)  int64{
 if m<0 || n<0 {
    return 0
  }

  res := int64(0);
  switch {
    case m==0:
      res=n+1
    case m>0 && n==0:
      res=akkerman(m-1,int64(1))
    case m>0 && n>0:
      res=akkerman(m-1,akkerman(m,n-1))
  }
  return res
}

func main() {

  m := int64(3)
  n := int64(3)

  fmt.Println("m=3 n=3",akkerman(m, n))
  
  m = 4
  n = 1  

  fmt.Println("m=4 n=1",akkerman(m, n))
}