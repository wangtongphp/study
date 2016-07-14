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
    $str = array_shift($datas);
    $arr = explode(" ", trim($str));
    run1($arr);
}

//解法一，循环每个数，键值属于0,3,8.. 时输出
function run1($arr){
    $line = count($arr);
    $i = 0;
    $j = 0;
    foreach ($arr as $k=>$v){
        if ($k+1 >= $line){
            $res[] = $v;
            break;
        }
    
        if($j == $k ){
            $res[] = $v;
            $j += $i*2+3;
            $i++;
        }
    }
    echo implode(" ", $res);
    echo PHP_EOL;
}

//解法二，分为n行，取每行最后一个数
function  run2($arr){
    $k = 0;
    $res = array();
    for ($i=1; ; $i++){
        $n = ($i-1)*2+1;
        $row = array_slice($arr,$k, $n);
        $k += $n;
        if (empty($row)){
            break;
        }
        $res[] = array_pop($row);
    }
    echo implode(" ", $res);
    echo PHP_EOL;
}


/**
 神奇数字
 题目描述：
 某天你突然得到一串秘密数字，M 告诉你解密的方法是，将这些数字按照第
 一行 1 个，第二行 3 个，第三行 5 个，第 n 行(n-1)*2+1 个的顺序排放这些数，
 然后将每行最后一个数组合起来就是最终的密码。
 你需要完成这个任务。
 示例：
 1 2 3 4 5 6 7 8 9 10
 摆放:
1
2 3 4
5 6 7 8 9
10
 则密码为：1 4   9   10
 输入：
 输入第一行为 T，表示有 T 组测试数据， （1<=T<=10）。
 接下来有 T 组测试数据，每组数据第一行是一个数 n (   1<=n    <=  1000)，表示
 有 n 个神秘数字，接下来一行，包含 n 个整数 a0,  a1  …   an-1，相邻之间用一个空
 格分开，对每个数字 ai ( 0   <=  ai  <=  1000,   0   <=  int<    n)。
 输出：
 你的è¾出包含 T 行，每行是一组测试数据的结果，
 对每组测试数据，输出若干个数字，也即密码，数字之间用空格分开。
 样例输入：
3
10
1 2 3 4 5 6 7 8 9 10
2
2 1
6
4 3 2 2 5 2
 样例输出：
 1   4   9   10
 2   1
 4   2   2

 */
    