cd "$srcdir/$pkgname-$pkgver"
install -Dm644 "LICENSE" -t "$pkgdir/usr/share/licenses/$pkgname"
install -Dm644 "README.md" -t "$pkgdir/usr/share/doc/$pkgname"
install -Dm644 "CODE_OF_CONDUCT.md" -t "$pkgdir/usr/share/doc/$pkgname"
install -Dm644 "CONTRIBUTING.md" -t "$pkgdir/usr/share/doc/$pkgname"
install -Dm755 "bin/$pkgname" -t "$pkgdir/usr/bin"
