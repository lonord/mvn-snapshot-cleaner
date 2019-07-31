# mvn-snapshot-cleaner
A cli tool to clean outdated snapshot items in maven repository

## Install

1. Using go get

```bash
$ go get github.com/lonord/mvn-snapshot-cleaner
```

2. Download from [release](https://github.com/lonord/mvn-snapshot-cleaner/releases)

## Usage

Just run it

```bash
$ mvn-snapshot-cleaner
```

And output

```
DELETE com.some.package1:some-package1:1.0-SNAPSHOT [1 history items]
DELETE com.some.package2:some-package2:1.0-SNAPSHOT [1 history items]
DELETE com.some.package3:some-package3:1.0-SNAPSHOT [1 history items]
DELETE com.some.package4:some-package4:1.0-SNAPSHOT [1 history items]
DELETE com.some.package5:some-package5:1.0-SNAPSHOT [1 history items]
DELETE com.some.package6:some-package6:1.0-SNAPSHOT [1 history items]
DELETE com.some.package7:some-package7:1.0-SNAPSHOT [1 history items]
DELETE com.some.package8:some-package8:1.0-SNAPSHOT [1 history items]
DELETE com.some.package9:some-package9:1.0-SNAPSHOT [1 history items]
DELETE com.some.package10:some-package10:1.0-SNAPSHOT [1 history items]
DELETE com.some.package11:some-package11:1.0-SNAPSHOT [1 history items]
DELETE com.some.package12:some-package12:1.0-SNAPSHOT [1 history items]
DELETE com.some.package13:some-package13:1.0-SNAPSHOT [1 history items]
==============================================
Total 13 entries cleaned, 245.20MB recycled :)
```

If your maven repository is not in path `~/.m2/repository`, specify it with `-r` flag

```bash
$ mvn-snapshot-cleaner -r path/to/your/repository
```

## License

MIT
