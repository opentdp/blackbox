<?php

header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS');
header('Access-Control-Allow-Headers: Content-Type, Authorization');
header('Content-Type: application/json;charset=utf-8');

include dirname(__DIR__) . '/vendor.php';

// fetch nodes

if (!$proxies = get_nodes($_GET['name'] ?? '')) {
    exit(json_encode(['error' => 'server failed']));
}

$output = [];
foreach ($proxies as &$node) {
    $config = [
        'name' => $node['name'],
        'status' => $node['status'],
        'curConns' => $node['curConns'],
        'todayTrafficIn' => $node['todayTrafficIn'],
        'todayTrafficOut' => $node['todayTrafficOut'],
        'lastStartTime' => $node['lastStartTime'],
        'lastCloseTime' => $node['lastCloseTime'],
    ];
    $config_file = dirname(__DIR__) . '/nodes/' . $node['name'] . '.php';
    if (is_file($config_file)) {
        $config = array_merge($config, include($config_file));
        $output[$node['name']] = $config;
    }
}

ksort($output);

echo json_encode($output, JSON_UNESCAPED_UNICODE | JSON_UNESCAPED_SLASHES);
