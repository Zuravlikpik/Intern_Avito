# UI-тесты для [проекта](https://cerulean-praline-8e5aa6.netlify.app/)

### Требования

- Node.js 18+ (лучше LTS)
- npm 9+

### Установка 

1) Установить зависимости:

```bash
npm install
```

### Запуск тестов

перейдите в корень проекта:

```bash
   cd task_ui
```

Запустить все тесты:

```bash
npx playwright test
```

Запустить конкретный тестовый файл categoryFilter.test.ts:
```bash
npx playwright test --headed tests/list/categoryFilter.test.ts
```

Запустить конкретный тест по названию:
```bash
npx playwright test -g "Фильтр по категории"
```
### Важно перед запуском тестов
  
> Перед запуском тестов **обязательно подключите VPN** (или убедитесь, что используете сеть, из которой есть доступ к `cerulean-praline-8e5aa6.netlify.app`).  
> В противном случае все тесты завершатся с ошибкой.

## Линтер (ESLint)

#### Проверить проект линтером:
```bash
npx eslint . 
```
Пустота означает, что ошибок нет
#### Запустить линтер для конкретного файла:
```bash
npx eslint pages/listPage/listPage.ts --fix
```
#### Запустить линтер для всего проекта:
```bash
npx eslint . --fix
```

