<?PHP
/**
 * @TODO 结果错误
 * @author wangtong1
 */

//$f = file("./s/input-2.txt");
$f = file('php://stdin');
$k=0;
$group= $f[$k++];

for($g=1; $g<=$group; $g++){
    $line= $f[$k++];
    for($l=1; $l<$line; $l++){
        $n = explode(" ", $f[$k++]);
        if(!isset($N[$n[0]]) ){
            $N[$n[0]] = new Node();
            $N[$n[0]]->value = $n[0];
        }
        if(!isset($N[$n[1]]) ){
            $N[$n[1]] = new Node();
            $N[$n[1]]->value = $n[1];
        }
        $n[2] == 1 ? $N[$n[1]]->left=$n[0] : $N[$n[1]]->right=$n[0];
    }
    print_r($N);exit;
    //$N左右节点交换
}


Class Node{
    public $left;
    public $value;
    public $right;
}

/*
树交换
题目描述：
给出一个二叉树，将树中每个节点的左右儿子交换，输出交换后的树的先序
遍历和中序遍历。比如：
转换前：
    1
    / \
2      3
      /    \
    4        5
交换后：
      1
    /    \
3        2
       /    \
    5        4
交换后的先序遍历： 1 3 2 5 4
交换后的中序遍历： 3 1 5 2 4
输入：
输入第一行是 T（1<= T <=10），表示有 T 组测试数据。
接下来的每组测试数据，第一行是一个正整数 n（0<n<1000），表示二叉树
有 n 个节点。节点用 1…n 来表示。
接下来的 n-1 行，每行三个正整数 i  j k，中间用一个空格分开，表示节点 i
的父节点是 j，其中 k=1 表示 i 是 j 的左儿子，k=2 表示 i 是 j 的右儿子。
输出：
对于每组数据，输出两行，第一行是交换后的树先序遍历，第二行是交换后
的树中序遍历，节点之间用一个空格分开。
样例输入：
1
5
3 1 2
2 1 1
5 2 2
4 2 1
样例输出：
1 3 2 5 4
3 1 5 2 4

*/
