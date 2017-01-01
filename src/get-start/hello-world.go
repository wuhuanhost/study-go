package main
import "fmt"
/*json解析的包*/
import "github.com/json-iterator/go"

func main(){
  var str="我的第一个go语言代码";
  fmt.Printf("hello world!!!\n");
  fmt.Printf(str+"\n");
  iter := jsoniter.ParseString(`[0,1,2,3]`)
  val := []int{}
  iter.Read(&val)
  fmt.Println(val[3])
}
