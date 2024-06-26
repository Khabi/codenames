# codenames 🕵️ [![Go Reference](https://pkg.go.dev/badge/github.com/Khabi/codenames.svg)](https://pkg.go.dev/github.com/Khabi/codenames)

Codenames is a simple library for generating short combinatons of random words ala docker's name generator.

The starting dictionary of words was AI generated.

When using the default implementaton, words are chosen at random from the included dictionary.  While some words independently are innocuous, when combined with other words they may take on different meanings.  This library does not by itself try to restrict that, but it does have features that can be used to limit the impact of them.  Caution is advised and if the enviroment requires it, you may use your own wordlist with the package.

## Default Usage

```go
slug := codenames.New()
fmt.Println(slug)

// Output: victorious-stout-bone-impact
```

## Custom dictionaries
You may use your own custom dictionaries either with the built in ones or without.

```go

var planets = []string{"Abydos", "Chulak", "Hanka", "Othalla"}

func Planet() string {
	i := rand.Intn(len(planets))
	return planets[i]
}

slug := codenames.Generate("-", nil, codenames.Color, Planet)
fmt.Println(slug)

// Output: azure-Chulak
```