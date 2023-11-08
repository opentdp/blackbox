<?php

header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: GET, POST');
header('Access-Control-Allow-Headers: Content-Type');
header('Content-Type: application/json;charset=utf-8');

include dirname(__DIR__) . '/vendor.php';

// fetch sites

if (!$sites = get_sites($_GET['name'] ?? '')) {
    exit(json_encode(['error' => 'site failed']));
}


ksort($sites);

echo json_encode($sites, JSON_UNESCAPED_UNICODE | JSON_UNESCAPED_SLASHES);
