<?php
/**
 * @status 通过
 * @author wangtong1
 */


//echo str_pad(decbin(ip2long("77.8.143.64")), 32, "0", STR_PAD_LEFT);
//echo PHP_EOL;
//echo substr(decbin(ip2long("77.8.143.64")),0,26);
//echo PHP_EOL;
//echo substr(decbin(ip2long("77.8.143.99")),0,26);
//exit;

//$f = file("./s/input-1.txt");
$f = file('php://stdin');
$line = array_shift($f);

for($i=1; $i<=$line; $i++){
	$ip_main = array_shift($f);
	$ip_main_arr = explode("/",$ip_main);
	$ip_main_int = str_pad(decbin(ip2long($ip_main_arr[0])),32,"0",STR_PAD_LEFT);
    $bitsize = (int)$ip_main_arr[1] ;
	$diff = substr($ip_main_int,0,$bitsize);
	//echo $bitsize;exit;
	$ip_nums = array_shift($f);
	for($j=0; $j<$ip_nums; $j++){
		$ip = array_shift($f);
		$ip_int = str_pad(decbin(ip2long(trim($ip))),32,"0",STR_PAD_LEFT);
		$ip_sub = substr($ip_int,0,$bitsize);
        //echo $ip. ' '.ip2long(trim($ip));
		//echo $ip_int.PHP_EOL.$ip_sub.PHP_EOL.$diff;
        //echo $diff.PHP_EOL.$ip_sub.' ';
        //echo $ip_main.' '.$ip.' ';
		if($diff === $ip_sub){
			echo 'y'.PHP_EOL;
		}else{
			echo 'n'.PHP_EOL;
		}
	}
}

/**
A. 子网掩码 
题目描述:   
    
众所周知,子网掩码的作用,主要是将 IP 地址区分为网络地址和主机地址
两个部分,举个栗子:  
    
210.73.15.0/24  
    
定义了一个子网络,这个网络下的所有 IP 地址,化为二进制以后,前 24 位
跟 210.73.15.0 一致。   
    
现在给出子网的定义,以及一堆 IP 地址,判断这个 IP 是否属于这个子网。  
    
    
输入:   
    
输入第一行为 T,表示有 T 组测试数据,  (1<=T<=10)。   
    
接下来有 T 组测试数据,每组数据第一行是一个子网的定义,符合 a.b.c.d/m
的格式,比如 210.73.15.0/24,其中 0<m<=32,a.b.c.d 是一个合法的 IP 地址。  
    
接下来一行是一个正整数 n    (   0<  n   <=  100),表示接下来有 n 个 IP 地址。    
    
接下来的 n 行,每行一个合法的 IPV4 地址,满足 a.b.c.d 的格式要求。    
    
输出:   
    
对于每一组数据的每个 IP 地址,输出 y,表示属于所定义的子网,输出 n 表
示不属于。  
    
    
样例输入:   
2   
210.73.15.0/24  
3   
210.73.15.1 
210.73.15.2 
210.73.16.1 
10.10.0.0/16    
2   
10.10.1.1   
192.168.31.1    
    
样例输出:   
y   
y   
n   
y   
n


*/
