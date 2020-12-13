<?php session_start();
require_once "../functions/login.php";
include_once "../functions/query/queryUsers.php";
include_once "../functions/query/queryLogs.php";
require_once "../functions/database_connect.php";
require_once "../functions/update/updateUser.php";
verifyLogin();
if(isset($_POST["logout"])){
    logout();
}
$user = queryUser();
if(isset($_POST["Update"])){
    switch (true){
        case $_POST["email"] != $user->email:
            updateEmail($_POST["email"],$_SESSION["memberID"]);
            break;

        case $_POST["Shopify_Name"] != $user->shopify_name:
            UpdateShopifyName($_POST["Shopify_Name"],$_SESSION["memberID"]);
            break;

        case $_POST["API_Key_Shopify"] != $user->api_key_shopify:
            UpdateAPIShopify($_POST["API_Key_Shopify"],$_SESSION["memberID"]);
            break;

        case $_POST["API_Key_Fiken"] != $user->api_key_fiken:
            UpdateAPIFiken($_POST["API_Key_Fiken"],$_SESSION["memberID"]);
            break;

        case $_POST["Company_Slug"] != $user->company_slug:
            UpdateCompanySlug($_POST["Company_Slug"],$_SESSION["memberID"]);
            break;

    }
}
$user = queryUser();

?>

<head>
    <title>Welcome</title>
</head>
<body>
<h2>Your details</h2>
<table border="1">
        <tr>
            <td>Email</td>
            <td>Shopify name</td>
            <td>Shopify API</td>
            <td>Fiken API</td>
            <td>Company slug</td>
        </tr>
    <?php
        echo "<tr>";
        echo "<td>" . $user->email . "</td>";
        echo "<td>" . $user->shopify_name . "</td>";
        echo "<td>" . $user->api_key_shopify . "</td>";
        echo "<td>" . $user->api_key_fiken . "</td>";
        echo "<td>" . $user->company_slug . "</td>";
     ?>
</table>

<?php ?>
<h2>Update details</h2>
<form method="post">
    <p>Email</p>
    <input type="email" name="email" value="<?php echo $user->email; ?>" required>
    <p>Shopify name</p>
    <input type = "text" name = "Shopify_Name" value="<?php echo $user->shopify_name ?>" required>
    <p>API Key Shopify</p>
    <input type = "text" name = "API_Key_Shopify" value="<?php echo $user->api_key_shopify ?>" required>
    <p>API Key Fiken</p>
    <input type = "text" name = "API_Key_Fiken" value="<?php echo $user->api_key_fiken ?>" required>
    <p>Company slug</p>
    <input type = "text" name = "Company_Slug"  value="<?php echo $user->company_slug ?>" required> <br><br>
    <input type="submit" name="Update" value="Update">

</form>

<p>Action was last performed: <?php echo queryLastCalled($_SESSION["memberID"])->last_called_date ?></p>



<br><br>
<form method="post">
    <input type="submit" value="Logout" name="logout">
</form>
</body>
