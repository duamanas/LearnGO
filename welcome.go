package main
import "fmt"
import "math"
//import "strings"

func add(a int, b int) int {
	return a + b
}

func addtoslice( impslicenames []int, index int) {
 impslicenames[index] = impslicenames[index] * 2
}
func addpointer (impvar *int) {
	*impvar = *impvar * 2
}

func sqrt( impa float64) (float64,error) {
   if impa < 0.0{
   return 0.0, fmt.Errorf("Sqrt of value (%f)" , impa) } else{
   return math.Sqrt(impa), nil }
}
func main() {
  c := add(2,3)
  fmt.Println(c)
  
slicenames := make([]int,3)
slicenames[0] = 10
slicenames[1] = 20
slicenames[2] = 30
addtoslice(slicenames,2)
fmt.Println(slicenames)

pointerint := 10
addpointer(&pointerint)
fmt.Println(pointerint)

sqrtreturn, errreturn := sqrt(10.0)

if errreturn == nil {
	fmt.Println(sqrtreturn)
}else{fmt.Println(errreturn)}
/*	fmt.Println("Hello World")

fmt.Println("Hello Again")
x,y := 2, 3
var z int = (x+y)/2
println(z)
if x > 2 {
	fmt.Println("x is greater than 2")
} else { 
	fmt.Println("x is less than or equal to 2") 
}
switch {
case x > 2: 
fmt.Println("case x>2")
case x == 2:
fmt.Println("case x=2")	
}
var count int
count = 10
for i := 0; i <= count; i++ {
	if i < 1{continue}
	fmt.Println(i)
}*/

/*var count int
count = 20
var frizz string
frizz = "frizz"
var buzz string
buzz = "buzz"
var frizzbuzz string
frizzbuzz = frizz + " " + buzz
for i := 1; i <= count; i++ {
if i%3 == 0 && i%5 == 0 { fmt.Println(frizzbuzz) 
} else if i%3 == 0 { fmt.Println(frizz) 
} else if i%5 == 0 { fmt.Println(buzz) 
} else if i%3 != 0 || i%5 != 0 { fmt.Println(i) } 
}
*/
/*var frizz string
frizz = "frizz"
var buzz string
buzz = "buzz"
var frizzbuzz string
frizzbuzz = frizz + " " + buzz
fmt.Println(len(frizzbuzz))
fmt.Println(frizzbuzz[:5])
fmt.Println(frizzbuzz[1:5])
fmt.Println(frizzbuzz[0:2])
count :=100
strcount := fmt.Sprintf("%d",count)
println(strcount[:1])
for i := 10; i < count; i++ {
numstring := fmt.Sprintf("%d",i)
length := len(numstring)
beg := numstring[:1]
end := numstring[length-1:length]
	if beg == end {fmt.Println(i)}
}
slicenames := make([]string, 3)
slicenames[0] = "Manas"
slicenames[1] = "Dua"
slicenames = append(slicenames,"Test")
fmt.Println(slicenames[0])
fmt.Println(slicenames[1])
fmt.Println(slicenames[2])
fmt.Println(slicenames[3])*/

/*slicenum := make([]int, 6)
slicenum[0] = 100
slicenum[1] = 5
slicenum[2] = 4
slicenum[3] = 67
slicenum[4] = 22
slicenum[5] = 88

slicelen := len(slicenum)
var greatest int
greatest = 0
for i := 0; i < slicelen; i++ {
	if slicenum[i] > greatest{
		greatest = slicenum[i]
	} 
}
fmt.Println(greatest) */
/*mapstocks := make( map[string]float32)
mapstocks["TSLA"] = 123.1
mapstocks["GOOGL"] = 999.1
fmt.Println(mapstocks["TSLA"])
value,exists := mapstocks["TSLA"]
fmt.Println(exists)
fmt.Println(value)
fmt.Println(mapstocks["GOOGL"])
delete(mapstocks,"GOOGL")
fmt.Println(mapstocks["GOOGL"])
value,exists = mapstocks["GOOGL"]
fmt.Println(exists)
fmt.Println(value) */

/*mapwords := make(map[string]int)
text := "This is a needle This ss is a Needle needles and pins"
strings.ToLower(text)
words := strings.Fields(text)
lenght := len(words)
var count int
for i := 0; i < length; i++ {
	if flag, wordvalue:= mapwords[words[i]] flag = false{
	mapwords[words[i]] = count }
}
fmt.Println(mapwords)*/



}