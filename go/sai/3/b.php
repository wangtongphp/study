<?PHP
/**
 * @TODO 通过
 * @tips: 输出的preorder inorder的两行最后不能有空格，所以不能直接输出，需要用引用的形式放到变量里面，
          左右树交换，如果其中有个子树为空，依旧交换
 * @author wangtong1
 */

//$file = file_get_contents("./s/input-2.txt");
$file = file_get_contents('php://stdin');
$f = explode(PHP_EOL,$file);
$k=0;
$group= $f[$k++];

for($g=1; $g<=$group; $g++){
    $line= $f[$k++];
    //组装初始二叉树
    $N = array();
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
        
        if($n[2] == 1){
            $N[$n[1]]->left=$N[$n[0]];
        }else{
            $N[$n[1]]->right=$N[$n[0]];
        }
        $N[$n[0]]->p = $N[$n[1]]->value;
    }

    //找出最终的二叉树,$Node
    foreach($N as $kN=>$vN){
        if(!isset($vN->p)){
            $Node = $vN;
            break;
        }
    }

    //print_r($Node);
    //$N左右节点交换
    swapLR($Node);

    $preArr = array();
    preorder($Node, $preArr);
    echo implode(' ',$preArr);
    echo PHP_EOL;

    $inArr = array();
    inorder($Node, $inArr);
    echo implode(' ',$inArr);
    echo PHP_EOL;
}

//前序遍历，递归
function preorder($node, &$res){
    if(empty($node)){
        return;
    }
    $res[] = $node->value;
    //echo '_'; //最后不能输出空格
    preorder($node->left, $res);
    preorder($node->right, $res);
}

//中序遍历，递归
function inorder($node, &$res){
    if(empty($node)){
        return;
    }
    inorder($node->left, $res);
    $res[] = $node->value;
    inorder($node->right, $res);
}

//递归交换二叉树的左右节点
function swapLR(&$node){
    if(empty($node)){
        return ;
    }
    if(!empty($node->left) || !empty($node->right)){
        $tmp = $node->left;
        $node->left = $node->right;
        $node->right = $tmp;
    }
    swapLR($node->left);
    swapLR($node->right);
}

Class Node{
    public $left;
    public $value;
    public $right;
    public $p;
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
           /  \
          3     2
              /   \
             5      4
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
