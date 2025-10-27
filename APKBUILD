# Maintainer: newsworthy39 <newsworthy39@github.com>
pkgname=cniplugincloudflare
pkgver=1.0.0
pkgrel=1
pkgdesc="cniplugincloudflare is a smalle go-program to manage dns-adresses in cloudclare, through a cni-definitionan."
url="https://github.com/newsworthy39/golang-cniplugincloudflare"
arch="all"
license="MIT"
depends=""
makedepends="bash"
source="$pkgname"
options="!check"
builddir="."
install="$pkgname.post-install $pkgname.post-deinstall"

#build() {
#    go build -o "$builddir/$pkgname" .
#}

package() {
    install -Dm755 "$builddir/$pkgname" "$pkgdir/usr/lib/cni/$pkgname"
}

sha512sums="
2d043da0ce47bbdd3c2843806362a03620b5828ac7598373cc23d91c5933c80b59446de2de35606a678ff6313cd38af6191f72d9b6f58f3cf97a98892f2617d8  cniplugincloudflare
"
