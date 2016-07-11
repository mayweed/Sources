package main

import "fmt"
//import "os"
import "strings"

//TODO:write a encode func and a decode func
//use containsRune to make the diff between small and capital
//treat special char like spaces hyphen etc...
//dirty trick with modulo:
//see:https://groups.google.com/forum/#!topic/golang-nuts/xj7CV857vAg
//Strings processing:
//cf http://golangcookbook.com/chapters/strings/processing/ excellent!!
func main(){
    /*
    numbers:=map[string]int{
        "one":1,
        "two":2,
        "three":3,
        "four":4,
        "five":5,
        "six":6,
        "seven":7,
        "eight":8,
        "nine":9,
    }
    */
    //str_to_crypt:="ave caesar"
    var offset rune=7
    str_to_crypt:="zpe-mvby-zpe-mvby-aoyll"
	//removePunctuation := func(r rune) rune {
	//	if strings.ContainsRune(".,:;-", r) {
	//		return -1
	//	} else {
	//		return r
	//	}
	//}
	//str_to_crypt = strings.Map(removePunctuation, s)
	//words := strings.Fields(str_to_crypt)
    //QUESTION:how to preserve my tokens once translation done??
	decode_caesar := func(r rune) rune {
        var dec rune
        //just add !!! and then sub r, no need for strings index!!
		switch {
		case r >= 'A' && r <= 'Z':
			dec= 'A' +(((r-'A')-offset)%26)
		case r >= 'a' && r <= 'z':
			dec='a'+(((r-'a')-offset)%26)
            //if dec < 0 {dec+=26}
        }
        //that trick does not work here!!
        if dec > 0{
            return dec
        }else{
            return dec+26
        }
	}
	fmt.Println(strings.Map(decode_caesar,str_to_crypt))
/*
    var combi []int
    for _,v:=range translated{
        for key,_ := range numbers{
            fmt.Println(key,v)
            if v==key{
                combi=append(combi,numbers[key])
                break
            }
        }
    }
    //fmt.Fprintln(os.Stderr,"DEBUG:...")
*/
}
