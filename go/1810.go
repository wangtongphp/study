package main

func main(){
    data:=[6]int{1,5,2,7,2,4}
    exec1(data, 6, 2)
}
func exec1(data [6]int , len int, elem int ){
    for l,r,i:=0,(len-1),0; i<=r;{
        if data[i] < elem{
            data[i++],data[l++] = data[i++],data[l++]
        }
        else if data[i] > elem{
            data[i],data[r--] = data[r--],data[i]
        }
        else{
            ++i;
        }
    }
}
