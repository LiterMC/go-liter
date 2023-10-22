
- **English**
- [中文](./README_zh.MD)

# Script

We use [goja](https://github.com/dop251/goja) as our javascript engine.  
Goja support full ES5, but it **DOES NOT** support full ES6 yet.  
Please read their [Features](https://github.com/dop251/goja#features) and [Known incompatibilities](https://github.com/dop251/goja#known-incompatibilities-and-caveats) to see what you can use and what you cannot

# TS support

1. Install the package below:
   ```sh
   npm install --save-dev go-liter-plugin
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

# Events

TODO
