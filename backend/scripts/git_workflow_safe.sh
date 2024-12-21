#!/bin/bash

# Проверяем, передано ли сообщение коммита
if [ -z "$1" ]; then
  echo "Ошибка: Вы должны передать сообщение коммита в качестве аргумента."
  echo "Пример: ./git_workflow_safe.sh \"Ваше сообщение коммита\""
  exit 1
fi

# Сохраняем сообщение коммита
COMMIT_MESSAGE="$1"

# Функция для обработки ошибок
handle_error() {
  echo "Ошибка на этапе: $1"
  echo "Скрипт завершён с ошибкой."
  exit 1
}

# Сбрасываем последний коммит в soft режиме
echo "Сбрасываем последний коммит в soft режиме..."
git reset --soft HEAD~1 || handle_error "git reset --soft HEAD~1"

# Сохраняем изменения в stash
echo "Сохраняем изменения в stash..."
git stash || handle_error "git stash"

# Получаем изменения из удалённого репозитория
echo "Получаем изменения из удалённого репозитория..."
git pull || handle_error "git pull"

# Восстанавливаем изменения из stash
echo "Восстанавливаем изменения из stash..."
git stash pop || handle_error "git stash pop"

# Добавляем изменения в индекс
echo "Добавляем изменения в индекс..."
git add . || handle_error "git add ."

# Проверяем, нужно ли продолжать rebase
if git status | grep -q "rebase in progress"; then
  echo "Продолжаем rebase..."
  git rebase --continue || handle_error "git rebase --continue"
fi

# Создаём новый коммит
echo "Создаём новый коммит с сообщением: \"$COMMIT_MESSAGE\""
git commit -m "$COMMIT_MESSAGE" || handle_error "git commit"

# Отправляем изменения в удалённый репозиторий
echo "Отправляем изменения в удалённый репозиторий..."
git push || handle_error "git push"

echo "Готово! Все изменения успешно отправлены."

