<?php

header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: POST');
header('Access-Control-Allow-Headers: Content-Type');
header('Content-Type: application/json;charset=utf-8');

include dirname(__DIR__) . '/vendor.php';

$post = json_decode(file_get_contents('php://input'), true);

if (empty($post) || empty($post['project']) || empty($post['targets']) || empty($post['email']) || empty($post['desc'])) {
    exit('{"error":1,"message":"表单项不能为空"}');
}

$post['targets'] = array_filter(explode("\n", $post['targets']));
foreach ($post['targets'] as $target) {
    if (!filter_var($target, FILTER_VALIDATE_URL)) {
        exit('{"error":1,"message":"站点地址校验失败"}');
    }
}

if (!empty($post['logo']) && !filter_var($post['logo'], FILTER_VALIDATE_URL)) {
    exit('{"error":1,"message":"图标地址校验失败"}');
}

if (!filter_var($post['email'], FILTER_VALIDATE_EMAIL)) {
    exit('{"error":1,"message":"邮箱校验失败"}');
}

$config = [
    'targets' => $post['targets'],
    'labels' => [
        'project' => $post['project'],
        'email' => $post['email'],
        'desc' => $post['desc'],
    ]
];

set_site($config);

exit('{"error":0,"message":"保存成功"}');
