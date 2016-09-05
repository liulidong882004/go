package main
import ("fmt")
var temp int
func use_array(args *[4]int){
     for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] < args[j] {
				temp = args[i]
				args[i] = args[j]
				args[j] = temp
	                }       
                }
    } 
}
func main() {
    var args = [4]int{1, 2, 3, 4};
    use_array(&args);
    fmt.Println(args);
}
