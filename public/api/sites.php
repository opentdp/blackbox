<?php

$format = 'json';
if (isset($_GET['format']) && in_array($_GET['format'], ['yaml'])) {
    $format = $_GET['format'];
}

header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: GET');
header('Access-Control-Allow-Headers: Content-Type');
header('Content-Type: application/' . $format . ';charset=utf-8');

include dirname(__DIR__) . '/vendor.php';

// fetch sites

if (!$sites = get_sites($_GET['name'] ?? '')) {
    if ($format == 'json') {
        exit(json_encode(['error' => 'not found']));
        exit;
    }
    exit('error: not found');
}

ksort($sites);

// output

if ($format == 'json') {
    echo json_encode($sites, JSON_UNESCAPED_UNICODE | JSON_UNESCAPED_SLASHES);
    exit;
}

if ($format == 'yaml') {
    $output = [];
    foreach ($sites as &$site) {
        $output[] = "- targets:";
        foreach ($site['targets'] as $url) {
            $output[] = "    - $url";
        }
        $output[] = "  labels:";
        foreach ($site['labels'] as $key => $value) {
            $output[] = "    $key: $value";
        }
    }
    echo implode("\n", $output);
    exit;
}
