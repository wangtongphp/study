<?PHP

$b_ti = microtime_float();

//$data = file_get_contents("php://stdin");
//$data = file_get_contents("./3/s/input-3.txt");
//$datas = explode(PHP_EOL,$data);

for($i=0;$i<999999;$i++){
    $s = "a bb c d";
    $p = explode(' ',$s);
    $arr[$p[0]]= $i;    
    $arr[$p[1]] = $i;    
    $arr[$p[2]] = $i;    
    //list($p0,$p1,$p2,$p3) = explode(' ',$s);
    //$arr[$p0] = $i;
    //$arr[$p1] = $i;
    //$arr[$p2] = $i;
    //$r_{$i} = 0;
    //$k = rand(0,99999);
}



$e_ti =microtime_float();
echo $e_ti-$b_ti;

function microtime_float()
{
        list($usec, $sec) = explode(" ", microtime());
        return ((float)$usec + (float)$sec);
}
