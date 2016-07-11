//dirty trick with modulo:
//see:https://groups.google.com/forum/#!topic/golang-nuts/xj7CV857vAg
//Strings processing:
//cf http://golangcookbook.com/chapters/strings/processing/ excellent!!

package main

import (
    "fmt"
    "strings"
    "unicode"
    )

func main(){
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

    //str_to_crypt:="ave caesar"
    var offset rune=1
    str_to_crypt:="zpe-mvby-zpe-mvby-aoyll"

    //str_to_crypt:="Aol zhml jvtipuhapvu pz"
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
        } else {
            dec=r
        }
        return dec
	}
	translated:=strings.Map(decode_caesar,str_to_crypt)
    fmt.Println(translated)

    //Searching for the right offset should be separated...
    for _,word :=range strings.Split(translated,"-") {
            //here should use
            number,ok:=numbers[word]
            if !ok {
                offset+=1
            } else {
                fmt.Printf("%d",number)
            }
    }
    //just cleaner no?
    fmt.Printf("\n")
    //fmt.Fprintln(os.Stderr,"DEBUG:...")
}
