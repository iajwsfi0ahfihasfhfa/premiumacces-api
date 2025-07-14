# PremiumAccess API

PremiumAccess API to lekki serwis REST do zarządzania subskrypcjami premium.  
Oferuje symulowaną obsługę płatności i podstawowe endpointy do aktywacji subskrypcji, sprawdzania statusu i anulowania.

## Funkcje

- Aktywacja subskrypcji premium (symulowana płatność)
- Sprawdzanie statusu subskrypcji
- Anulowanie subskrypcji
- Endpoint zdrowia serwisu

## Szybki start

1. Klonuj repozytorium:
    ```sh
    git clone https://github.com/yourorg/premiumaccess.git
    cd premiumaccess
    ```

2. Uruchom serwer:
    ```sh
    go run main.go
    ```

3. API nasłuchuje na porcie `8080`.

## Endpointy

| Metoda | Endpoint           | Opis                           |
|--------|--------------------|--------------------------------|
| POST   | /api/subscribe     | Aktywuj subskrypcję            |
| GET    | /api/status        | Sprawdź status subskrypcji     |
| POST   | /api/cancel        | Anuluj aktywną subskrypcję     |
| GET    | /health            | Health check                   |

---

## Przykłady użycia

### Aktywacja subskrypcji

```sh
curl -X POST http://localhost:8080/api/subscribe \
  -H "Content-Type: application/json" \
  -d '{"user_id":"user42"}'
```

### Sprawdzenie statusu

```sh
curl "http://localhost:8080/api/status?user_id=user42"
```

### Anulowanie subskrypcji

```sh
curl -X POST http://localhost:8080/api/cancel \
  -H "Content-Type: application/json" \
  -d '{"user_id":"user42"}'
```

### Health check

```sh
curl http://localhost:8080/health
```

---

## Struktura projektu

```
premiumaccess/
├── main.go
├── go.mod
└── README.md
```

## Licencja

MIT License
