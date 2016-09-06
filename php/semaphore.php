<?PHP

define('PROC_NUM',10);
define('TIME_OUT', 60);                                                                                           
$ppid = posix_getpid();
$arr = array(3,4,9,2,9);
$s_arr = serialize($arr);                                                                                         

$shmR = shm_attach(ftok(__FILE__,'i').$ppid, strlen($s_arr)*2);                                                   

if($shmR == false){
    echo 'attach fail; ';                                                                                         
}
if ( false == shm_put_var($shmR, 1, $arr)){                                                                       
    echo 'put var fail ';                                                                                         
}                                                                                                                 

$semR = sem_get($ppid, 1);                                                                                        
                                                                                                                  
                                                                                                                  

$shm_id = $pid;
for($i=0; $i<PROC_NUM; $i++){                                                                                     
    $pid = pcntl_fork();                                                                                          
    if($pid == 0){                                                                                                
        while(1){
            sem_acquire($semR);
            $arrCur = shm_get_var($shmR, 1);                                                                      
            $row = array_pop($arrCur);                                                                            
            sleep(rand(1,3));
            if(false == shm_put_var($shmR, 1, $arrCur)){                                                          
                dir('2 put var fail');                                                                            
            }
            sem_release($semR);                                                                                   
            $res[] = $row;                                                                                        
        }
        echo 'child proc done ; ';                                                                                
        echo implode(' ',$res);                                                                                   
        exit(0);                                                                                                  
    }                                                                                                             
}
$start = time();                                                                                                  
$n = 0;
while( $n < PROC_NUM && time()-$start<TIME_OUT){                                                                  
    $status = -1;
    $npid = pcntl_wait($status, WNOHANG);                                                                         
    if($npid>0){
        echo $npid.' exit; ';                                                                                     
        ++$n;                                                                                                     
    }
    sleep(1);                                                                                                     
}                                                                                                                 

sem_remove($semR);                                                                                                
sem_remove($semR);                                                                                                
echo 'Finished';                                                                                                  
               
