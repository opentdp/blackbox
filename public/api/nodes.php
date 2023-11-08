<?php

$format = 'json';
if (isset($_GET['format']) && in_array($_GET['format'], ['yaml'])) {
    $format = $_GET['format'];
}

header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: GET, POST');
header('Access-Control-Allow-Headers: Content-Type');
header('Content-Type: application/' . $format . ';charset=utf-8');

// fetch nodes

include dirname(__DIR__) . '/vendor.php';

if (!$proxies = get_nodes($_GET['name'] ?? '')) {
    if ($format == 'json') {
        exit(json_encode(['error' => 'not found']));
        exit;
    }
    exit('error: not found');
}

$nodes = [];
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
        $nodes[$node['name']] = $config;
        touch($config_file);
    }
}

ksort($nodes);

// output

if ($format == 'json') {
    echo json_encode($nodes, JSON_UNESCAPED_UNICODE | JSON_UNESCAPED_SLASHES);
    exit;
}

if ($format == 'yaml') {
    $output = [];
    foreach ($nodes as $name => $node) {
        $output[] = "$name:";
        foreach ($node as $key => $value) {
            $output[] = "  $key: $value";
        }
    }
    echo implode("\n", $output);
    exit;
}
