<?php

$row=<<<EOF
INSERT INTO `tk_$1`.`tb_logic_que_type` VALUES (13, '朗读题', 0, 1, 1, '2,');
EOF;
for ($d=1;$d<=19;$d++){
	//echo str_ireplace($row,"$1",$d);
	echo str_replace('$1',$d,$row);
	echo PHP_EOL;
}

/*
$a="";
for ($d=1;$d<=19;$d++){
	$c = "";
	for ($i = 0; $i < 16; $i++) {
		$id0=300+$i*6;
		$id1=301+$i*6;
		$id2=302+$i*6;
		$id3=303+$i*6;
		$id4=304+$i*6;
		$s=31+$i;
		$o=128;
		$c = $c. "
			({$id0},'真题','k','2021-02-01 00:00:00','2021-02-01 00:00:00',{$o},0,{$s}),
			({$id1},'模拟题','k','2021-02-01 00:00:00','2021-02-01 00:00:00',{$o},0,{$s}),
			({$id2},'真题改编','k','2021-02-01 00:00:00','2021-02-01 00:00:00',{$o},0,{$s}),
			({$id3},'自研题目','k','2021-02-01 00:00:00','2021-02-01 00:00:00',{$o},0,{$s}),
			({$id4},'外部试题','k','2021-02-01 00:00:00','2021-02-01 00:00:00',{$o},0,{$s}),";
	}
	$c = substr($c,0,-1).";";
	$a = $a. "
	INSERT INTO tk_{$d}.`tb_paper_type` VALUES ". $c;
}
echo $a;
*/
