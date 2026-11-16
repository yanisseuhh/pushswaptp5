# Push-Swap

## Compilation
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker

## Exécution
./push-swap "2 1 3 6 5 8"
./push-swap "2 1 3 6 5 8" | ./checker "2 1 3 6 5 8"

## Tests
go test -v ./...
cd testeur && go run . ..

## Stratégie
- N <= 6 : force brute (BFS).
- N > 6 : normalisation par rangs puis tri Radix binaire.
- Complexité : O(N log N).

## Gestion des erreurs
Affiche Error sur stderr et quitte avec le code 1 :
- Entrées non numériques ou mal formatées
- Doublons
- Dépassement des limites int32
- Instruction inconnue sur le checker