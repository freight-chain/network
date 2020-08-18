#!/bin/bash
touch .prettierrc.json
cat << EOF > .prettierrc.json
{
  "trailingComma": "es5",
  "tabWidth": 4,
  "semi": false,
  "singleQuote": true
}
EOF
