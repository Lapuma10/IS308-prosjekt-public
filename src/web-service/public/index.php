<?php session_start();
require_once "../functions/database_connect.php";
require_once "../functions/query/queryUsers.php";
require_once "../functions/insert/insertUser.php";
include_once "../functions/login.php";
if (isLoggedIn()){
    header("location: welcome.php");
}

if(isset($_POST["submit"])){
    if(verifyCredentials($_POST["email"],$_POST["password"])){
        $_SESSION["memberID"] = queryIDFromEmail($_POST["email"]);
        $_SESSION["loggedin"] = 1;
        header("location: welcome.php");
    }
}

?>
<html>
<body>
<form method="post" action="index.php">
    <p>Email</p>
    <input type="email" name="email">
    <p>Password</p>
    <input type="password" name="password">
    <input type="submit" name="submit" value="Login">
</form>
<a href="register.php">Register</a>


</body>
</html>