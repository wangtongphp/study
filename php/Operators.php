<?PHP

$values = array(0,1,2,4,8);

$r = 5;

foreach($values as $v){
    echo $r & $v;
    echo PHP_EOL;
}
