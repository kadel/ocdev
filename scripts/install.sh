#!/bin/bash
set -e

# The version of ocdev to install. Possible values - "master" and "latest"
# master - builds from git master branch
# latest - released versions specified by LATEST_VERSION variable
OCDEV_VERSION="latest"

# Latest released ocdev version
LATEST_VERSION="v0.0.1"

GITHUB_RELEASES_URL="https://github.com/redhat-developer/ocdev/releases/download/${LATEST_VERSION}"
BINTRAY_URL="https://dl.bintray.com/ocdev/ocdev/latest"

INSTALLATION_PATH="/usr/local/bin/"
PRIVILEGED_EXECUTION="sh -c"

DEBIAN_GPG_PUBLIC_KEY="https://bintray.com/user/downloadSubjectPublicKey?username=bintray"
DEBIAN_MASTER_REPOSITORY="https://dl.bintray.com/ocdev/ocdev-deb-dev"
DEBIAN_LATEST_REPOSITORY="https://dl.bintray.com/ocdev/ocdev-deb-releases"

RPM_MASTER_YUM_REPO="https://bintray.com/ocdev/ocdev-rpm-dev/rpm"
RPM_LATEST_YUM_REPO="https://bintray.com/ocdev/ocdev-rpm-releases/rpm"

SUPPORTED_PLATFORMS="
darwin-amd64
linux-amd64
linux-arm
"

echo_stderr ()
{
    echo "$@" >&2
}

command_exists() {
	command -v "$@" > /dev/null 2>&1
}

check_platform() {
    kernel="$(uname -s)"
    if [ "$(uname -m)" = "x86_64" ]; then
        arch="amd64"
    fi

    platform_type=$(echo "${kernel}-${arch}" | tr '[:upper:]' '[:lower:]')

    if ! echo "# $SUPPORTED_PLATFORMS" | grep "$platform_type" > /dev/null; then
        echo_stderr "
# The installer has detected your platform to be $platform_type, which is
# currently not supported by this installer script.

# Please visit the following URL for detailed installation steps:
# https://github.com/redhat-developer/ocdev/#installation

        "
        exit 1
    fi
    echo "$platform_type"
}

get_distribution() {
	lsb_dist=""
	if [ -r /etc/os-release ]; then
		lsb_dist="$(. /etc/os-release && echo "$ID")"
	fi
	echo "$lsb_dist"
}

set_privileged_execution() {
	if [ "$(id -u)" != "0" ]; then
        if command_exists sudo; then
            echo "# Installer will run privileged commands with sudo"
            PRIVILEGED_EXECUTION='sudo -E sh -c'
        elif command_exists su ; then
            echo "# Installer will run privileged commands with \"su -c\""
            PRIVILEGED_EXECUTION='su -c'
        else
    	    echo_stderr "# 
This installer needs to run as root. The current user is not root, and we could not find "sudo" or "su" installed on the system. Please run again with root privileges, or install "sudo" or "su" packages.
"
        fi
    else
        echo "# Installer is being run as root"
	fi
}

invalid_ocdev_version_error() {
    echo_stderr "# Invalid value of ocdev version provided, provide master or latest."
    exit 1
}

# install ocdev binary to /usr/local/bin/
install_ocdev_bin() {
    case "$OCDEV_VERSION" in
    master)
        BINARY_URL="$BINTRAY_URL/$platform/ocdev"
        echo "# Downloading ocdev from $BINARY_URL"
        curl -Lo ocdev "$BINARY_URL"
        ;;
    latest)
        BINARY_URL="$GITHUB_RELEASES_URL/ocdev-$platform.gz"
        echo "# Downloading ocdev from $BINARY_URL"
        curl -Lo ocdev.gz "$BINARY_URL"
        echo "# Extracting ocdev.gz"
        gunzip -d ocdev.gz
        ;;
    *)
        invalid_ocdev_version_error
    esac

    echo "# Setting execute permissions on ocdev"
    chmod +x ocdev
    echo "# Moving ocdev binary to $INSTALLATION_PATH"
    $PRIVILEGED_EXECUTION "mv ocdev $INSTALLATION_PATH"
    echo "# ocdev has been successfully installed on your machine"
}

# install ocdev on macOS via homebrew
install_ocdev_homebrew() {
    if ! command_exists brew; then
        echo_stderr "# brew command not found. Please install Homebrew and run this installer again. See https://brew.sh/ on how to install Homebrew."
    fi

    echo "# Enabling kadel/ocdev tap... "
    brew tap kadel/ocdev

    echo "Installing ocdev..."
    case "$OCDEV_VERSION" in
    master)
        brew install kadel/ocdev/ocdev -- HEAD
        ;;
    latest)
        brew install kadel/ocdev/ocdev
    esac

    return 0
}

# install ocdev using deb package
install_ocdev_deb(){
    echo "# Installing pre-requisites..."
    $PRIVILEGED_EXECUTION "apt-get update"
    $PRIVILEGED_EXECUTION "apt-get install -y gnupg apt-transport-https curl"

    echo "# "Adding GPG public key...
    $PRIVILEGED_EXECUTION "curl -L \"$DEBIAN_GPG_PUBLIC_KEY\" |  apt-key add -"

    echo "# Adding repository to /etc/apt/sources.list"
    case "$OCDEV_VERSION" in
    master)
        $PRIVILEGED_EXECUTION "echo \"deb $DEBIAN_MASTER_REPOSITORY stretch main\" |  tee -a /etc/apt/sources.list"
        ;;
    latest)
        $PRIVILEGED_EXECUTION "echo \"deb $DEBIAN_LATEST_REPOSITORY stretch main\" | tee -a /etc/apt/sources.list"
        ;;
    *)
        invalid_ocdev_version_error
    esac

    $PRIVILEGED_EXECUTION "apt-get update"
    $PRIVILEGED_EXECUTION "apt-get install -y ocdev"
}

# install ocdev using rpm pacakge
install_ocdev_rpm() {
    package_manager=""
    case "$distribution" in
    fedora)
        package_manager="dnf"
        ;;
    centos)
        package_manager="yum"
        ;;
    esac
    
    echo "# Installing pre-requisites..."
    $PRIVILEGED_EXECUTION "$package_manager install -y curl"

    echo "# Adding ocdev repo under /etc/yum.repos.d/"
    case "$OCDEV_VERSION" in

    master)
        $PRIVILEGED_EXECUTION "curl -L $RPM_MASTER_YUM_REPO -o /etc/yum.repos.d/bintray-ocdev-ocdev-rpm-dev.repo"
        ;;
    latest)
        $PRIVILEGED_EXECUTION "curl -L $RPM_LATEST_YUM_REPO -o /etc/yum.repos.d/bintray-ocdev-ocdev-rpm-releases.repo"
        ;;
    *)
        invalid_ocdev_version_error
    esac

    $PRIVILEGED_EXECUTION "$package_manager install -y ocdev"
}

# install ocdev using package manager
# automaticaly detects system and choses proper way to install ocdev (deb/rpm/homebrew)
# if detected system is uknown uses install_ocdev_bin to copy binary to /usr/local/bin
install_ocdev_pkg() {
    echo "# Starting ocdev installation..."
    echo "# Detecting distribution..."

    platform="$(check_platform)"
    echo "# Detected platform: $platform"

    if command_exists ocdev; then
        echo_stderr "# 
ocdev version \"$(ocdev version)\" is already installed on your system, running this installer script might case issues with your current installation. If you want to install ocdev using this script, please remove the current installation of ocdev from you system.
Aborting now!
"
        exit 1
    fi

    # macOS specific steps
    if [ "$platform" = "darwin-amd64" ]; then
        install_ocdev_homebrew
    fi

    set_privileged_execution

    distribution=$(get_distribution)
  	echo "# Detected distribution: $distribution"
  	echo "# Installing ocdev version: $OCDEV_VERSION"

    case "$distribution" in

    ubuntu|debian|linuxmint)
        install_ocdev_deb
        ;;

    centos|fedora)
        install_ocdev_rpm
        ;;

    *)
        echo "# Could not identify distribution, proceeding with a binary install..."
        install_ocdev_bin
        ;;
    esac
}

verify_ocdev() {
    if command_exists ocdev; then
        echo "
# Verification complete!
# ocdev version \"$(ocdev version)\" has been installed at $(type -P ocdev)
"
    else
        echo_stderr "
# Something is wrong with ocdev installation, please run the installaer script again. If the issue persists, please create an issue at https://github.com/redhat-developer/ocdev/issues"
        exit 1
    fi
}

install_ocdev_pkg
verify_ocdev