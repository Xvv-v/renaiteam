
var storage = window.localStorage;
function SendLogin(url,information) {
    var httpRequest =new XMLHttpRequest();
    httpRequest.open("POST",url,true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    httpRequest.send(JSON.stringify(information));
    httpRequest.onreadystatechange =()=>{
        if (httpRequest.readyState == 4 && httpRequest.status==200){
            var data=JSON.parse(httpRequest.responseText);
            if(data.code==200){
                alert("登陆成功！正在跳转......");
                 addshow();
            }else{
                alert(data.msg);
                location.reload();
            }
        }
    }
}
function SendRehister(url,information) {
    var httpRequest =new XMLHttpRequest();
    httpRequest.open("POST",url,true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    httpRequest.send(JSON.stringify(information));
    httpRequest.onreadystatechange =()=>{
        if (httpRequest.readyState == 4 && httpRequest.status==200){
            var data=JSON.parse(httpRequest.responseText);
            if(data.code==200){
                alert("注册成功！正在跳转......");
                addLogin();
            }else{
                alert(data.msg);
                location.reload();
            }
        }
    }
}
function SendShow(url,information) {
    var httpRequest =new XMLHttpRequest();
    httpRequest.open("POST",url,true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    httpRequest.send(JSON.stringify(information));
    httpRequest.onreadystatechange =()=>{
        if (httpRequest.readyState == 4 && httpRequest.status==200){
            var data=JSON.parse(httpRequest.responseText);
            var oPNode = document.getElementById("temp");//获取p节点bai的对象
            for (i=0;i<data.length;i++){
                var node = document.createElement('p');//创建一个文本节点
                node.innerHTML =data[i].num+data[i].text;
                //var node1 = document.createElement('tr');
                oPNode.appendChild(node);//将创建的文本内容插zhi入到p节点中dao，追加方式
            }
        }
    }
}
function SendAdd(url,information) {
    var httpRequest = new XMLHttpRequest();
    httpRequest.open("POST", url, true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    httpRequest.send(JSON.stringify(information));
    httpRequest.onreadystatechange = () => {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            var data = JSON.parse(httpRequest.responseText);
            if(data.code==200){
                alert("添加成功！正在加载......");
                location.reload();
                //myShow();
            }else{
                alert(data.msg);
                location.reload();
            }
        }
    }

}
function SendDel(url,information) {
    var httpRequest = new XMLHttpRequest();
    httpRequest.open("POST", url, true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    httpRequest.send(JSON.stringify(information));
    httpRequest.onreadystatechange = () => {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            var data = JSON.parse(httpRequest.responseText);
            if(data.code==200){
                alert("删除成功！正在加载......");
                location.reload();
                //myShow();
            }else{
                alert(data.msg);
                location.reload();
            }
        }
    }
}
function SendMod(url,information) {
    var httpRequest = new XMLHttpRequest();
    httpRequest.open("POST", url, true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    httpRequest.send(JSON.stringify(information));
    httpRequest.onreadystatechange = () => {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            var data = JSON.parse(httpRequest.responseText);
            if(data.code==200){
                alert("修改成功！正在加载......");
                location.reload();
            }else{
                alert(data.msg);
                location.reload();
            }
        }
    }
}
function validateForm(x,y) {
    if ((x == null || x == "")) {
        alert("需要相应的信息（例如：输入账号和密码编号等）！");
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
    var storage = window.localStorage;
    storage["username"]=num1;
    password1=password.value;
    console.log(num1,password1);
   var b= validateForm(num1,password1);
   if(b){
       var url = "http://localhost:8080/Login";
       SendLogin(url,{"accout":num1,"password":password1});
   }else{
       return false
   }
}
function addregister(){
    window.location.href="http://localhost:63342/mygoProject/main/my/HTML/register.html?_ijt=on3l2j3tk402qepst0n4k96huj"; //你需要跳转的地址
}
function addshow(){
    window.location.href="http://localhost:63342/mygoProject/main/my/HTML/show.html?_ijt=rhhlm4ea85bkc8vhpcjrbf0evi"; //你需要跳转的地址
}
function addLogin() {
    window.location.href="http://localhost:63342/mygoProject/main/my/HTML/Login.html?_ijt=hh48cneiaj6ul2269u37i03qj";
}

function MyReset() {
        var num1;
        var password1;
        var num = document.getElementById("username1");
        var password = document.getElementById("password1");
        num1=num.value;
        password1=password.value;
        console.log(num1,password1);
        var b= validateForm(num1,password1);
        if(b){
            var url = "http://localhost:8080/register";
            SendRehister(url,{"accout":num1,"password":password1});
        }else{
            return false
        }
}
function myShow(data) {
    var url = "http://localhost:8080/show";
    var accout=storage["username"];
    console.log(accout);
    SendShow(url,{"user":accout});
}
function myDel() {
   var num1;
   num=document.getElementById("num");
   num1=num.value;
    var accout=storage["username"];
    console.log(accout);
   if(validateForm(num1,null)){
       var url = "http://localhost:8080/delete";
       SendDel(url,{"num":num1,"user":accout});
   }else{
       return false
   }

}
function myAdd() {
   var num1;
   var text1;
   var num = document.getElementById("num1");
   var text = document.getElementById("text1");
   num1=num.value;
   text1=text.value;
   if(validateForm(num1,null)){
       var url = "http://localhost:8080/addition";
       var accout=storage["username"];
       console.log(accout);
       SendAdd(url,{"num":num1,"text":text1,"user":accout});
   }else{
       return false
   }

}
function myMod() {
    var num1;
    var text1;
    var num = document.getElementById("num2");
    var text = document.getElementById("text2");
    num1=num.value;
    text1=text.value;
    if(validateForm(num1,null)){
        var url = "http://localhost:8080/modification";
        var accout=storage["username"];
        console.log(accout);
        SendMod(url,{"num":num1,"text":text1,"user":accout});
    }else{
        return false
    }
}