package main

//import ("fmt"
//        "os"
//        "bufio")
import "fmt"

func calcArbitrage(parity []float32) float32{
    var start float32=100000.0
    var arbitrage float32=start/parity[0]
    for k:=1;k<len(parity);k++{
        arbitrage/=parity[k]
    }
    if arbitrage-start > 0{
        fmt.Println(arbitrage-start)
    }else{
        fmt.Println(0)
    }
    return 0
}

func main() {
    //scanner:=bufio.NewScanner(os.Stdin)

    var numTestCases int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&numTestCases)

    //Should I write a custom splitFunc for my floats??
    //var num float32
    //var parity []float32
    for i:=0;i<numTestCases;i++{
        for i:=0;i<3;i++{
             _,err:=fmt.Fscan(in,&num)
            if err != nil{break}
            parity=append(parity,num)
        }

    //parity:= []float32{1.1837,1.3829,0.6102}
    parity:= []float32{1.1234,1.2134,1.2311}

    //for{
    //    for i:=0;i<3;i++{
    //        _,err:=fmt.Fscan(in,&num)
    //        if err != nil{break}
    //        parity=append(parity,num)
    //    }
        calcArbitrage(parity)
    }
}
