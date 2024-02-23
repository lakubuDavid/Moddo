# Why ?

So I started experimenting with Go last week and decided that this would be a good fit as a test project because :
- It teaches me text manipulation in Go
- It teaches me how to uses bascic data structures
- Also I use a lot of different different languages with interaction between them in some projects and it was the tool I needed.

Anyway I will write a small doc later since it's still pretty early and I also need to add some comments to the code (even though I believe the code is understandable...maybe)


# Build & Run

`go run . <model_definition_file> --lang=<...> [--output-dir="..."] [--file-case=...]`

or if already built :

`moddo <model_definition_file> --lang=<...> [--output-dir="..."] [--file-case=...]`


# Arguments

- `--lang` : The output language
    - `ts`,`ts-int` : Typescript and Typescript interface
    - `cs`,`cs-record`,`cs-props` : C# class, C# record class,C# class with props
    - `java`,`java-props` : Java class, Java class with props
    - `py` : Python class
    - `php` : Php class
    - `teal` : Teal record, teal is superset of lua with types ,a little like typescript with javascript
- `--output-dir` (optional): The output dir
- `--file-case` (optional): Specify how the file names will be formated, eache language already as it's own default file case rule but if needed you can always change it
    - camel : CamelCase
    - lowerCamel : lowerCamelCase
    - snake : snake_case

# Example model

```
# This is a comment
package Models

model User
    id      int     @id
    name    string
    email   string
    posts   string  @many
    password string @writeonly
    birthdate string @readonly

model Car
    id      int
    owner   string
    color   string
    paint   gold
    fuel    number
```

> The syntaxe is pretty basic,but **the indentation is important**
>
> Use either `Tab` or 4 spaces *

- `package NAME` : The namespace to use,only applies to languages that use it
- `model MODEL_NAME` : begining of model definition
- Properties are written like this :
```name type @modifier```
- Comments starts with a `#`

> Modifiers don't work for now except `@many` for collections/arrays (will write it later this week or if I ever retouch this code)

While you can have multiple models in a single file they will be generated in separate files.
Why ? Because it's easier and avoids having to check per language if I can or not put everything in a single file or ask if the user (you) want it in a signel file or not.
It's all about simplicity here.

## Types
Only basic primitive types :
- int
- string
- bool
- number
- Any unknown type will be set to a generic type for each language (`any` for typescript, `object` for C# and Java, `mixed` for php...)
- Array are made using the `@many` modifier

# What doesn't work

- Modifiers are useless for now except `@many`
- Only basic types

# The goal

- [x] Generate base classes
- [x] Have the `@many` modifier work
- [ ] Have the `@readonly` and `@writeonly` work for generators with properties like `java-props` and `cs-props`
- [ ] Some new ideas...

# Wanna help ?

**Don't.**

At least not now, but you can look at the code and get inspired from it.

# Special thanks to

- God
- My bed
- Youtube
- Perplexity & ChatGPT
- My fan (it's getting really hot here)
