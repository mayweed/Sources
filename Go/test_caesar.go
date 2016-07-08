package main

import "fmt"
//import "os"
import "strings"

/*
const str="abcdefghijklmnopqrstuvwxyz"
const STR="ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func index (s string,r rune) int{
    for k,v :=range s{
        if r==v{
            return k
        }
    }
    return 0
}
*/
//TODO:write a encode func and a decode func
//use containsRune to make the diff between small and capital
//treat special char like spaces hyphen etc...
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
	//cf http://golangcookbook.com/chapters/strings/processing/ excellent!!
    //removePunctuation := func(r rune) rune {
	//	if strings.ContainsRune(".,:;-", r) {
	//		return -1
	//	} else {
	//		return r
	//	}
	//}
	//str_to_crypt = strings.Map(removePunctuation, s)
	//words := strings.Fields(str_to_crypt)
    //var tokens_trans []string
    //QUESTION:how to preserve my tokens once translation done??
    //HERE should loop over tokens

	decode_caesar := func(r rune) rune {
        var dec rune
        //just add !!! and then sub r, no need for strings index!!
		switch {
		case r >= 'A' && r <= 'Z':
			dec= 'A' +(((r-'A')-offset)%26)
		case r >= 'a' && r <= 'z':
			dec='a'+(((r-'a')-offset)%26)
            if dec < 0 {dec+=26}
        }
        fmt.Println(dec)
		//if dec < 0{
        //    dec+=26
        //    return dec
        //}//else{
		 //   return dec
		//}
        return dec
	}
	fmt.Println(strings.Map(decode_caesar,str_to_crypt))
/*
    for _,word := range words{
        //and then loop over each string in tokens
        for _,c :=range word{
            //enc:=(index(str,c)+offset)%26
            //translated=append(translated,string(str[enc]))
            dec:=(index(str,c)-offset)%26
            //dirty trick
            //see:https://groups.google.com/forum/#!topic/golang-nuts/xj7CV857vAg
                        translated=append(translated,string(str[dec]))
        }
        tokens_trans=append(tokens_trans,strings.Join(translated,""))
    }
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
    //f:=strings.ContainsRune(str,'A')
    //fmt.Println(tokens_trans)
    //fmt.Println(strings.Join(translated,""))
    //fmt.Fprintln(os.Stderr,"DEBUG:...")
*/
}
