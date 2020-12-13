<?php
/**
 * @param $email
 * @param $userID
 * @return
 */
function updateEmail($email,$userID){
    $conn = dbConnect();
    $sql = "Update Credentials
            SET email = ?
            WHERE user_id = ?";
    $stmt = $conn->prepare($sql);
    $stmt->bind_param("si",$email,$userID);
    $stmt->execute();
}

/**
 * @param $shopifyName
 * @param $userID
 */
function UpdateShopifyName($shopifyName,$userID){
    $conn = dbConnect();
    $sql = "Update Users
            SET shopify_name = ?
            WHERE user_id = ?";
    $stmt = $conn->prepare($sql);
    $stmt->bind_param("si",$shopifyName,$userID);
    $stmt->execute();
}

/**
 * @param $API
 * @param $userID
 */
function UpdateAPIShopify($API,$userID){
    $conn = dbConnect();
    $sql = "Update Users
            SET api_key_shopify = ?
            WHERE user_id = ?";
    $stmt = $conn->prepare($sql);
    $stmt->bind_param("si",$API,$userID);
    $stmt->execute();
}

/**
 * @param string $API
 * @param $userID
 */
function  UpdateAPIFiken($API,$userID){
    $conn = dbConnect();
    $sql = "Update Users
            SET api_key_fiken = ?
            WHERE user_id = ?";
    $stmt = $conn->prepare($sql);
    $stmt->bind_param("si",$API,$userID);
    $stmt->execute();
}

/**
 * @param $companySlug
 * @param $userID
 */
function UpdateCompanySlug($companySlug,$userID){
    $conn = dbConnect();
    $sql = "Update Users
            SET company_slug = ?
            WHERE user_id = ?";
    $stmt = $conn->prepare($sql);
    $stmt->bind_param("si",$companySlug,$userID);
    $stmt->execute();
}

