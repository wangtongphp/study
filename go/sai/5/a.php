<?php
/**
 * @desc 
 * @author wangtong1@xiaomi.com
 */

//$file = file_get_contents("php://stdin");
$file = file_get_contents("./input-1-t.txt");
$f = explode(PHP_EOL,$file);
$k = 0;
$group = $f[$k++];
for($g=0; $g<$group; $g++){

    $d = trim($f[$k++]);
    $d_cnt = strlen($d);
    $data = array(); 
    for($i=0;$i<$d_cnt;$i++){
        $data[] = $d{$i};
    }
    $res = 0;
    f1($data,$res);
    echo $res;
    echo PHP_EOL;

}

function f1(&$data,&$res){

    $cnt = count($data);
    $i=0;
    if(empty($data) ){
        return;
    }
    foreach($data as $k=>$v){
        if(++$i == $cnt || $cnt < 3){
            return;
        }
        if($v==8 && $data[$k+1] == 1 && $data[$k+2] == 6){
            $res++;
            unset($data[$k]);
            unset($data[$k+1]);
            unset($data[$k+2]);
            break;
        }
    }
    $data = array_values($data);
    f1($data,$res);
}
