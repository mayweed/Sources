package main

import "fmt"
//import "os"
import "strings"

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
    offset:=7
    //str_to_crypt:="ave caesar"
    str_to_crypt:="zpe mvby zpe mvby aoyll"

    //split here first!!
    var translated []string
    tokens:=strings.Split(str_to_crypt," ")
    fmt.Println(tokens)
    //QUESTION:how to preserve my tokens once translation done??
    //HERE should loop over tokens
    for i:=0;i<len(tokens);i++{
        //and then loop over each string in tokens
        for _,c :=range tokens[i]{
            if c !=' '{
                //enc:=(index(str,c)+offset)%26
                //translated=append(translated,string(str[enc]))
                dec:=(index(str,c)-offset)%26
                //dirty trick
                //see:https://groups.google.com/forum/#!topic/golang-nuts/xj7CV857vAg
                if dec < 0{
                    dec+=26
                }
                translated=append(translated,string(str[dec]))
            }
        }
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
    */
    //f:=strings.ContainsRune(str,'A')
    //fmt.Println(translated)
    fmt.Println(strings.Join(translated,""))
    //fmt.Fprintln(os.Stderr,"DEBUG:...")
}
