
# File structure

## Go

### `std` — Standard Go project

```
<project>/
├── cmd/
├── internal/
└── pkg/
```

### `api` — Go API server

```
<project>/
├── cmd/
│   └── api/
├── internal/
│   ├── handlers/
│   └── models/
└── pkg/
```

### `cli` — Go CLI application

```
<project>/
├── cmd/
├── internal/
│   └── config/
└── pkg/
    └── utils/
```

---

## TypeScript

### `std` — Standard TypeScript project

```
<project>/
├── src/
└── dist/
```

### `api` — TypeScript API server

```
<project>/
├── src/
│   ├── config/
│   ├── controllers/
│   ├── middlewares/
│   ├── models/
│   ├── routes/
│   ├── services/
│   └── types/
└── tests/
```

### `nxtjs` — Next.js project

```
<project>/
├── public/
└── src/
    ├── app/
    ├── components/
    ├── features/
    ├── lib/
    ├── hooks/
    └── types/
```

### `lib` — TypeScript library

```
<project>/
└── src/
    ├── utils/
    ├── types/
    └── __tests__/
```

---

## C++

### `std` — Standard C++ project

```
<project>/
├── src/
├── tests/
└── build/
```

### `app` — C++ application/games

```
<project>/
├── src/
├── include/
├── tests/
├── assets/
└── cmake/
```

### `lib` — C++ library

```
<project>/
├── src/
├── include/
│   └── <project>/
├── tests/
└── examples/
```
