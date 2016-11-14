<?PHP

$arr = array(
    1<<0,
    1<<1,
    1<<2,
    1<<3,
    1<<4,
    1<<5
    );
var_dump($arr);
$sum = 9;

//找到所有的组成部分
$sum_s = strrev((string)decbin($sum));
for($i = 0; ;$i++){
    if(!isset($sum_s[$i])){
        break;
    }
    if($sum_s[$i] == 1){
        $r[] = 1<<$i;
    }
}

var_dump(bitToArr(2097144));

function bitToArr($sum){

    $sum_s = strrev((string)decbin($sum));
    $sum_s_len = strlen($sum_s);
    $r = array();
    for($i = 0; $i<$sum_s_len;$i++){
        if($sum_s{$i} == 1){  
            $r[] = $i;
        }    
    }    
    return $r;

} 
