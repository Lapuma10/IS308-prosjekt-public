<?php
/**
 * @param $userID
 * @param $ShopifyName
 * @param $APIKeyShopify
 * @param $APIKEYFiken
 * @param $CompanySlug
 */
function addUser($userID,$ShopifyName, $APIKeyShopify, $APIKEYFiken, $CompanySlug){
    $conn = dbConnect();
    $stmt = $conn->prepare("INSERT INTO Users VALUES (?, ?, ?, ?, ?)");
    $stmt->bind_param("issss",$userID, $ShopifyName, $APIKeyShopify, $APIKEYFiken, $CompanySlug);
    echo $stmt->error;
    $stmt->execute();
}

/**
 * @param $email
 * @param $password
 * @return int
 */
function addCredentials($email,$password){
    $conn = dbConnect();
    $stmt = $conn->prepare("INSERT INTO Credentials(email,password) VALUES(?,?)");
    echo $stmt->error;
    $stmt->bind_param("ss", $email, $password);
    $stmt->execute();
    echo $stmt->error;
    return $stmt->insert_id;
}

function addJob($userID){
    $jobType = "update-user";
    $intervalDays = 2;
    $conn = dbConnect();
    $stmt = $conn->prepare("INSERT INTO Cronjobs(job_type,interval_days,user_id) VALUES(?,?,?)");
    $stmt->bind_param("ssi", $jobType, $intervalDays, $userID);
    $stmt->execute();
    return $stmt->insert_id;
}


/**
 * @param $email
 * @param $password
 * @return bool
 */
function verifyCredentials($email,$password){
    $conn = dbConnect();
    $stmt = $conn->prepare("SELECT * FROM Credentials WHERE email = ? and password = ?");
    echo $stmt->error;
    $stmt->bind_param("ss",$email,$password);
    $stmt->execute();
    echo $stmt->error;
    if($stmt->get_result()->num_rows == 1){
        return true;
    }else{
        return false;
    }
}