<?php
/**
 * @status 通过
 * @author wangtong1@xiaomi.com
 */

$data = file_get_contents("php://stdin");
$datas = explode(PHP_EOL,$data);
$k = 0;
$group = $datas[$k++];
$max = pow(2,32)/2-1;
for($g=0; $g<$group; $g++){

    $d = explode(" ",$datas[$k++]);
    $t = 1;
    for($i=0;$i<$d[1];$i++){
        $t = $t*$d[0];
        if($t >= $max){
            echo 'overflow';
            break;
        }
        if($i == $d[1]-1){
            echo 'OK';
        }
    }
    echo PHP_EOL;
    
}




/**


B.  边界检查
题目描述：
Saerdna  最近在研究 golang  的编译过程,对 go  的类型推导和类型检查很
感兴趣.打算给 golang  增加一个 int  整型越界的编译检查功能.怎奈能力不足,你
可以帮一下他么.
输入：
输入第一行为 T，表示有 T 组测试数据，（1<=T<=100）。
对每组数据，每行 2 个整数 a,b (0<=a<=100, 0<=b<=100),以空格分割
输出：
对于每组数组，如果 �!的结果可以用 int 类型进行保存,则输出 OK
如果结果不能用 int  类型进行存储,则输出 overflow
样例输入：
4
2 3
2 30
2 31
3 99
样例输出：
OK
OK
overflow
overflow
Hint:
由于经费原因,Saerdna 手头只有 32 位的小霸王,所以默认环境是 32 位机器

*/
