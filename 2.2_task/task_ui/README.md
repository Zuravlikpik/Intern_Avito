# testboard-e2e-tests

E2E тесты на Playwright для Testboard.

### Требования

- Node.js 18+ (лучше LTS)
- npm 9+

### Установка

1) Установить зависимости:

```bash
npm install
```

### Использование вечного пользователя в тестах

Тесты логинятся в существующий аккаунт через ручку и прокидывают авторизацию в localStorage.

Для использования СВОЕГО пользователя нужно ввести его данные в файле .env

```bash
const USER_EMAIL = "email;
const USER_PASSWORD = "password";
```

### Запуск тестов

Запустить все тесты:

```bash
npx playwright test
```

Запустить конкретный тестовый файл reateAccount.test.ts:
```bash
npx playwright test tests/accountTests/createAccount.test.ts
```

Запустить конкретный тест по названию:
```bash
npx playwright test -g "Открытие страницы Мои объявления"
```
## Линтер (ESLint)

#### Проверить проект линтером:
```bash
npx eslint . 
```
Пустота означает, что ошибок нет
#### Запустить линтер для конкретного файла:
```bash
npx eslint fixtures/auth.fixture.ts --fix
```
#### Запустить линтер для всего проекта:
```bash
npx eslint . --fix
```
Если какие-то ошибки останутся, их нужно будет подправить руками

