package main
import ("fmt")
var temp int
func use_array(args [] int,num int){
     for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if args[i] < args[j] {
				temp = args[i]
				args[i] = args[j]
				args[j] = temp
	                }       
                }
    } 
}
func main() {
    var args = []int{1, 2, 3, 4,6,8,9};
    use_array(args,7);
    fmt.Println(args);
}
