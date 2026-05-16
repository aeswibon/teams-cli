# Regenerate with: vhs docs/demo.tape

Output "docs/demo.gif"

Set FontSize 20
Set Width 1200
Set Height 700
Set Theme "Builtin Solarized Dark"
Set TypingSpeed 60ms

Type "go run . --help"
Enter
Sleep 1s

Type "go run . init --help"
Enter
Sleep 1s

# Demo is intentionally tokenless; doctor should fail safely.
Type "go run . doctor || true"
Enter
Sleep 1s

Type "go run . chats --help"
Enter
Sleep 1s

Type "go run . messages --help"
Enter
Sleep 1s

Type "go run . request --help"
Enter
Sleep 1s
