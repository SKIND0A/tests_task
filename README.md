## Тестирование и покрытие

Все тест-кейсы покрывают позитивные сценарии, граничные условия и негативные ветки валидации.

| Пакет | Тесты | Результат | Покрытие | Время |
| :--- | :---: | :---: | :---: | :---: |
| `validator` | 14 / 14 |  PASS | **100.0%** | 0.009s |


```text
=== RUN   TestValidatePassword
=== RUN   TestValidatePassword/password_is_empty
=== RUN   TestValidatePassword/password_is_empty#01
=== RUN   TestValidatePassword/password_is_too_short
=== RUN   TestValidatePassword/password_is_too_long
=== RUN   TestValidatePassword/password_is_too_common
=== RUN   TestValidatePassword/password_is_too_common#01
=== RUN   TestValidatePassword/password_is_too_common#02
=== RUN   TestValidatePassword/password_is_too_common#03
=== RUN   TestValidatePassword/password_is_too_common#04
=== RUN   TestValidatePassword/password_must_not_contain_spaces
=== RUN   TestValidatePassword/password_must_contain_uppercase_letter
=== RUN   TestValidatePassword/password_must_contain_lowercase_letter
=== RUN   TestValidatePassword/password_must_contain_digit
=== RUN   TestValidatePassword/strong_password
--- PASS: TestValidatePassword (0.00s)
    --- PASS: TestValidatePassword/password_is_empty (0.00s)
    --- PASS: TestValidatePassword/password_is_empty#01 (0.00s)
    --- PASS: TestValidatePassword/password_is_too_short (0.00s)
    --- PASS: TestValidatePassword/password_is_too_long (0.00s)
    --- PASS: TestValidatePassword/password_is_too_common (0.00s)
    --- PASS: TestValidatePassword/password_is_too_common#01 (0.00s)
    --- PASS: TestValidatePassword/password_is_too_common#02 (0.00s)
    --- PASS: TestValidatePassword/password_is_too_common#03 (0.00s)
    --- PASS: TestValidatePassword/password_is_too_common#04 (0.00s)
    --- PASS: TestValidatePassword/password_must_not_contain_spaces (0.00s)
    --- PASS: TestValidatePassword/password_must_contain_uppercase_letter (0.00s)
    --- PASS: TestValidatePassword/password_must_contain_lowercase_letter (0.00s)
    --- PASS: TestValidatePassword/password_must_contain_digit (0.00s)
    --- PASS: TestValidatePassword/strong_password (0.00s)
PASS
coverage: 100.0% of statements
ok      study/go/tests_tasks    0.009s
