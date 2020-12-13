<?php session_start();

function queryUser(){
        $conn = dbConnect();
        $sql = "SELECT email,shopify_name,api_key_shopify,api_key_fiken,company_slug
                FROM Users
                INNER JOIN Credentials on Credentials.user_id = Users.user_id
                WHERE Users.user_id = ?";
        $stmt = $conn->prepare($sql);
        $stmt->bind_param("i",$_SESSION["memberID"]);
        $stmt->execute();
        return $stmt
                    ->get_result()
                    ->fetch_object();
}

function queryIDFromEmail($email){
    $conn = dbConnect();
    $sql = "SELECT user_id
                FROM Credentials
                WHERE email = ?";
    $stmt = $conn->prepare($sql);
    $stmt->bind_param("s",$email);
    $stmt->execute();
    return $stmt
        ->get_result()
        ->fetch_object()
        ->user_id;

}