<?php


/**
 * https://stackoverflow.com/questions/26812248/php-flush-dont-work
 */
ini_set('output_buffering', 0);
ini_set('zlib.output_compression', 0);
if( !ob_get_level() ){ ob_start(); }
else { ob_end_clean(); ob_start(); }
for ($i = 0; $i < 10; $i++) {
  //For Nginx we have to reach minimum  buffer size, 
  //so if it is not enough increment output
  echo str_pad( $i . '<br>', 1024 + 10, ' ', STR_PAD_RIGHT ); 
  ob_flush();
  flush();
  sleep(1);
}

/*
for ($i = 0; $i < 3; $i++) {
echo str_repeat('0', 4096). $i . '<br/>';
sleep($i +1);
}
*/

/**
 * php.net
 */
/*
header("Cache-Control: no-cache, must-revalidate");
header("Expires: Mon, 26 Jul 1997 05:00:00 GMT");


echo str_repeat('0', 4096);

if (ob_get_level() == 0) ob_start();


for ($i = 0; $i<3; $i++){

        echo "<br> Line to show.";
        echo str_pad('',4096)."\n";    

        ob_flush();
        flush();
        sleep(2);
}

echo "Done.";

ob_end_flush();
*/


/*
 @apache_setenv('no-gzip', 1);
    @ini_set('zlib.output_compression', 0);
    @ini_set('implicit_flush', 1);
    for ($i = 0; $i < ob_get_level(); $i++) { ob_end_flush(); }
    o_implicit_flush(1);

    */