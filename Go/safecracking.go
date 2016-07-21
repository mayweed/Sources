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
        "zero":0,
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
    var offset rune
    var translated string
    //TrimSpace is important, w/o it wont work!!
    var str_to_decrypt="xxx: wkuhh-ilyh-rqh-vla-wzr"
    var str_to_crypt=strings.TrimSpace(strings.Split(str_to_decrypt,":")[1])

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

    for offset<26{
	    translated=strings.Map(decode_caesar,str_to_crypt)
        for _,word :=range strings.Split(translated,"-") {
            for key,_:= range numbers{
                if key==word{
                    fmt.Printf("%d",numbers[word])
                }
            }
        }
        offset+=1
    }

    //just cleaner no?
    fmt.Printf("\n")
    //fmt.Fprintln(os.Stderr,"DEBUG:...")
}
