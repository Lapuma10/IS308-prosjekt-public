<?php
    function dbConnect(){
    $host = 'db';
    $username = "root";
    $password = "BSBACIT2020";
    $dbname = "project";
    $conn = new mysqli($host, $username, $password, $dbname);
    if ($conn->connect_error) {
        die("Connection failed: " . $conn->connect_error);
    }
    return $conn;
}
