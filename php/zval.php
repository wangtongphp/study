<?PHP
$arr1 = array(1);
echo "\nbefore:\n";
echo "\$arr1[0] == {$arr1[0]}\n";
$arr2 = $arr1;
$arr2[0]++;
echo "\nafter:\n";
echo "\$arr1[0] == {$arr1[0]}\n";
echo "\$arr2[0] == {$arr2[0]}\n";

// Example two
$arr3 = array(1);
$a =& $arr3[0];
echo "\nbefore:\n";
echo "\$a == $a\n";
echo "\$arr3[0] == {$arr3[0]}\n";
$arr4 = $arr3;
$arr4[0]++;
echo "\nafter:\n";
echo "\$a == $a\n";
echo "\$arr3[0] == {$arr3[0]}\n";
echo "\$arr4[0] == {$arr4[0]}\n";



// Example
$arr5 = 5;
$arr6 =& $arr5;
echo "\nbefore:\n";
echo "\$arr6 == $arr6\n";
echo "\$arr5 == {$arr5}\n";
$arr7 = $arr5;
$arr7++;
echo "\nafter:\n";
echo "\$arr6 == $arr6\n";
echo "\$arr5 == {$arr5}\n";
echo "\$arr7 == {$arr7}\n";

$h = array(1,2);
$j =& $h[0];
$i = $h;
//$i[0]++;
$i[0] = $i[0] + 1;
var_dump($h,$i);

$m = 3;
$n =& $m;
$m = 4;
echo $n;


echo PHP_EOL;
$o = 3;
$p =& $o;
unset($o);
xdebug_debug_zval('o');
xdebug_debug_zval('p');
echo PHP_EOL;
