# Why ?

So I started experimenting with Go last week and decided that this would be a good fit as a test project because :
- It as uses a lot of text manipulation 
- It uses bascic data structures 
- Also I use a lot of different different language with interaction between them in some projects and it was the tool I needed

Anyway I will write a small doc later since it's still pretty early and I also need to add some comment (even though I believe the code is understandable...maybe)


# Build & Run

Just :
`go run . <model_definition_file> --lang=<java,ts,py,php,cs> [--output-dir="output/models/"]` 

or if already built :

`moddo <model_definition_file> --lang=<java,ts,py,php,cs> [--output-dir="output/models/"]`


# Arguments 

- `--lang` : The output language
    - `ts`,`ts-int` : Typescript and Typescript interface
    - `cs`,`cs-record`,`cs-props` : C# class, C# record class,C# class with props
    - `java`,`java-props` : Java class, Java class with props
    - `py` : Python class
    - `php` : Php class
- `--output-dir` (optional): The output dir

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
- properties : `name type @modifier`

> For now there are no collections
>
> It's supposed to use the @many modifier (makes the parser easy)
>
> Modifiers don't work (will write it later this week or if I ever retouch this code)

## Types 
Only basic primitive types :
- int
- string
- bool
- number
- Any unknown type will be set to a generic type for each language (`any` for typescript, `object` for C# and Java, `mixed` for php...)

# What doesn't work

- Modifiers are useless for now
- Only basic types
- No collections

# The goal

- [x] Generate base classes
- [ ] have modifiers work (especially the @many one that will be used for collections)
- [ ] Some new ideas...

# Wanna help ?

Don't.

At least not now but you can look at the code and get inspired from it.

# Special thanks to

- God
- My bed
- Youtube
- Perplexity & ChatGPT
- My fan (it's getting really hot here)
  
