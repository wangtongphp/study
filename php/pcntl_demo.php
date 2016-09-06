<?php


define('PROC_NUM',3);
define('TIME_OUT',33);
$ppid = posix_getpid();

for($i=0; $i<PROC_NUM; $i++){
    sleep(1);
    $pid = pcntl_fork();
    if($pid == 0){
        echo 'child proc done ; ';
        exit(0);
    } 
}
$start = time();
$n = 0;
while( $n < PROC_NUM && time()-$start<TIME_OUT){
    $status = -1;
    $npid = pcntl_wait($status, WNOHANG);
    if($npid>0){
        echo $npid.' exit; ';
        ++$n;
    }
    sleep(1);
}   

echo 'Finished';


/*

define('PROC_NUM', 12);
define('SRC_NUM', 20);
define('TIME_OUT', 6);
$arr = array();
for ($i = 0; $i < SRC_NUM; ++$i) {
    $arr[] = $i;
}
$sArr = serialize($arr);
 
echo "parent id: " . posix_getpid() . "\n";
 
// create share memory
$nShmID = shm_attach(ftok(__FILE__, 'i') . + posix_getpid(), strlen($sArr) * 2);
if ($nShmID == FALSE) {
    die("Failed to create shm\n");
}
 
// write the array to the shared memory
$nArrKey = 1;
if (FALSE == shm_put_var($nShmID, $nArrKey, $arr)) {
    die('Failed to write array to shm');
}
 
// create semphore
$nSemID = sem_get(posix_getpid(), 1); // 采用主进程id作为key，避免主进程之间干扰
 
// child process consume the data in the shm
for($i = 0; $i < PROC_NUM; ++$i) {
    $nPID = pcntl_fork();
     
    if ($nPID == 0) {
        // child
        $arrProcess = array();
        while (true) {
            sem_acquire($nSemID);
             
            // get the value
            $arrCur = shm_get_var($nShmID, $nArrKey);
             
            if (0 == count($arrCur) || $arrCur == FALSE) {
                // value out
                sem_release($nSemID);
                break;
            }
            $nVal = array_pop($arrCur);
            if (FALSE == shm_put_var($nShmID, $nArrKey, $arrCur)) {
                die('Failed to write array to shm');
            }
            sem_release($nSemID);
             
            // simulate process value
            sleep(rand(1,3));
            $arrProcess[] = $nVal;
        }
         
        echo "Child process " . posix_getpid() . " consume: ";
        echo implode(' ',$arrProcess) . "\n";
        exit(0); // 子进程逻辑一定要在这里退出，不能执行父进程代码
    }
}
 
// wait for children
echo "wait children\n";
$n = 0;
$nStart = time();
while ($n < PROC_NUM && (time() - $nStart < TIME_OUT)) { // 父进程等待有超时机制，避免无限等待
    //echo "wait ... \n";
    $nStatus = -1;
    $nPID = pcntl_wait($nStatus, WNOHANG);
    if ($nPID > 0) {
        echo "{$nPID} exit\n";
        ++$n;
    }
    sleep(1); // 避免忙等
}
 
// clear shm
//shm_detach($nShmID);
sem_remove($nSemID);
shm_remove($nShmID);
echo "finished\n";

define("PC","10");
define("TO",4);
define("TS",4);
 

$sPipePath = "my_pipe.".posix_getpid();
if (!posix_mkfifo($sPipePath, 0666)) {
    die("create pipe {$sPipePath} error");
}
 
// 模拟任务并发
for ($i = 0; $i < PC; ++$i ) {
    $nPID = pcntl_fork(); // 创建子进程
    if ($nPID == 0) {
        // 子进程过程
        sleep(rand(1,TS)); // 模拟延时
        $oW = fopen($sPipePath, 'w');
        fwrite($oW, $i."\n"); // 当前任务处理完比，在管道中写入数据
        fclose($oW);
        exit(0); // 执行完后退出
    }
}
 
// 父进程
$oR = fopen($sPipePath, 'r');
stream_set_blocking($oR, FALSE); // 将管道设置为非堵塞，用于适应超时机制
$sData = ''; // 存放管道中的数据
$nLine = 0;
$nStart = time();
while ($nLine < PC && (time() - $nStart) < TO) {
    $sLine = fread($oR, 1024);
    if (empty($sLine)) {
        continue;   
    }   
     
    echo "current line: {$sLine}\n";
    // 用于分析多少任务处理完毕，通过‘\n’标识
    foreach(str_split($sLine) as $c) {
        if ("\n" == $c) {
            ++$nLine;
        }
    }
    $sData .= $sLine;
}
echo "Final line count:$nLine\n";
fclose($oR);
unlink($sPipePath); // 删除管道，已经没有作用了
 
// 等待子进程执行完毕，避免僵尸进程
$n = 0;
while ($n < PC) {
    $nStatus = -1;
    $nPID = pcntl_wait($nStatus, WNOHANG);
    if ($nPID > 0) {
        echo "{$nPID} exit\n";
        ++$n;
    }
}
 
// 验证结果，主要查看结果中是否每个任务都完成了
$arr2 = array();
foreach(explode("\n", $sData) as $i) {// trim all
    if (is_numeric(trim($i))) {
        array_push($arr2, $i);  
    }
}
$arr2 = array_unique($arr2);
if ( count($arr2) == PC) {  
    echo 'ok'; 
} else {
    echo  "error count " . count($arr2) . "\n";
    var_dump($arr2);
}

/*
function tick_handler()
{
      echo "tick_handler() called\n";
}

$a = 1;
tick_handler();

if ($a > 0) {
        $a += 2;
            tick_handler();
                print($a);
                    tick_handler();
}
tick_handler();


declare(ticks=1);

// A function called on each tick event
function tick_handler()
{
        echo "tick_handler() called\n";
}

register_tick_function('tick_handler');

$a = 1;


            print($a);


    set_time_limit(0);
    
    function profiler($return = false)
    {
        static $m = 0;
        if($return) return $m . " bytes";
        if(($mem = memory_get_usage()) > $m) $m = $mem;
    }
    
    register_tick_function('profiler');
    declare(ticks = 1);
    
    $numbers = array();
    for($i=0; $i<10; $i++)
    {
        print($i . "<br />");
    }
    
    print(profiler(true));

*/
