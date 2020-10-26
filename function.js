function Send(url,information) {
    var httpRequest =new XMLHttpRequest();
    httpRequest.open("POST",url,true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    httpRequest.send(JSON.stringify(information));
    httpRequest.onreadystatechange =()=>{
        if (httpRequest.readyState == 4 && httpRequest.status==200){
            var data=JSON.parse(httpRequest.responseText);
            if(data.code==200){
                alert("登陆成功！正在跳转......");
                 add1();
            }else{
                alert(data.msg);
            }
        }
    }
}
function validateForm(x,y) {
    if ((x == null || x == "")) {
        alert("需要输入账号和密码！");
        return false;
    }
    else{
        return true
    }
}
function myLogin() {
    console.log("data");
    var num1;
    var password1;
    var num = document.getElementById("username");
    var password = document.getElementById("password");
    num1=num.value;
    password1=password.value;
    console.log(num1,password1);
   var b= validateForm(num1,password1);
   if(b){
       var url = "http://localhost:8080/Login";
       Send(url,{"num":num1,"password":password1});
   }else{
       return false
   }
}
function add(){
    window.location.href="http://localhost:63342/mygoProject/main/my/HTML/register.html?_ijt=on3l2j3tk402qepst0n4k96huj"; //你需要跳转的地址
}
function add1(){
    window.location.href="http://localhost:63342/mygoProject/main/my/HTML/show.html?_ijt=rhhlm4ea85bkc8vhpcjrbf0evi"; //你需要跳转的地址
}

function MyReset() {
        var num1;
        var password1;
        var num = document.getElementById("username1");
        var password = document.getElementById("password1");
        num1=num.value;
        password1=password.value;
        console.log(num1,password1);
        var url = "http://localhost:8080/register";
       Send(url,{"num":num1,"password":password1});
    }
    function myShow() {
        var oPNode = document.getElementById("temp");//获取p节点bai的对象
        for(i =0;i<10;i++){
            var node = document.createElement('p');//创建一个文本节点
            node.innerHTML = "hello world!";//文本内容du
            oPNode.appendChild(node);//将创建的文本内容插zhi入到p节点中dao，追加方式
        }

    }