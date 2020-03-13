<?php

function genTraceID() {
        return base_convert(intval(microtime(true)) * 100000 + mt_rand(0, 99999), 10, 36)
. base_convert(mt_rand(10000, 19999) * 100000000000000 + ip2long(gethostbyname(gethostname())) * 10000 + getmypid(), 10, 36);
}


//echo genTraceID();


for($i = 0; $i < 1000000; $i++) {
    echo genTraceID() . "\n";
}


