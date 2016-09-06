<?PHP

$a = <<<EOF
< ' " > / \ : ( ) _ % like null nul & ;
EOF;

$mysqli = new mysqli("172:3306",'','','');
echo htmlspecialchars($a).PHP_EOL;
echo (addslashes($a)).PHP_EOL;
echo $mysqli->real_escape_string($a).PHP_EOL;


echo htmlspecialchars(addslashes($a)).PHP_EOL;

$arg = 3; // $_GET['arg'];
$sql ="";
echo $sql.PHP_EOL;


$q = $mysqli->query($sql);

include("./1.html");
