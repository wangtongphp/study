<?PHP

$a = $argv[1];
$a = array(3,4,5,6,8);
$b = $argv[2];

$c = count($a);
$l = 0;
$h = $c-1;

while($l < $h){
    $m = intval(($l+$h)/2);

    if($b == $a[$m]){
        echo $m;
    }
    if($b > $a[$m]){
        $l = $m + 1;
    }else{
        $h = $m -1;
    }
}

