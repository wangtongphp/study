<?php
/**
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

