
type SubType struct{
    sub1 int
    sub2 float
}

type TestType struct{
    member1 [3]int
    member2 float
    member3 [3]SubType
}

var user [2]TestType

func main(){
    result := user[1].member3[1].sub1
    __dump("構造体テスト途中ログ.log")
}
