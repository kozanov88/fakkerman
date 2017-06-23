package main

import "fmt"

func akkerman(m,n uint64)  uint64{
 if m<0 || n<0 {
    return 0
  }

  res := uint64(0);
  switch {
    case m==0:
      res=n+1
    case m>0 && n==0:
      res=akkerman(m-1,uint64(1))
    case m>0 && n>0:
      res=akkerman(m-1,akkerman(m,n-1))
  }
  return res
}

func main() {

  m := uint64(3)
  n := uint64(3)

  fmt.Println(akkerman(m, n))
}