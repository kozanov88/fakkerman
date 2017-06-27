package main

import "fmt"
//import "strconv"

var temp map[string]int

func find(m,n int)int {
  key:=fmt.Sprintf("%d",m) + "," + fmt.Sprintf("%d",n)
  res:=temp[key]
  if res==0 {

    res= int(1)
  }
  return res
}

func akkerman(m,n int)  int{
 if m<0 || n<0 {
    return 0
  }

  res := int(0);
  switch {
    case m==0:{
      res=n+1
      key:=fmt.Sprintf("%d",m) + "," + fmt.Sprintf("%d",n)
      temp[key] = res
      fmt.Println(temp[key])
    }

      
    case m>0 && n==0:{

      res=akkerman(m-1,int(1))
      }
      
    case m>0 && n>0:{
      res=akkerman(m-1,akkerman(m,n-1))}
  }
  return res
}

func main() {

  m := int(3)
  n := int(3)

  fmt.Println("m=3 n=3",akkerman(m, n))
  
  m = 4
  n = 1  

  fmt.Println("m=4 n=1",akkerman(m, n))
}