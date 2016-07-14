<?php
/**
 * @author wangtong
 */

/**
Are You Ok?
题目描述：
“我叫你三声，你敢答应吗？”
“Are you  ok? 
Are you ok? 
Are you ok? ….”
风靡网络的《Are you  ok》相信已经街知巷闻，如今你作为神曲的编曲，收
集了大量语料文本，而你需要从这些文本中按字母进行拆解并重新组合，统计最
后到底能重组成多少句神曲 are  you ok
比如原始文本是:  
Are you ok? I am  fine, thank you and what’s  your  name  ? I am Richard.
按字母拆解以后并重组（大小写不区分，不包含空格），能组成 2 句 are you  ok
输入：
输入第一行为 T，表示有 T 组测试数据， （1<=T<=10）。
接下来有 T 组测试数据，每组数据包含一行，是一段英文文本（文本长度在
10^6 以内）。
输出：
你的输出包含 T 行，每行是一组测试数据的结果，
对每组测试数据，输出一个数字，表示文本最多能重组多少句神曲。
样例输入：
2
Are you ok? I am  fine, thank you and what’s  your  name  ? I am  Richard.
How are you.
样例输出：
2
0
*/

//$f = file("./a/input-1.txt");
$f = file('php://stdin');
$line = $f[0];

for($i=1; $i<=$line; $i++){
	$arr = array();

	$row = $f[$i];
	$rowlen = strlen($row);
	for($j=0; $j<$rowlen; $j++){
		if(in_array($row{$j}, array('a','A'))){
			$arr['a'] = isset($arr['a']) ? ($arr['a']+1) : 1;
		}elseif(in_array($row{$j}, array('r','R'))){
			$arr['r'] = isset($arr['r']) ? ($arr['r']+1) : 1;
		}elseif(in_array($row{$j}, array('e','E'))){
			$arr['e'] = isset($arr['e']) ? ($arr['e']+1) : 1;
		}elseif(in_array($row{$j}, array('y','Y'))){
			$arr['y'] = isset($arr['y']) ? ($arr['y']+1) : 1;
		}elseif(in_array($row{$j}, array('o','O'))){
			$arr['o'] = isset($arr['o']) ? ($arr['o']+1) : 1;
		}elseif(in_array($row{$j}, array('u','U'))){
			$arr['u'] = isset($arr['u']) ? ($arr['u']+1) : 1;
		}elseif(in_array($row{$j}, array('k','K'))){
			$arr['k'] = isset($arr['k']) ? ($arr['k']+1) : 1;
		}
	}
	//print_r($arr);
	if (count($arr) < 7 || $arr['o'] < 2 ){
		$res[$i] = 0;
	}else{
		$arr['o'] = $arr['o'] / 2;
		sort($arr, SORT_NUMERIC);
		$one = array_shift($arr);
		$res[$i] = intval($one);
	}	
	//print_r($o_arr);
}

$out = '';
foreach($res as $v){
	$out .= $v.PHP_EOL;
	echo $v.PHP_EOL;	
}
//file_put_contents("./a/out-1.txt",$out);
