
- **English**
- [中文](./README_zh.MD)

# Script

We use [goja](https://github.com/dop251/goja) as our javascript engine.  
Goja support full ES5, but it **DOES NOT** support full ES6 yet.  
Please read their [Features](https://github.com/dop251/goja#features) and [Known incompatibilities](https://github.com/dop251/goja#known-incompatibilities-and-caveats) to see what you can use and what you cannot

# Setup

## Lazy setup

```sh
npm init glp@latest "<directory name>"
```

Note: Typescript will be automaticly configured.

## Manually setup

You may want to configure your project manually, then you have to do the commands below

1. Install [go-liter-plugin](https://www.npmjs.com/package/go-liter-plugin):
   ```sh
   npm install go-liter-plugin
   ```

2. Add the following element to your `tsconfig.json`
   ```js
   {
      "compilerOptions": {
         "typeRoots": [
            "node_modules/@types",
            "node_modules/go-liter-plugin"
         ]
      }
   }
   ```

# Typescript

Check files under [./types/lib/](./types/lib/) for more information

# Events

TODO
