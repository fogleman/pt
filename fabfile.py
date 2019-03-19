from fabric.api import *

env.user = 'ubuntu'
env.key_filename = '/Users/fogleman/home.pem'

env.hosts = [
    '127.0.0.1',
]

def init():
    run('sudo apt-get --assume-yes install git')
    run('wget https://storage.googleapis.com/golang/go1.6.3.linux-amd64.tar.gz')
    run('sudo tar -C /usr/local -xzf go1.6.3.linux-amd64.tar.gz')
    run('echo "export PATH=$PATH:/usr/local/go/bin" >> .profile')
    run('echo "export GOPATH=$HOME/go" >> .profile')
    run('source .profile')
    run('go get github.com/hborntraeger/pt')

def fetch():
    i = env.hosts.index(env.host)
    filename = 'fetch%d.tar.gz' % i
    with cd('~/go/src/github.com/hborntraeger/pt'):
        run('tar czf ~/fetch.tar.gz *.png')
    get('fetch.tar.gz', filename)
    local('tar xzf ' + filename)
    local('rm ' + filename)

def latest():
    i = env.hosts.index(env.host)
    filename = 'latest%d.png' % i
    with cd('~/go/src/github.com/hborntraeger/pt'):
        run('cp `ls out* | tail -n 2 | head -n 1` ~/latest.png')
    get('latest.png', filename)
