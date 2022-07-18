# Roman

[![Code Coverage](https://codecov.io/gh/UltiRequiem/roman/branch/main/graph/badge.svg)](https://codecov.io/gh/UltiRequiem/roman)
[![Go Reference](https://pkg.go.dev/badge/github.com/UltiRequiem/roman.svg)](https://pkg.go.dev/github.com/UltiRequiem/roman)

Convert Numbers to Roman numerals.

## Usage

```go
package main

import (
	"fmt"
	"github.com/UltiRequiem/roman"
)

func main() {
	roman, _ := roman.ConvertToRoman(33)

	fmt.Println(roman) //=> XXXIII

        fmt.Println(roman.ParseRoman("XXXIII")) //=> 33
}
```

It includes other helper methods, check the docs.

## Documentation

Is hosted on [Go Doc](https://pkg.go.dev/github.com/UltiRequiem/roman) 📄

## Support

Open an Issue, I will check it a soon as possible 👀

If you want to hurry me up a bit
[send me a tweet](https://twitter.com/UltiRequiem) 😆

Consider [supporting me on Patreon](https://patreon.com/UltiRequiem) if you like
my work 🙏

Don't forget to star the repo ⭐

## Authors

[Eliaz Bobadilla](https://ultirequiem.com) - Creator and Maintainer 💪

See also the full list of
[contributors](https://github.com/UltiRequiem/roman/contributors) who
participated in this project ✨

## Versioning

We use [Semantic Versioning](http://semver.org). For the versions available, see
the [tags](https://github.com/UltiRequiem/roman/tags) 🏷️

## Licence

Licensed under the MIT License 📄
