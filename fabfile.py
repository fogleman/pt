from fabric.api import *
from fabric_gce_tools import *

# env.user = 'ubuntu'
# env.key_filename = '/Users/fogleman/home.pem'

# env.hosts = [
# ]

@parallel
@roles('pt')
def init():
    run('sudo apt-get update')
    # run('sudo apt-get --assume-yes upgrade')
    run('sudo apt-get --assume-yes install git build-essential')
    run('wget https://github.com/embree/embree/releases/download/v2.16.2/embree-2.16.2.x86_64.linux.tar.gz')
    run('tar xzf embree-2.16.2.x86_64.linux.tar.gz')
    run('rm embree-2.16.2.x86_64.linux.tar.gz')
    run('echo "source embree-2.16.2.x86_64.linux/embree-vars.sh" >> .profile')
    run('wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz')
    run('sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz')
    run('rm go1.8.3.linux-amd64.tar.gz')
    run('echo "export PATH=$PATH:/usr/local/go/bin" >> .profile')
    run('echo "export GOPATH=$HOME/go" >> .profile')
    run('source .profile')
    run('go get -u github.com/fogleman/go-embree')
    run('go get -u github.com/fogleman/pt/pt')
    with cd('~/go/src/github.com/fogleman/pt'):
        run('git checkout embree')

@parallel
@roles('pt')
def upload(path):
    put(path)

@parallel
@roles('pt')
def start(command):
    command += ' 2>stderr.txt >stdout.txt &'
    run(command)

@parallel
@roles('pt')
def latest():
    i = env.hosts.index(env.host)
    filename = 'latest%d.png' % i
    run('cp `ls out* | tail -n 2 | head -n 1` latest.png')
    get('latest.png', filename)

update_roles_gce(False)
