<?php

/**
 *
 * @param $userID
 * @return object|stdClass
 */
function queryLastCalled($userID){
    $conn = dbConnect();
    $sql = "SELECT last_called_date
                FROM Cronjobs
                WHERE user_id = ?";
    $stmt = $conn->prepare($sql);
    $stmt->bind_param("i",$userID);
    $stmt->execute();
    return $stmt
        ->get_result()
        ->fetch_object();
}