//utils

func In(c int,d []int)bool{
    for _,v := range d{
        if v==c{
            return true
        }
    }
    return false
}
