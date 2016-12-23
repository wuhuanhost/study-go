package main
import "fmt"
/*json*/
import "github.com/json-iterator/go"
func main(){
  iter := jsoniter.ParseString(`[0,1,2,3]`)
  val := []int{}
  iter.Read(&val)
  fmt.Println(val[3])
}
