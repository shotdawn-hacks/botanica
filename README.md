## Botanica

> Intellectual medical herbs database

### Architecture

Currently frontend, plants microservice, pdf-extractor

### How to start your own DB

You must provide mongoDB connection uri as env variable

```shell
export MONGO=mongodb:user:password@host:port
```

### How to start with our DB

To start services on machine simply run

```shell
docker compose up --build -d
```

### Team

shotdawn

1. Danila Artamonov (Backend\ML, MISIS)
2. Tani Suksina (Backend, MISIS)
3. Vsevolod (Frontend)
4. Yana (UX/UI)
