go version

if [ $? -eq 127 ] 
then
    wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
    sudo -i
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
    go version
    return 0
fi
