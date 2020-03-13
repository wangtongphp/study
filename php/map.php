<?php

$a["a"] = 'a';
echo $a["b"] == "";
echo json_decode($a["b"], true);


$b = false;
echo $b["b"] == false;
echo json_decode($b["b"], true);


$c = null;
echo $c["c"] == "";
echo json_decode($c["c"], true);



