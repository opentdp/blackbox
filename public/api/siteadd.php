<?php

header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: POST');
header('Access-Control-Allow-Headers: Content-Type');
header('Content-Type: application/json;charset=utf-8');

include dirname(__DIR__) . '/vendor.php';

// check formdata

$post = json_decode(file_get_contents('php://input'), true);

if (empty($post)) {
    exit('{"error":1,"message":"表单项不能为空"}');
}

if (empty($post['secret'])) {
    exit('{"error":1,"message":"密钥不能为空"}');
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

if (empty($post['project'])) {
    exit('{"error":1,"message":"名称不能为空"}');
}

if (!empty($post['logo']) && !filter_var($post['logo'], FILTER_VALIDATE_URL)) {
    exit('{"error":1,"message":"图标地址无效"}');
}

if (empty($post['email']) || !filter_var($post['email'], FILTER_VALIDATE_EMAIL)) {
    exit('{"error":1,"message":"邮箱格式错误"}');
}

if (empty($post['desc'])) {
    exit('{"error":1,"message":"描述不能为空"}');
}

// save site config

$config = [
    'targets' => $post['targets'],
    'labels' => [
        'project' => $post['project'],
        'email' => $post['email'],
        'logo' => $post['logo'],
        'desc' => $post['desc'],
    ]
];

set_site($config);

exit('{"error":0,"message":"保存成功"}');
