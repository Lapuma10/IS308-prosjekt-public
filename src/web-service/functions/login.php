<?php session_start();

function isLoggedIn(){
    if($_SESSION["loggedin"] == 1){
        return true;
    }else{
        return false;
    }
}

function verifyLogin(){
    if($_SESSION["loggedin"] == 1){
        return true;
    }else{
        header("location: index.php");
    }
}

function logout(){
    unset($_SESSION);
    session_destroy();
    header("location: index.php");
}