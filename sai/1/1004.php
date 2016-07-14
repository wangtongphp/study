<?PHP
/**
 * @author wangtong1
 * @TODO 运行超时
 */

$d = file_get_contents("php://stdin");
$ds = explode(PHP_EOL,$d);
$group = array_shift($ds);
for($i=0;$i<$group;$i++){
    $line = array_shift($ds);
    for($j=0;$j<$line;$j++){
        $cmd = explode(" ",array_shift($ds));
        if($cmd[0] == 'set'){
            $r[$cmd[1]] = $cmd[2];
        }elseif($cmd[0] == 'get'){
            if(isset($r[$cmd[1]])){
                echo $r[$cmd[1]];
            }else{
                echo 'nil';
            }
            echo PHP_EOL;
        }elseif($cmd[0] == 'min'){
            if(empty($r)){
                echo 'nil';
            }else{
               asort($r);
               echo current($r);
            }
            echo PHP_EOL;
        }elseif($cmd[0] == 'del'){
            unset($r[$cmd[1]]);
        }
    }
}

//var_dump($r);


/**
数据回放
题目描述：
数据恢复是所有 dba 的噩梦，在某次数据库灾难之后，dba 终于挽救回一批
数据，但只找回了一些数据库操作日志，你需要根据这些操作日志，重放整个数
据库的操作结果。
数据库的操作命令有以下几种：
1. set  <key> <value>
对 key  设置对应的 value 值（所有 value 都是整数）
2. get  <key>
获取 key 的 value 值，并输出。如果 key 不存在，则输出 nil
3. del  <key>
删除 key
4. min
输出整个数据库中 value 最小的值，如果数据库为空，则输出 nil
所有 key 都是由 a-z，0-9 字符组成的字符串， value 为整数且 0 <= value  < 
2^32。
输入：
输入第一行是 t(1<=t<=11)，表示有 t 组数据。
接下来 t 组数据，每组第一行是 n(1<=n<=500000)， 表示有 n 个操作。
接下来的 n 行，每行是上述 4 个操作的其中一个。
操作命令，key，value 之间用一个空格隔开。
输出：
对于每个 get 命令，输出对应 key 的值，如果 key 已不存在，则输出 nil
对于每个 min 命令，输出当前数据库中最小的 value 值，如果数据库为空，
则输出 nil。
样例输入：
1
9
set a 10
set b 5
set c 30
get d
min
del b
min
set b 15
min
例输出：
nil
5
10
10

*/
