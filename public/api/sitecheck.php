<?php

header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: POST');
header('Access-Control-Allow-Headers: Content-Type');
header('Content-Type: application/json;charset=utf-8');

include dirname(__DIR__) . '/vendor.php';

// check formdata

$post = json_decode(file_get_contents('php://input'), true);

if (empty($post['secret'])) {
    exit('{"error":1,"message":"密码不能为空"}');
}

if (empty($post['targets'])) {
    exit('{"error":1,"message":"网址不能为空"}');
}

$post['targets'] = array_filter(explode("\n", $post['targets']));
foreach ($post['targets'] as &$target) {
    $target = trim($target, '/');
    if (!filter_var($target, FILTER_VALIDATE_URL)) {
        exit('{"error":1,"message":"站点地址无效"}');
    }
    if (!check_site($target, $post['secret'])) {
        exit('{"error":1,"message":"站点验证失败"}');
    }
}

exit('{"error":0,"message":"验证成功"}');
