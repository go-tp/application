#!/bin/bash
<< EOF
    @Author: gouzi 
    @Date: 2022-03-31 16:37:50 
    @Last Modified by: gouzi
    @Last Modified time: 2022-03-31 17:12:30
EOF


# env judge
# Linux
if [ $(uname) = "Linux" ];
then
    SED="sed"
    GETOPT="getopt"
fi
# Mac
if [ $(uname) = "Darwin" ];
then
    SED="gsed"
    GETOPT="/usr/local/opt/gnu-getopt/bin/getopt"
    # install gnu-sed
    if [ ! $(brew list|grep gnu-sed) ];
    then
        brew install gnu-sed
    fi
    # install gnu-getopt
    if [ ! $(brew list|grep gnu-getopt) ];
    then
        brew install gnu-getopt
    fi
fi

ARGS=`$GETOPT -o hva:b: --long help,version,arch,system-os -- "$@"`

help(){
    echo "usage: ./build.sh [-a arch] [-s os] [-v version]"
    echo "    -h, --help         : 帮助信息"
    echo ""
    echo "Options: "
    echo "    -a, --arch         : 指令集 选项 [amd64, 386, arm]"
    echo "    -s, --system-os     : 系统 选项 [linux, darwin, windows, mac(不成功的话去选darwin)]"
    echo "    -v, --version      : 版本 打包后的版本"
    exit 1
}

while true; do
    case "$1" in
        -h|--help ) help ; shift ;;
        -a|--arch ) ARCH=$2 ; shift 2 ;;
        -s|--system-os ) SYSTEMOS=$2 ; shift 2 ;;
        -v|--version ) VERSION=$2 ; shift 2 ;;
        -- ) shift; break ;;
        * ) break ;;
    esac
done

echo "start make\n"

# IS NULL
if [ ! -n "$ARCH" ]; then
    ARCH="amd64"
fi

if [ ! -n "$SYSTEMOS" ]; then
    SYSTEMOS="linux"
fi

if [ ! -n "$VERSION" ]; then
    VERSION=0.0.1
fi

# Check WAIT
echo "==============="
echo "Check Wait:\n"
echo "ARCH: $ARCH"
echo "SYSTEMOS: $SYSTEMOS"
echo "VERSION: $VERSION"
echo "===============\n"
sleep "5"

# RUN MAKE

# ARCH sed
if [ $ARCH = "amd64" ] || [ $ARCH = "Amd64" ] || [ $ARCH = "AMD64" ];
then
    ARCH="amd64"
    $SED -i 's/ARCH="amd64"/ARCH="'${ARCH}'"/' Makefile
fi

if [ $ARCH = "386" ];
then
    ARCH="386"
    $SED -i 's/ARCH="amd64"/ARCH="'${ARCH}'"/' Makefile
fi

if [ $ARCH = "arm" ] || [ $ARCH = "ARM" ] || [ $ARCH = "Arm" ];
then
    ARCH="arm"
    $SED -i 's/ARCH="amd64"/ARCH="'${ARCH}'"/' Makefile
fi

# system sed
if [ $SYSTEMOS = "linux" ] || [ $SYSTEMOS = "Linux" ] || [ $SYSTEMOS = "LINUX" ];
then
    SYSTEMOS="linux"
    $SED -i 's/SYSTEMOS="linux"/SYSTEMOS="'${SYSTEMOS}'"/' Makefile
fi

if [ $SYSTEMOS = "darwin" ] || [ $SYSTEMOS = "Darwin" ] || [ $SYSTEMOS = "DARWIN" ];
then
    SYSTEMOS="darwin"
    $SED -i 's/SYSTEMOS="linux"/SYSTEMOS="'${SYSTEMOS}'"/' Makefile
fi

if [ $SYSTEMOS = "mac" ]  || [ $SYSTEMOS = "Mac" ] || [ $SYSTEMOS = "MAC" ] || [ $SYSTEMOS = "MACOS" ] || [ $SYSTEMOS = "macos" ];
then
    SYSTEMOS="mac"
    $SED -i 's/SYSTEMOS="linux"/SYSTEMOS="'${SYSTEMOS}'"/' Makefile
fi

if [ $SYSTEMOS = "windows" ] || [ $SYSTEMOS = "Windows" ] || [ $SYSTEMOS = "WINDOWS" ] || [ $SYSTEMOS = "WIN" ] || [ $SYSTEMOS = "win" ] || [ $SYSTEMOS = "Win" ];
then
    SYSTEMOS="windows"
    $SED -i 's/SYSTEMOS="linux"/SYSTEMOS="'${SYSTEMOS}'"/' Makefile
fi

# version sed
$SED -i 's/VERSION=0.0.1/VERSION='${VERSION}'/' Makefile

# RUN Makefile
make

# RESET MAKE
$SED -i 's/ARCH="'${ARCH}'"/ARCH="amd64"/' Makefile
$SED -i 's/SYSTEMOS="'${SYSTEMOS}'"/SYSTEMOS="linux"/' Makefile
$SED -i 's/VERSION='${VERSION}'/VERSION=0.0.1/' Makefile

echo "end make Finish\n"