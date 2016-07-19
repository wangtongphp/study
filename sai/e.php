<?PHP

$b_ti = microtime();

for($i=0;$i<99999;$i++){
    //$r_{$i} = 0;
    //$k = rand(0,99999);
    $arr[] = $i;
}


$j = 0;
for($i=0;$i<99999;$i++){
    //echo next($arr); 
    //unset($arr[$i]);
    //$arr[$j++] ;
    //$arr[++$j] ;
    //$arr[$i];
    //array_shift($arr); //巨慢无比,等不到结果
}
    

foreach($arr as $k=>$v){

    for($i=0;$i<999;$i++){
        
    }
}


$e_ti = microtime();
echo $e_ti-$b_ti;
