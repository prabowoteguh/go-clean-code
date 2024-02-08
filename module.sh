go mod edit -module go-assignment-bootcamp
find . -type f -name '*.go' \
  -exec sed -i -e 's,github.com/d-alejandro/go-code-examples,go-assignment-bootcamp,g' {} \;