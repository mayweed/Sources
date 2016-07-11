package main

import "fmt"
import "strings"
import "unicode"

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
        if unicode.IsLetter(r){
		    switch {
		    case r >= 'A' && r <= 'Z':
                delta:=(r-'A')-offset
                if delta < 0 {delta +=26}
			    dec= 'A' +(delta%26)
		    case r >= 'a' && r <= 'z':
                //in golang -3%26=-3 and NOT 23...
                delta:=(r-'a')-offset
                if delta <0 {delta += 26}
			    dec='a'+(delta%26)
            }
        } else{
            dec=r
        }
        return dec
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
