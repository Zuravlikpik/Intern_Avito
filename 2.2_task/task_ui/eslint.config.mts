import js from "@eslint/js";
import globals from "globals";
import tseslint from "typescript-eslint";

export default [
    {
        ignores: ["node_modules/**"]
    },
    {
        files: ["**/*.{js,mjs,cjs,ts,mts,cts}"],
        ...js.configs.recommended,
        languageOptions: {
            globals: globals.browser
        },
        plugins: {
            js
        }
    },
    {
        rules: {
            "quotes": ["error", "double"],
            "eol-last": ["error", "always"],
            "semi": ["error", "always"],
            "indent": ["error", 4],
            "linebreak-style": ["error", "unix"],
            "prefer-const": "error",
            "@typescript-eslint/no-unused-vars": [
                "error",
                {
                    "argsIgnorePattern": "^_",
                    "varsIgnorePattern": "^_",
                    "caughtErrorsIgnorePattern": "^_"
                }
            ]
        }
    },
    ...tseslint.configs.recommended
];
