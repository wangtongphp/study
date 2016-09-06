<?php
/**
 * @desc 
 * @author wangtong1@xiaomi.com
 */
$file = file_get_contents("./input-2-t.txt"); 
//$file = file_get_contents("php://stdin");
$f = explode(PHP_EOL,$file);
$k = 0;
$group = $f[$k++];
for($g=0; $g<$group; $g++){
    
    $yuan = $f[$k++];
    
    $t_cost = explode(' ', trim($f[$k++]));
    for($t = 1; $t<=9;$t++){
        $cost[$t] = $t_cost[$t-1];
    }
    
    $cost_orig = $cost;
    asort($cost);
    $fi = 0;
    $res = '';
    foreach($cost as $ck=>$cv){
        $fi++;
        if($yu - $cv < 0 && $yu2-$cv <0 && isset($preck)){
            break;
        }
        if($fi == 1){
            //one
            $len = intval($yuan/$cv);
            $len2 = intval($yuan/$cv)-1;
            if($len == 0){
                echo -1;
                echo PHP_EOL;
                continue 2;
            }
            $yu = $yuan - $len*$cv;
            $yu2 = $yuan - $len2*$cv;
            $res = str_repeat($ck,$len);
            $res2 = str_repeat($ck,$len2);
            //$res[$ck]=$cv;
            if($yu == 0){
                echo $res;
                echo PHP_EOL;
                continue 2;
            }
        }else{

            if($yu>=$cv){
                $res = $ck.$res;
            }
            if($yu2>=$cv){
                $res2 = $ck.$res2;
            }
        }
        $preck = $ck;
        
    }

    echo $res2> $res1 ?$res2 : $res1;

}

function f1(){

}
