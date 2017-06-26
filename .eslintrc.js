module.exports = {
  "parser": "babel-eslint",
  "env": {
    "es6": true,
    "node": true,
    "mocha": true,
    "browser": true,
  },
  "plugins": [ "eslint-plugin-babel", "react" ],
  "extends": "eslint:recommended",
  "parserOptions": {
    "sourceType": "module",
    "ecmaFeatures": {
      "jsx": true
    }
  },
  "rules": {
    "accessor-pairs": 2,
    "array-bracket-spacing": 0,
    "array-callback-return": 2,
    "arrow-body-style": 0,
    "arrow-parens": 0,
    "arrow-spacing": [
      2,
      {
        "after": true,
        "before": true
      }
    ],
    "babel/generator-star-spacing": 1,
    "babel/new-cap": 1,
    "babel/array-bracket-spacing": 1,
    "babel/object-curly-spacing": [ "error", "always", {
      "objectsInObjects": false,
      "arraysInObjects": false
    } ],
    "babel/object-shorthand": 1,
    "babel/arrow-parens": 0,
    "babel/no-await-in-loop": 1,
    "babel/flow-object-type": 1,
    "block-scoped-var": 2,
    "block-spacing": [
      2,
      "always"
    ],
    "brace-style": [
      2,
      "1tbs",
      {
        "allowSingleLine": true
      }
    ],
    "callback-return": 2,
    "camelcase": 0,
    "comma-dangle": 0,
    "comma-spacing": [
      2,
      {
        "after": true,
        "before": false
      }
    ],
    "comma-style": [
      2,
      "last"
    ],
    "complexity": 2,
    "computed-property-spacing": [
      2,
      "never"
    ],
    "consistent-return": 0,
    "consistent-this": 2,
    "curly": 0,
    "default-case": 2,
    "dot-location": [
      2,
      "property"
    ],
    "dot-notation": [
      2,
      {
        "allowKeywords": true
      }
    ],
    "eol-last": 2,
    "eqeqeq": 0,
    "func-names": 0,
    "func-style": 0,
    "generator-star-spacing": 0,
    "global-require": 0,
    "guard-for-in": 2,
    "handle-callback-err": 0,
    "id-blacklist": 2,
    "id-length": 0,
    "id-match": 2,
    "indent": 0,
    "init-declarations": 0,
    "jsx-quotes": 2,
    "key-spacing": 2,
    "keyword-spacing": [
      2,
      {
        "after": true,
        "before": true
      }
    ],
    "linebreak-style": [
      2,
      "unix"
    ],
    "lines-around-comment": 2,
    "max-depth": 2,
    "max-len": 0,
    "max-nested-callbacks": 2,
    "max-params": 0,
    "new-cap": 0,
    "new-parens": 2,
    "newline-after-var": 0,
    "newline-per-chained-call": 0,
    "no-alert": 2,
    "no-array-constructor": 2,
    "no-bitwise": 2,
    "no-caller": 2,
    "no-catch-shadow": 0,
    "no-confusing-arrow": 0,
    "no-continue": 2,
    "no-console" : 0,
    "no-div-regex": 2,
    "no-else-return": 2,
    "no-empty-function": 0,
    "no-eq-null": 2,
    "no-eval": 2,
    "no-extend-native": 2,
    "no-extra-bind": 2,
    "no-extra-label": 2,
    "no-extra-parens": 2,
    "no-extra-semi": 2,
    "no-floating-decimal": 2,
    "no-implicit-coercion": 2,
    "no-implicit-globals": 2,
    "no-implied-eval": 2,
    "no-inline-comments": 0,
    "no-invalid-this": 2,
    "no-iterator": 2,
    "no-label-var": 2,
    "no-labels": 2,
    "no-lone-blocks": 2,
    "no-lonely-if": 2,
    "no-loop-func": 2,
    "no-magic-numbers": 0,
    "no-mixed-requires": 2,
    "no-multi-spaces": 2,
    "no-multi-str": 2,
    "no-multiple-empty-lines": 2,
    "no-native-reassign": 2,
    "no-negated-condition": 0,
    "no-nested-ternary": 2,
    "no-new": 2,
    "no-new-func": 2,
    "no-new-object": 2,
    "no-new-require": 2,
    "no-new-wrappers": 2,
    "no-octal-escape": 2,
    "no-param-reassign": 0,
    "no-path-concat": 2,
    "no-plusplus": 0,
    "no-process-env": 0,
    "no-process-exit": 0,
    "no-proto": 2,
    "no-restricted-imports": 2,
    "no-restricted-modules": 2,
    "no-restricted-syntax": 2,
    "no-return-assign": 2,
    "no-script-url": 2,
    "no-self-compare": 2,
    "no-sequences": 2,
    "no-shadow": 0,
    "no-shadow-restricted-names": 2,
    "no-spaced-func": 2,
    "no-sync": 0,
    "no-ternary": 0,
    "no-throw-literal": 2,
    "no-trailing-spaces": 2,
    "no-undef-init": 2,
    "no-undefined": 2,
    "no-underscore-dangle": 0,
    "no-unmodified-loop-condition": 2,
    "no-unneeded-ternary": 2,
    "no-unreachable": 2,
    "no-unused-vars": [
      2,
      {
        "vars": "all",
        "args": "none"
      }
    ],
    "no-use-before-define": 0,
    "no-useless-call": 2,
    "no-useless-concat": 2,
    "no-useless-constructor": 2,
    "no-var": 0,
    "no-void": 2,
    "no-warning-comments": 0,
    "no-whitespace-before-property": 2,
    "no-with": 2,
    "object-curly-spacing": 0,
    "object-shorthand": 0,
    "one-var": 0,
    "one-var-declaration-per-line": 2,
    "operator-assignment": 2,
    "operator-linebreak": 2,
    "padded-blocks": 0,
    "prefer-arrow-callback": 0,
    "prefer-const": 0,
    "prefer-reflect": 0,
    "prefer-rest-params": 2,
    "prefer-spread": 2,
    "prefer-template": 0,
    "quote-props": 0,
    "quotes": 0,
    "radix": 2,
    "require-jsdoc": 0,
    "require-yield": 2,
    "semi": 0,
    "semi-spacing": 2,
    "sort-imports": [ 0, {
      "ignoreCase": true
    } ],
    "sort-vars": 2,
    "space-before-blocks": 2,
    "space-before-function-paren": 0,
    "space-in-parens": [
      2,
      "never"
    ],
    "space-infix-ops": 2,
    "space-unary-ops": 2,
    "spaced-comment": 0,
    "strict": 0,
    "template-curly-spacing": 2,
    "valid-jsdoc": 0,
    "vars-on-top": 2,
    "wrap-iife": 2,
    "wrap-regex": 0,
    "yield-star-spacing": 2,
    "yoda": [
      2,
      "never"
    ],
    "react/jsx-uses-react": "error",
    "react/jsx-uses-vars": "error",
  }
};
