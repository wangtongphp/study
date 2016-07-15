<?php
/**
 * @author wangtong1@xiaomi.com
 * @desc 第四届编程比赛第1题
 */

$data = file_get_contents("php://stdin");
$datas = explode(PHP_EOL,$data);
$k = 0;
$group = $datas[$k++];
for($g=0; $g<$group; $g++){
    $line = $datas[$k++];
    $mihomes = explode(" ", $datas[$k++]);
    $fens = explode(" ",$datas[$k++]);
    run2($mihomes, $fens);
}


function run1($mihomes,$fens){
    sort($mihomes);
    sort($fens);
    //var_dump($fens);
    //var_dump($mihomes);
    $cnt = count($mihomes);
    $total =0; 
    foreach($fens as $k=>$v){
        foreach($mihomes as $kk=>$vv){
            $len = abs($v - $vv);
            if(!isset($min)){
                $min = $len;
            }
            if($len > $min){
                $total += $min;
                //echo $min.PHP_EOL;
                //var_dump($v, $vv);
                unset($min);
                break;
            }else{
                $min = $len;
            }

            if($kk==$cnt-1){
                $total += $len;
                unset($min);
                break;
            }

        }
    }
    echo $total;
    
    echo PHP_EOL;
}

function run2($mihomes,$fens){
    sort($mihomes);
    sort($fens);
    //var_dump($mihomes);
    $cnt = count($mihomes);
    $total =0; 
    $j = 0;
    foreach($fens as $k=>$v){
        for($i=$j;$i<$cnt;$i++){
            $len = abs($v-$mihomes[$i]);
            if(!isset($min)){
                $min = $len;
                continue;
            }
            if($len > $min){
                $j = $i-1;
                $total += $min;
                unset($min);
                break; 
            }else{
                $min = $len;
            }
            if($i==$cnt-1){
                $j = $i-1;
                $total += $len;
                unset($min);
                break;
            }
        }
    }
    echo $total;
    
    echo PHP_EOL;
}


/**
题目描述：
小米今年开设了许多家零售店，米粉们可以很轻松的找到离自己最近的零售
店购买手机，现在给出 n 个米家，以及 m 个米粉的位置，求出所有米粉走到距
离自己最近米家的总距离。
为了简化问题，米家和米粉都在一个一维空间上，坐标范围在[0, 10!]之间。
输入：
输入第一行为 T，表示有 T 组测试数据， （1<=T<=10）。
对每组数据，第一行是两个正整数 n，m（0 < n <= 100000,  0 < m <=
100000），分别表示有 n 个米家，以及 m 个米粉。
接下来一行，有 n 个整数， �!�! … �!!!，分别表示每个米家的坐标(�!< 
0 ≤ �! ≤ 10!)，中间用一个空格分隔。
接下来一行，有 m 个整数，�!�! … �!!!，分别表示每个米粉的坐标
(0 ≤ �! ≤ 10!)，中间用一个空格分隔。
输出：
对于每组数组，输出一个整数，求出所有米粉找到离自己最近米家的总距离。
样例输入：
1
3 4
1 5 6
8 2 3 5

样例输出：
5
*/
