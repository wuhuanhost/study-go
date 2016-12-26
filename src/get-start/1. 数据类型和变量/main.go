package main  /**定义一个包**/
import "fmt"  /**导入fmt包**/
/**
**main函数是执行的入口
**/
func main(){
  /**　打印‘你好’
  **　go语言中对方的约定
  **　公开的方法名用帕斯卡命名法命名，
  **  私有方法用驼峰命名法命名。
  **/
  fmt.Println("你好");	
  Test();//方法调用
  TestNumber();
  setComplex();
}

/**
**变量
**/
func Test(){
/**
**基本的数据类型：布尔型，数字类型，字符串类型和派生类型
**派生类型包括：指针类型，数组类型，结构化类型，联合体类型，函数类型，切片类型，接口类型，map类型，channel类型 
**/
	var b bool = true;//定义一个bool型的变量b
	var i int = 100;//定义一个int型的变量i
	var s string = "你好中国";//定义一个字符串
	fmt.Println(b);
	fmt.Println(i);
	fmt.Println(s+"\n");

}



/**
**数字类型
**/

func TestNumber(){
  var i1 uint8=100;//定义一个无符号的8位整数(0_255)
  var i2 int8=-100;//定义一个有符号的8位整数(-128_127)
  var i3 uint16=10000;//无符号 16 位整型 (0 到 65535)
  var i4 int16=-10000;//有符号 16 位整型 (-32768 到 32767)
  var i5 uint32=1000000;//无符号 32 位整型 (0 到 4294967295)
  var i6 int32=-1000000;//有符号 32 位整型 (-2147483648 到 2147483647)
  var i7 uint64=10000000000;//无符号 64 位整型 (0 到 18446744073709551615)
  var i8 int64=-10000000000;//有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

   
  var f1 float32=121325.1234563;//IEEE-754 32位浮点型数
  var f2 float64=4548454874654845.231455;//IEEE-754 64位浮点型数


  var b1 byte=100;//类似 uint8
  var r1 rune=100000;//类似 int32
  var u1 uint=100;//32 或 64 位
  var i11 int=100;//与 uint 一样大小
  var u11 uintptr=100;//无符号整型，用于存放一个指针



  var c1 complex64=3.2+12i*12i;//32 位实数和虚数
  var c2 complex128=10000+12332i;//64 位实数和虚数

  fmt.Println(i1);
  fmt.Println(i2);
  fmt.Println(i3);
  fmt.Println(i4);
  fmt.Println(i5);
  fmt.Println(i6);
  fmt.Println(i7);
  fmt.Println(i8);

  fmt.Println(f1);
  fmt.Println(f2);

  fmt.Println(b1);
  fmt.Println(r1);
  fmt.Println(u1);
  fmt.Println(i11);
  fmt.Println(u11);

  fmt.Println(c1);
  fmt.Println(c2);

}

/**
 * 虚数使用
 */
 func setComplex() { 
   var value1 complex64 = 3.2 + 12i 
   value2 := 3.2 + 12i 
   value3 := complex(3.2, 12) 
   fmt.Println(value1, value2, value3) 
}