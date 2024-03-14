<?php
$siteRootPath = $argv[1];
$configPath = $siteRootPath."/app/etc/config.php";
$composerJson = $siteRootPath."/composer.json";
$composerLockJson = $siteRootPath."/composer.lock";

$appCodePath = $siteRootPath."/app/code";
$vendorPath = $siteRootPath."/vendor";

if(file_exists($configPath)){
    $configData = include_once($configPath);
    $modules = $configData['modules'];
    $psr4 = include_once($vendorPath."/composer/autoload_psr4.php");
    $composer = file_get_contents($composerJson);
    $composer = json_decode($composer, true);
    $composerLock = file_get_contents($composerLockJson);
    $composerLock = json_decode($composerLock, true);
    print("Magento Version\n");
    $magentoVersion = "Unknown";
    if(!empty($composer['require']['magento/product-enterprise-edition'])){
        $magentoVersion = "Enterprise edition ".$composer['require']['magento/product-enterprise-edition'];
    } elseif(!empty($composer['require']['magento/product-community-edition'])) {
        $magentoVersion = "Community edition ".$composer['require']['magento/product-community-edition'];
    } elseif(!empty($composer['require']['magento/magento-cloud-metapackage'])) {
        $magentoVersion = "Enterprise edition on Adobe Cloud ".$composer['require']['magento/magento-cloud-metapackage'];
    }
    
    print($magentoVersion."\n\n");
    print("Third-parties modules\n");
    print("Name, Current version,  Latest version, Status\n");
    
    foreach($modules as $moduleName => $isActive) {
        if(strpos($moduleName, "Magento_", 0) !== 0){
            $version = "\"no version\"";
            $latestVersion = null;
            $modulePath = $psr4[str_replace("_", "\\", $moduleName)."\\"][0]??false;
            if($modulePath && file_exists($modulePath."/composer.json")) {
                $composerData = file_get_contents($modulePath."/composer.json");
                $composerData = json_decode($composerData, true);
                if(empty($composer['require'][$composerData['name']]) && !searchInLock($composerLock['packages'], $composerData['name'])){
                    continue;
                }
                $version = $composerData['version']??$version;
                $latests = [];
                if(exec("composer show --all -f json ".$composerData['name'], $latests) !== false){
                    $latests = json_decode(implode("", $latests), true);
                    if(!empty($latests['versions'])) {
                        if(count($latests['versions'])==1){
                            $latestVersion = $latests['versions'][0]??null;
                        } else {
                            foreach($latests['versions'] as $v) {
                                if(preg_match("/[a-z]+/i", $v) === 0){
                                    $latestVersion = $v;
                                    break;
                                }
                            }
                        }
                    }
                }
            } else {
                $modVN = explode("_", $moduleName);
                if(file_exists($appCodePath."/".$modVN[0]."/".$modVN[1]."/composer.json")) {
                    $composerData = file_get_contents($appCodePath."/".$modVN[0]."/".$modVN[1]."/composer.json");
                    $composerData = json_decode($composerData, true);
                    $version = $composerData['version']??$version;
                }
            }

            if(!$latestVersion) {
                $latestVersion = $version;
            }

            print($moduleName.", ".$version.", ".$latestVersion.", ".($isActive==1?"enabled":"disabled")."\n");
        }
    }
}

function searchInLock($packages, $name) {
    foreach($packages as $package) {
        if($package['name'] == $name) {
            return true;
        }
    }

    return false;
}

function prepareLatest($latests) {
    $arr = [];
    foreach($latests as $latest) {
        $arr[$latest['name']] = $latest;
    }

    return $arr;
}