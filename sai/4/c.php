<?php
/**
 * @TODO 运行超时
 * @desc 基础思路：将所有用户喜欢的所有类目统计，循环3个类目的组合，每种组合内遍历用户喜欢的类目和金额，如果用户喜欢的类目在组合中则金额记录总金额，对比每个组合的总金额即可
    tips: 为了优化每个组合都要遍历所有用户的问题，提前汇总每个类目的所有用户，然后这里将三个类目的所有用户提取出来合并去重余额相加即是金额。
 * @author wangtong1@xiaomi.com
 */

$data = file_get_contents("php://stdin");
$datas = explode(PHP_EOL,$data);
$k = 0;
$group = $datas[$k++];
for($g=0; $g<$group; $g++){

    $d = array();
    $d = explode(" ", $datas[$k++]);

    //run
    $t = 0;
    //<4，全部余额相加
    if($d[0]<4){
        for($i=0;$i<$d[1];$i++){
            $p = explode(" ", $datas[$k++]);
            $t +=$p[3];
        }
        echo $t;
        echo PHP_EOL;
    }else{
       //cates
       $cates = array();
       $users = array();
       for($i=0;$i<$d[1];$i++){
            $p = explode(" ", $datas[$k++]);
            
            $cates[$p[0]][] = $i;
            $cates[$p[1]][] = $i;
            $cates[$p[2]][] = $i;
            $users[$i] = $p[3];
       }
       $cate_cnt = count($cates);
       $cks = array_keys($cates);
       //var_dump($cates, $cks);exit;

       //所有类目的组合
        for($ik=0;$ik<$cate_cnt-2;$ik++){
           for($ikk=$ik+1;$ikk<$cate_cnt-1;$ikk++){
                for($ikkk=$ikk+1;$ikkk<$cate_cnt;$ikkk++){
                    $t_amount = 0;
                    $user_keys = array_unique(array_merge($cates[$cks[$ik]], $cates[$cks[$ikk]], $cates[$cks[$ikkk]]));
                    //var_dump($user_keys);exit;
                    foreach($user_keys as $kv){
                            $t_amount += $users[$kv];
                    }
                    if(!isset($max_amount) || $max_amount < $t_amount){
                        $max_amount = $t_amount;
                    }
                }
           }
        }

        echo $max_amount;
        unset($max_amount);

        echo PHP_EOL;
    }
}


/**
网站布局
题目描述：
网站入口布局是一个非常考验技术的活儿,能不能吸引到用户直接决定了网
站转化率的高低.
我们假设每个用户会对三个类目感兴趣,当浏览网站页面时,如果用户发现感
兴趣的类目,用户会点击对应类目并花费他钱包中所有的余额购买相关产品,如果
前三个类目中没出现用户感兴趣的类目时,用户会选择离开.
现在问题来了,给你每个用户小米钱包中的余额数以及每个用户感兴趣的类
目信息,你能通过调整官网类目的排列顺序让用户尽可能购买我们的产品么?

输入：
第一行一个整数 T,    表示一个 T(T<=100)组测试数据
接下来 T 组测试数据
第一行是 2 个整数 N(3<=N<=15)表示官网的类目数量.一个整数
M(M<=10^6)表示来官网浏览的用户数量.
接下来 M     行,每行 4 个数,P0,P1,P2,cost.数之间用空格隔开
Px 表示用户喜欢的类目,优先级 P0>P1>P2,保证 P0,P1,P2 互不相同.   Cost 表
示用户小米钱包的余额,0<=Px<N.   cost<=10^3.
出：
对每组测试数据输出一个整数 total,    表示用户最多能产生多少销售额
样例输入：
3
3 3
0 1 2 1
1 2 0 2
2 0 1 4
3 3
0 1 2 1
1 0 2 2
2 1 0 4
12 4
0 1 2 1
3 4 5 2
6 7 8 4
11 10 9 8

样例输出：
7
7
14

*/
