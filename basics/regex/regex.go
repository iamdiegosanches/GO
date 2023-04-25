package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	//reStr := "The ape was at the apex"
	//match, _ := regexp.MatchString("(ape[^ ]?", reStr) // ape not followed by space
	//pl(match)
	reStr2 := "Cat rat mat fat pat"
	match2, _ := regexp.Compile("([crmfp]at)")
	pl("Matchstring:", match2.MatchString(reStr2))
	pl("Findstring:", match2.FindString(reStr2))
	pl("Index:", match2.FindStringIndex(reStr2))
	pl("All String:", match2.FindAllString(reStr2, -1))
	pl("1st 2 String:", match2.FindAllString(reStr2, 2))
	pl("All submatch index:", match2.FindAllStringSubmatchIndex(reStr2, -1))
	pl(match2.ReplaceAllString(reStr2, "Dog"))
}
