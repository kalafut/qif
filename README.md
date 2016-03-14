# qif
QIF import for Go.

# About
QIF is an old, flawed financial data format. But it is also very common, and may be the only usable output from older financial software. This library will help you get data out of QIF and into Go structs. Only import is supported, as are only a subset of the possible QIF commands.

# Usage
...TODO...
The only conversion to native Go types besides `string` is the date (TODO: locale stuff). Converting Amount was tempting, but it's unclear that other formats would be any better. Strings are simple and universally consumable.

# Sample
...TODO...

# Contributing
Though the QIF format itself is lacking, it is at least pretty simple and just plain text. A loop and a big ass switch statement is usually parser enough to extract what you're looking for, and that's exactly what I've done. I've added support for the commands I have an actual need for and example of, but others can be added pretty easily. Submit a PR along with a sample file / test and I'll be happy to add it.

