# Kata Setup for CoffeeScript

### run specs
```
npm install
jasmine-node spec --coffee
```

### getting started

clone repository
```
git clone https://github.com/fnbk/kata_setup_coffeescript
```

delete .git files, delete package.json, rename directory, create new repository
```
rm -rf .git/
rm package.json
cd ../
mv kata_setup_coffeescript fizz_buzz_kata
cd fizz_buzz_kata
git init
```

setup node project
```
npm init
```

add jasmine_node to package.json
```
"devDependencies": {
  "jasmine-node": "~1.4.x"
}
```


