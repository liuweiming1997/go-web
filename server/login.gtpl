<html>
<head>
<title>memory something</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link rel="stylesheet" href="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/css/bootstrap.min.css">  
<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
<script src="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>

<!-- <form action="/login" method="post">
	user:<input type="text" name="user">
	money:<input type="text" name="money">
	<input type="submit" value="登录">
</form> -->

<div>
<form name = "myform"  onsubmit = "validateForm()" method = "post" action ="/login" >
<table align="center">
 <tr><td><label>&nbsp;&nbsp;user：</label></td><td><input type = "text" name = "user" class="form-control"></td></tr>
 <tr><td><label>money：</td></label><td><input type = "text" name = "money" class="form-control"></td></tr>
 <tr>
<td>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td><td><input type = "submit" value = "submit" class="btn btn-default"> 
<input type = "reset" value = "reset" class="btn btn-default"></td>
</tr>	
  </table>
  </form>
</div>

</body>
</html>