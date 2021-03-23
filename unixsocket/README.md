# Unix socket example in golang


## php client

```
<?php
  $sock = stream_socket_client('unix:///home/utsaina/mio/echo.sock', $errno, $errst);
  fwrite($sock, 'message');
  $resp = fread($sock, 4096);
  fclose($sock);
  echo ("resp: " . $resp);  
?>
```