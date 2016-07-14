<?PHP


$b_ti = microtime();
for($i=0;$i<99999;$i++){
    $arr[] = $i;
}


$j = 0;
for($i=0;$i<99999;$i++){
    $arr[$j++] ;
    //$arr[++$j] ;
    //$arr[$i];
    //array_shift($arr);
}
    

$e_ti = microtime();
echo $e_ti-$b_ti;
