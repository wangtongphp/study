<?PHP

$b_ti = microtime_float();

//$data = file_get_contents("php://stdin");
//$data = file_get_contents("./3/s/input-3.txt");
//$datas = explode(PHP_EOL,$data);

f1();

function f1(){
    //$p = array();
    //$arr = array();
    for($i=0;$i<999999;$i++){
        $s = "a bb c";
        $p = explode(' ',$s);
        $arr[$p[0]]= $i;    
        $arr[$p[1]] = $i;    
        $arr[$p[2]] = $i;    
    }
}


function f2(){
    $p0 = 0;
    $p1 = 0;
    $p2 = 0;
    $arr = array();
    for($i=0;$i<999999;$i++){
        $s = "a bb c";
        list($p0,$p1,$p2) = explode(' ',$s);
        $arr[$p0] = $i;
        $arr[$p1] = $i;
        $arr[$p2] = $i;
    }
}

$e_ti =microtime_float();
echo $e_ti-$b_ti;

function microtime_float()
{
        list($usec, $sec) = explode(" ", microtime());
        return ((float)$usec + (float)$sec);
}
