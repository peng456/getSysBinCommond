<?php
//var_dump($_ENV);
if (!is_cli()){
echo "please run me whit php-cli!!! \n";
exit;
}

echo "PHP VERSION:".phpversion();
echo "\n";

$env = getenv();

foreach (explode(":",$env['PATH']) as $value){
  echo "\n";
  echo "=======".$value."===start=============";
  echo "\n";
  $filenames = get_filenamesbydir($value);
  //打印所有文件名，包括路径 
  foreach ($filenames as $commandvalue) { 
      echo $commandvalue."\n"; 
  }
  echo "=======".$value."===end=============";
  echo "\n";
}

 /*
 判断当前的运行环境是否是cli模式
 */
 function is_cli(){
     return preg_match("/cli/i", php_sapi_name()) ? true : false;
 }

function get_allfiles($path,&$files) { 

  if(is_dir($path)){ 

    $dp = dir($path); 

    while ($file = $dp ->read()){ 

      if($file !="." && $file !=".."){ 

        get_allfiles($path."/".$file, $files); 

      } 

    } 

    $dp ->close(); 

  } 

  if(is_file($path)){ 

    $files[] = $path; 

  } 

} 


function get_filenamesbydir($dir){ 

  $files = array(); 

  get_allfiles($dir,$files); 

  return $files; 
}
