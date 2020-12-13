<?php
require_once "../functions/database_connect.php";
require_once "../functions/query/queryUsers.php";
require_once "../functions/insert/insertUser.php";
    if(isset($_POST['submit'])) {
        $member_id = addCredentials($_POST["email"], $_POST["password"]);
        addUser($member_id, $_POST["Shopify_Name"], $_POST["API_Key_Shopify"], $_POST["API_Key_Fiken"], $_POST["Company_Slug"]);
        addJob($member_id);
        header("location: index.php");
    }
?>


<body>
<form method="post" action="register.php">
    <p>Email</p>
    <input type="email" name="email">
    <p>Password</p>
    <input type="password" name="password">
    <p>Shopify name</p>
    <input type = "text" id= "Shopify_Name" name = "Shopify_Name" placeholder = "Shopify_Name" required>
    <p>API Key Shopify</p>
    <input type = "text" name = "API_Key_Shopify" id="API_Key_Shopify" placeholder = "API_Key_Shopify" required>
    <p>API Key Fiken</p>
    <input type = "text" name = "API_Key_Fiken" id="API_Key_Fiken" placeholder = "API_Key_Fiken" required>
    <p>Company slug</p>
    <input type = "text" name = "Company_Slug" id="Company_Slug" placeholder = "Company_Slug" required> <br><br>
    <input type="submit" name="submit" value="Register">
</form>
