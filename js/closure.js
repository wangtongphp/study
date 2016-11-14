
//closure 1
var a = 1;
function fn(){

    var b = 10;
    
    function bar(){
        console.log(a+b);
    }
    return bar;
}();

//var x=fn(), b=100;
//x(); //11



//closure 2
