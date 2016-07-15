<?php
/**
 * @author wangtong1@xiaomi.com
 */

$data = file_get_contents("php://stdin");
$datas = explode(PHP_EOL,$data);
$k = 0;
$group = $datas[$k++];
for($g=0; $g<$group; $g++){

    $line = $datas[$k++];
    for($i=0;$i<$line;$i++){
        $p = $datas[$k++];

    }
}

function run($data){

    echo PHP_EOL;
}
