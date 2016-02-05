from fabric.api import *

env.user = 'ubuntu'
env.key_filename = '/Users/fogleman/home.pem'

env.hosts = [
    'a.name.com',
    'b.name.com',
]

def init():
    run('sudo apt-get --assume-yes install git')
    run('wget https://storage.googleapis.com/golang/go1.5.3.linux-amd64.tar.gz')
    run('sudo tar -C /usr/local -xzf go1.5.3.linux-amd64.tar.gz')
    run('echo "export PATH=$PATH:/usr/local/go/bin" >> .profile')
    run('echo "export GOPATH=$HOME/go" >> .profile')
    run('source .profile')
    run('go get github.com/fogleman/pt')

def run():
    i = env.hosts.index(env.host)
    n = len(env.hosts)
    print i, n

def fetch():
    with cd('~/go/src/github.com/fogleman/pt'):
        run('tar czf ~/fetch.tar.gz *.png')
    get('fetch.tar.gz', 'fetch.tar.gz')
    local('tar xzf fetch.tar.gz')
