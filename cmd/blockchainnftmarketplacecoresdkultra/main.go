// cmd/blockchainnftmarketplacecoresdkultra/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftmarketplacecoresdkultra/internal/blockchainnftmarketplacecoresdkultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftmarketplacecoresdkultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
