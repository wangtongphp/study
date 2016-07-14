<?php
/**
 * @author wangtong1@xiaomi.com
 */

$data = file_get_contents("php://stdin");
$datas = explode(PHP_EOL,$data);
array_pop($datas);
$group = array_shift($datas);
for($g=0; $g<$group; $g++){

    $line = array_shift($datas);
    for($i=0;$i<$line;$i++){
        $p = array_shift($datas);

    }
}