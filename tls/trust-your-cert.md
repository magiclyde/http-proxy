### 添加自签名证书

#### ubuntu
```
sudo cp server.crt /usr/local/share/ca-certificates/
sudo update-ca-certificates
```

#### centos
```
sudo cp server.crt /etc/pki/ca-trust/source/anchors/
sudo update-ca-trust
```
