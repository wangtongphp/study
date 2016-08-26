<?php
/**
 * @desc 
 * @author wangtong1@xiaomi.com
 */

$file = file_get_contents("./input-3-t.txt");
//$file = file_get_contents("php://stdin");
$f = explode(PHP_EOL,$file);
$k = 0;
$group = $f[$k++];
for($g=0; $g<$group; $g++){
    list($num, $line, $ask) = explode(' ',$f[$k++]);
    for($l=0;$l<$line;$l++){
        $p = explode(' ',$f[$k++]);
        $res[$p[0]] = '-'.$p[1];
        if($res[$p[0]]<0){
            echo  'liar';
            echo PHP_EOL;
            break;
        }
        $res[$p[0]+$p[2]] = $p[1];
        $res[$p[0]
    }
}
