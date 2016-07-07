<?PHP

/**
 * @TODO 运行超时
 * @author wangtong1
 */

$data = file_get_contents("php://stdin");
$datas = explode(PHP_EOL,$data);
array_pop($datas);
$group = array_shift($datas);
for($g=0; $g<$group; $g++){
    //省配置
    $provs = array_shift($datas);
    for($i=0;$i<$provs;$i++){
        $p = array_shift($datas);
        $ps = explode(" ", $p);
        $pr[] = array('prov'=>$ps[2], 'min'=>ip2long($ps[0]), 'max'=>ip2long($ps[1]));
        $res[$ps[2]] = 0;
    }
    $pr = array_sort($pr,'min'); 

    //ip
    $ips = array_shift($datas);
    for($j=0;$j<$ips;$j++){
        $ip = ip2long(array_shift($datas));
        //二分查找 
        $low = 0;
        $high = count($pr)-1;
        while($low <= $high){
            $mid = intval(($low+$high)/2);
            if($ip >=  $pr[$mid]['min'] && $ip <= $pr[$mid]['max']){
                $res[$pr[$mid]['prov']] += 1;
                break;
            }elseif($ip < $pr[$mid]['min']){
                $high = $mid-1;
            }elseif($ip > $pr[$mid]['max']){
                $low = $mid+1;
            }
        }

    }
    ksort($res);
    foreach($res as $k=>$v){
        echo $k." ".$v.PHP_EOL;
    }



}


    function array_sort($arr, $col = "", $order = "SORT_ASC")
    {
        $new_array  = array();
        $sort_array = array();

        if(is_array($arr) && !empty($arr))
        {
            foreach($arr as $k => $v)
            {
                if(is_array($v))
                {
                    foreach ($v as $kk => $vv)
                    {
                        if($kk == $col)
                        {
                            $sort_array[$k] = $vv;
                        }
                    }
                }
                else
                {
                    $sort_array[$k] = $v;
                }
            }

            switch ($order)
            {
                case "SORT_ASC":
                    asort($sort_array);
                break;
                case "SORT_DESC":
                    arsort($sort_array);
                break;
                default :
                    return array();
                    break;
            }

            foreach ($sort_array as $k => $v) {
                $new_array[$k] = $arr[$k];
            }
        }

        return $new_array;
    }
//var_dump($datas);

/**
IP 地址统计
题目᧿述：
数据中心有一些 ip 地址库，包括每个地址区间以及归属省份(ip 地址区间不
会重叠交叉)，现在要对用户访问日志统计每个省份的访问数量。
输入：
输入第一行为 T，表示包含 T 组测试数据。
对于每组测试数据，第一行为 n（1<=n<=500000），表示 ip 地址库的 ip 区
间数量。
接下来的 n 行，每行包含三个数据，ip_start,   ip_end，province_id，表示 ip
区间的开始区间，结束区间，以及对应的省份 id。ip 区间为闭区间，而且各个
区间不会重叠交叉。ip 地址以 a.b.c.d 形式ᨀ供。
接下来是 m(1<=m<=500000)，表示用户访问日志数。
接下来的 m 行，每行是一个 ip 地址，表示用户访问的 ip 记录。
输出：
针对每组测试数据，按省份 id 输出每个省的访问数，每行第一个为省份 id，
第二个为该省份的访问数量，中间用空格分开。如果某个省份没有访问过，则访
问次数为 0。每组数据按省份 id 从小到大输出。
样例输入：
2
1
192.168.12.2 192.168.12.3 1
2
10.100.108.34
192.168.12.2
3
10.1.1.1 10.2.255.255 1
223.199.12.2 224.200.19.12 2
172.2.12.0 172.16.2.2 3
3
10.2.123.123
10.2.255.255
223.199.13.255
样例输出：
1   1
1   2
2   1
3   0
*/
