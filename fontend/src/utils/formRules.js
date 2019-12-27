//检查字符串字节长度
function checkStrLength(strTemp) { 
    var i,sum;
    sum=0;
    var length = strTemp.length ;
    for(i=0;i<length;i++) { 
      if ((strTemp.charCodeAt(i)>=0) && (strTemp.charCodeAt(i)<=255)) {
        sum=sum+1;
      }else {
        sum=sum+2;
      }
    }
    return sum; 
  }


  //校验邮箱格式
  function checkEmailStr(emailTemp){
    var myReg = /^[a-zA-Z0-9_-]+@([a-zA-Z0-9]+\.)+(com|cn|net|org)$/
    return  myReg.test(emailTemp)
  }
  

  