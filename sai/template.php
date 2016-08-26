<?php
/**
 * @desc 
 * @author wangtong1@xiaomi.com
 */

$file = file_get_contents("php://stdin");
$f = explode(PHP_EOL,$file);
$k = 0;
$group = $f[$k++];
for($g=0; $g<$group; $g++){

    $line = $f[$k++];
    for($l=0;$l<$line;$l++){
        $p = explode(' ',$f[$k++]);

    }
}
